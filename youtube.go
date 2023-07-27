package youtube

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/fatih/color"
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
	ResponseToYouTubeBodyData []byte
	ResponseSubtitlesBodyData []byte
	countRequests             uint
}

const (
	secondsSleepDelayIfStatusCode403 = 1
	countRequestsConst               = 4
	baseurl                          = "https://www.youtube.com/watch?v="
	//embedurl                   = "https://www.youtube.com/embed/"
	regexInitialPlayerResponse = `\{\"responseContext\".*?\"nanos\"\:[\d]+[\}]{4}`
)

// Init YouTube struct
// make Get request to web page YouTube
// with video ID
func Init(linkOrId string) (yt *YouTube, err error) {
	yt = &YouTube{}
	regex := regexp.MustCompile(`[a-zA-Z0-9-_]{11}`)
	if regex.Match([]byte(linkOrId)) != true {
		return nil, errors.New("error: Init() ID is doesn't match([a-zA-Z0-9-_]{11})")
	}
	yt.VideoID = string(regex.Find([]byte(linkOrId)))
	link := fmt.Sprintf("%s%s%s", baseurl, yt.VideoID, "&hl=en")
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
	yt.ResponseToYouTubeBodyData, err = io.ReadAll(res.Body)
	return err
}

type writeCounter struct {
	contentLength int
	part          int
	currentChunk  int
	timeStart     time.Time
	videoID       string
}

// Download try download file from YouTube
func (yt *YouTube) Download(format Format, filePath string) (err error) {
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
			time.Sleep(secondsSleepDelayIfStatusCode403 * time.Second)
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
	contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))

	if err != nil {
		fmt.Println("error: strconv.ParseFloat")
		return err
	}
	if contentLength == 0 {
		return errors.New(
			fmt.Sprintf("error: resp.Header.Get(\"Content-Length\") %d", contentLength))

	}
	counter := &writeCounter{
		contentLength: contentLength,
		timeStart:     time.Now(),
		videoID:       yt.VideoID,
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
	return string(yt.ResponseSubtitlesBodyData), err
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

	err = os.WriteFile(filePath+".xml", yt.ResponseSubtitlesBodyData, 777)
	return err
}
func (yt *YouTube) convertResponseToYouTubeBodyDataToStruct() (err error) {
	playerResponse := regexp.MustCompile(regexInitialPlayerResponse)
	if playerResponse.Match(yt.ResponseToYouTubeBodyData) != true {
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
	if err = json.Unmarshal(playerResponse.Find(yt.ResponseToYouTubeBodyData), cs); err != nil {
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
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("error: StatusCode %d != http.StatusOK", resp.StatusCode)
		return errors.New(err)
	}
	if yt.ResponseSubtitlesBodyData, err = io.ReadAll(resp.Body); err != nil {
		return err
	}
	return xml.Unmarshal(yt.ResponseSubtitlesBodyData, &yt.Subtitles)
}
func readCloser(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	amountChunks := wc.contentLength / n
	wc.part += n
	currentChunk := wc.part / n
	aux := (float64(currentChunk) / float64(amountChunks)) * 100
	percent := math.Floor(aux) //█ ▓ ▒ ░

	elapsedTime := float64(time.Now().Sub(wc.timeStart))
	chunksPerTime := float64(currentChunk) / elapsedTime
	estimatedTotalTime := float64(amountChunks) / chunksPerTime
	timeLeftInSeconds := (estimatedTotalTime - elapsedTime)

	str := fmt.Sprintf("%s %s %s%s%s%s",
		wc.videoID,
		fmtDuration(time.Now().Sub(wc.timeStart)),
		fmtDuration(time.Duration(uint64(timeLeftInSeconds))),
		humanizeContentLength(wc.contentLength),
		humanizeContentLength(wc.part),
		fmt.Sprintf("%4.f%%", percent))

	fmt.Printf("\r%s",
		decorateDownloadLine(
			str[:int(percent/2)],
			str[int(percent/2):],
		),
	)
	//fmt.Println(utf8.RuneCountInString(str))
	if percent == float64(100) {
		fmt.Println()
	}
	return n, nil
}
func humanizeContentLength(b int) string {
	const unit = 1024

	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%6.1f%cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
func decorateDownloadLine(revers, white string) string {
	return fmt.Sprintf("%s%s",
		color.New(color.ReverseVideo).SprintFunc()(revers),
		color.New(color.BgWhite, color.Bold).SprintFunc()(white),
	)
}
func fmtDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
