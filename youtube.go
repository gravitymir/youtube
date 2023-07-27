package youtube

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Details struct {
	VideoID          string   `json:"videoId"`
	Title            string   `json:"title"`
	LengthSeconds    string   `json:"lengthSeconds"`
	Keywords         []string `json:"keywords"`
	ChannelID        string   `json:"channelId"`
	IsOwnerViewing   bool     `json:"isOwnerViewing"`
	ShortDescription string   `json:"shortDescription"`
	IsCrawlable      bool     `json:"isCrawlable"`
	Thumbnail        struct {
		Thumbnails []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	AllowRatings      bool   `json:"allowRatings"`
	ViewCount         string `json:"viewCount"`
	Author            string `json:"author"`
	IsPrivate         bool   `json:"isPrivate"`
	IsUnpluggedCorpus bool   `json:"isUnpluggedCorpus"`
	IsLiveContent     bool   `json:"isLiveContent"`
}
type Format struct {
	Itag             int    `json:"itag"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength,omitempty"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate,omitempty"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
}
type SubtitlesTrack struct {
	BaseURL        string
	Name           string
	VssID          string
	LanguageCode   string
	Kind           string
	IsTranslatable bool
}

type YouTube struct {
	VideoID string
	Status  string // OK ERROR LOGIN_REQUIRED LIVE_STREAM_OFFLINE UNPLAYABLE CONTENT_CHECK_REQUIRED
	Reason  string
	Details
	VideoAndAudio   []Format
	VideoFormats    []Format
	AudioFormats    []Format
	SubtitlesTracks []SubtitlesTrack
	Subtitles       struct {
		Text []struct {
			Text  string `xml:",chardata"`
			Start string `xml:"start,attr"`
			Dur   string `xml:"dur,attr"`
		} `xml:"text"`
	}
	responseToYouTubeBodyData []byte
	responseSubtitlesBodyData []byte
	countRequests             uint
}

const (
	sleepDelayIfStatusCode403 = 1000
	countRequestsConst        = 4
	baseurl                   = "https://www.youtube.com/watch?v="
	//embedurl                   = "https://www.youtube.com/embed/"
	regexInitialPlayerResponse = `\{\"responseContext\".*?\"nanos\"\:[\d]+[\}]{4}`
)

// GetInfo init YouTube struct
// make Get request to web page YouTube
// with video ID
func GetInfo(ID string) (yt *YouTube, err error) {
	yt = &YouTube{}
	regex := regexp.MustCompile(`[a-zA-Z0-9-_]{11}`)
	if regex.Match([]byte(ID)) != true {
		return nil, errors.New("error: GetInfo() ID is doesn't match([a-zA-Z0-9-_]{11})")
	}
	ID = string(regex.Find([]byte(ID)))
	link := fmt.Sprintf("%s%s%s", baseurl, ID, "&hl=en")
	if err = yt.getRequestToYouTube(link); err != nil {
		return nil, err
	}
	if err = yt.convertResponseToYouTubeBodyDataToStruct(); err != nil {
		return nil, err
	}
	yt.countRequests = countRequestsConst
	return yt, nil
}
func (yt *YouTube) getRequestToYouTube(URL string) error {
	res, err := http.Get(URL)
	defer readCloser(res.Body)

	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("error getRequestToYouTube StatusCode is %d", res.StatusCode))
	}
	yt.responseToYouTubeBodyData, err = io.ReadAll(res.Body)
	return err
}

type writeCounter struct {
	ContentLength float64
	Part          float64
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Part += float64(n)
	percent := math.Floor((wc.Part / wc.ContentLength) * 100)
	downloads := strings.Repeat("█", int(percent/2)) //█ ▓ ▒ ░
	spaces := strings.Repeat("░", 50-int(percent/2)) //█ ▓ ▒ ░
	fmt.Printf("\rDownloading %s%s %.f%%", downloads, spaces, percent)
	return n, nil
}

func (yt *YouTube) Download(format Format, filePath string) (err error) {
	//regex := regexp.MustCompile(`(\bvideo\b|\baudio\b/)(\bwebm\b|\bmp4\b)`)
	regex := regexp.MustCompile(`\bwebm\b|\bmp4\b|\b3gp\b`)
	if regex.Match([]byte(format.MimeType)) != true {
		return errors.New("error: Download () Format don't recognised")
	}
	fileExt := "." + string(regex.Find([]byte(format.MimeType)))

	if filePath == "" {
		filePath = yt.VideoID
	}
	out, err := os.Create(filePath + fileExt)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(out)
	resp, err := http.Get(format.URL)
	if err != nil {
		fmt.Println("error: http.Get(format.URL)")
		return err
	}
	defer readCloser(resp.Body)

	if resp.StatusCode == 403 {
		if yt.countRequests > 0 {
			time.Sleep(sleepDelayIfStatusCode403)
			yt.countRequests--
			fmt.Println("error: StatusCode is 403, soon next try")
			return yt.Download(format, filePath)
		}
		yt.countRequests = countRequestsConst
		return errors.New(
			fmt.Sprintf("error: Download() imposible StatusCode = %d", resp.StatusCode))
	} else {
		yt.countRequests = countRequestsConst
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: resp.StatusCode == %d", resp.StatusCode)
		return errors.New(
			fmt.Sprintf("error: resp.StatusCode != http.StatusOK: %d", resp.StatusCode))
	}

	contentLength, err := strconv.ParseFloat(resp.Header["Content-Length"][0], 64)
	if err != nil {
		fmt.Println("error: strconv.ParseFloat")
		return err
	}
	counter := &writeCounter{
		ContentLength: contentLength,
	}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	return err
}

// GetSubtitlesXML get subtitles into XML format
func (yt *YouTube) GetSubtitlesXML(track SubtitlesTrack) (subtitles string, err error) {
	if len(yt.SubtitlesTracks) == 0 {
		return "", errors.New("error: GetSubtitlesXML() SubtitlesTracks is empty")
	}
	if err = yt.getSubtitles(track); err != nil {
		return "", err
	}
	return string(yt.responseSubtitlesBodyData), err
}

// GetSubtitlesPlainText get subtitles into only plain text, without timestamps
func (yt *YouTube) GetSubtitlesPlainText(track SubtitlesTrack) (subtitles string, err error) {
	if len(yt.SubtitlesTracks) == 0 {
		return "", errors.New("error: GetSubtitles() SubtitlesTracks is empty")
	}
	if err = yt.getSubtitles(track); err != nil {
		return "", err
	}
	for _, v := range yt.Subtitles.Text {
		subtitles += fmt.Sprintf("%s\n", v.Text)
	}
	return strings.TrimRight(subtitles, "\n"), err
}

// GetSubtitlesJson subtitles in string JSON
func (yt *YouTube) GetSubtitlesJson(track SubtitlesTrack) (subtitles string, err error) {
	if len(yt.SubtitlesTracks) == 0 {
		return "", errors.New("error: GetSubtitles() SubtitlesTracks is empty")
	}
	if err = yt.getSubtitles(track); err != nil {
		return "", err
	}
	if res, err := json.Marshal(yt.Subtitles); err != nil {
		return string(res), err
	}
	return subtitles, err
}

// GetSubtitlesJsonPretty return subtitles in string JSON with 4 spaces for better humans reading
// pretty style JSON print
func (yt *YouTube) GetSubtitlesJsonPretty(track SubtitlesTrack) (subtitles string, err error) {
	if len(yt.SubtitlesTracks) == 0 {
		return "", errors.New("error: GetSubtitles() SubtitlesTracks is empty")
	}
	if err = yt.getSubtitles(track); err != nil {
		return "", err
	}
	s, err := json.MarshalIndent(yt.Subtitles, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), err
}

// SaveSubtitlesJsonPretty save subtitles to file in pretty format
func (yt *YouTube) SaveSubtitlesJsonPretty(track SubtitlesTrack, filePath string) (err error) {
	jsonString, err := yt.GetSubtitlesJsonPretty(track)
	if err != nil {
		return err
	}
	if filePath == "" {
		filePath = yt.VideoID
	}

	err = os.WriteFile(filePath+".json", []byte(jsonString), 777)
	return err
}

// SaveSubtitlesXML save subtitles to file in XML format, plain format
func (yt *YouTube) SaveSubtitlesXML(track SubtitlesTrack, filePath string) (err error) {
	err = yt.getSubtitles(track)
	if err != nil {
		return err
	}
	if filePath == "" {
		filePath = yt.VideoID
	}

	err = os.WriteFile(filePath+".xml", yt.responseSubtitlesBodyData, 777)
	return err
}

func (yt *YouTube) convertResponseToYouTubeBodyDataToStruct() (err error) {
	playerResponse := regexp.MustCompile(regexInitialPlayerResponse)
	if playerResponse.Match(yt.responseToYouTubeBodyData) != true {
		return errors.New("error parser: can't find \"regexInitialPlayerResponse\"\n" +
			"Could be YouTube change struct page\n" +
			"If it true, need modify github.com/gravitymir/youtube, string 945")

	}
	cs := new(struct {
		PlayabilityStatus struct {
			Status      string `json:"status"`
			ErrorScreen struct {
				PlayerErrorMessageRenderer struct {
					Reason struct {
						SimpleText string `json:"simpleText"`
					} `json:"reason"`
				} `json:"playerErrorMessageRenderer"`
			} `json:"errorScreen"`
		} `json:"playabilityStatus"`
		StreamingData struct {
			Formats         []Format `json:"formats"`
			AdaptiveFormats []Format `json:"adaptiveFormats"` //here
		} `json:"streamingData"`
		Captions struct {
			PlayerCaptionsTracklistRenderer struct {
				Captions []struct {
					BaseURL string `json:"baseUrl"`
					Name    struct {
						SimpleText string `json:"simpleText"`
					} `json:"name"`
					VssID          string `json:"vssId"`
					LanguageCode   string `json:"languageCode"`
					Kind           string `json:"kind"`
					IsTranslatable bool   `json:"isTranslatable"`
				} `json:"captionTracks"`
			} `json:"playerCaptionsTracklistRenderer"`
		} `json:"captions"`
		Details `json:"videoDetails"`
	})
	if err = json.Unmarshal(playerResponse.Find(yt.responseToYouTubeBodyData), cs); err != nil {
		return err
	}
	yt.Status = cs.PlayabilityStatus.Status
	yt.Reason = cs.PlayabilityStatus.ErrorScreen.PlayerErrorMessageRenderer.Reason.SimpleText
	yt.Details = cs.Details

	yt.VideoAndAudio = cs.StreamingData.Formats
	for _, v := range append(cs.StreamingData.AdaptiveFormats) {
		if v.Fps == 0 {
			yt.AudioFormats = append(yt.AudioFormats, v)
		} else {
			yt.VideoFormats = append(yt.VideoFormats, v)
		}
	}

	for _, v := range cs.Captions.PlayerCaptionsTracklistRenderer.Captions {
		yt.SubtitlesTracks = append(yt.SubtitlesTracks, SubtitlesTrack{
			BaseURL:        v.BaseURL,
			Name:           v.Name.SimpleText,
			VssID:          v.VssID,
			LanguageCode:   v.LanguageCode,
			Kind:           v.Kind,
			IsTranslatable: v.IsTranslatable,
		})
	}
	return err
}
func (yt *YouTube) getSubtitles(track SubtitlesTrack) (err error) {
	if len(yt.SubtitlesTracks) == 0 {
		return errors.New("error: GetSubtitles() SubtitlesTracks is empty")
	}
	resp, err := http.Get(track.BaseURL)
	defer readCloser(resp.Body)
	if err != nil {
		fmt.Println("error: GetSubtitles() http.Get(track.BaseURL) ")
		fmt.Println(err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("error: StatusCode %d != http.StatusOK", resp.StatusCode)
		return errors.New(err)
	}
	if yt.responseSubtitlesBodyData, err = io.ReadAll(resp.Body); err != nil{
		return err
	}
	return xml.Unmarshal(yt.responseSubtitlesBodyData, &yt.Subtitles)
}

func readCloser(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
