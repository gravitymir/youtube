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
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	VideoID          string   `json:"videoId"`
	Title            string   `json:"title"`
	LengthSeconds    string   `json:"lengthSeconds"`
	Keywords         []string `json:"keywords"`
	ShortDescription string   `json:"shortDescription"`
	Author           string   `json:"author"`
}
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
	baseurl = "https://www.youtube.com/watch?v="
	//embedurl                   = "https://www.youtube.com/embed/"
	regexInitialPlayerResponse = `\{\"responseContext\".*?\"nanos\"\:[\d]+[\}]{4}`
	lengthFilepathExtPrint     = 56
)

// Init YouTube struct
// make Get request to web page YouTube
// with video ID
func Init(linkOrId string) (yt *YouTube, err error) {
	yt = &YouTube{}
	regex, err := regexp.Compile(`[a-zA-Z0-9-_]{11}`)
	if regex.Match([]byte(linkOrId)) != true {
		return nil, err
	}
	yt.VideoID = string(regex.Find([]byte(linkOrId)))
	link := fmt.Sprintf("%s%s%s", baseurl, yt.VideoID, "&hl=en")
	if err = yt.getRequestToYouTube(link); err != nil {
		return nil, err
	}
	if err = yt.convertResponseToYouTubeBodyDataToStruct(); err != nil {
		return nil, err
	}
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
	contentLength          int
	humanizedContentLength string
	part                   int
	currentChunk           int
	timeStart              time.Time
	videoID                string
	filePathExt            string
	filePathExtForPrint    string
}

// Exec d
func Exec() error {
	cmd := exec.Command("go", "run", "download2.go")
	if err := cmd.Start(); err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 2)
	//stderr, err := cmd.StderrPipe()
	err = cmd.Process.Signal(os.Interrupt)
	if err != nil {
		return err
	}

	out := make([]byte, 1024)

	// infinite loop to listen to stdin from that child process
	// it means what ever is printed using fmt.Print in that golang application will be captured here
	for {

		// reading the bytes
		n, err := stdout.Read(out)
		if err != nil {
			return err
		}
		fmt.Println(string(out[:n]))
	}

	// loop exited means process either crashed or stopped
	//fmt.Println("Process crashed or stopped...")
	//
	//cmd.Wait()
}

// Download try download2 file from YouTube
func (yt *YouTube) Download(format Format, filePath string) (err error) {
	req, err := http.NewRequest("GET", format.URL, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp, err := http.Get(format.URL)
	if err != nil {
		return err
	}
	defer readCloser(resp.Body)
	if resp.StatusCode == 403 {
		return errors.New(
			fmt.Sprintf("error: Download() imposible StatusCode = %d", resp.StatusCode))
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: resp.StatusCode == %d", resp.StatusCode)
		return errors.New(
			fmt.Sprintf("error: resp.StatusCode != http.StatusOK: %d", resp.StatusCode))
	}
	contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

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
	file := filePath + fileExt
	if len(file) > lengthFilepathExtPrint {
		file = file[len(file)-lengthFilepathExtPrint:]
	} else {
		file = strings.Repeat(" ", lengthFilepathExtPrint-len(file)) + file
	}

	counter := &writeCounter{
		contentLength:          contentLength,
		humanizedContentLength: humanizeContentLength(contentLength),
		timeStart:              time.Now(),
		videoID:                yt.VideoID,
		filePathExt:            filePath + fileExt,
		filePathExtForPrint:    file,
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

	err = os.WriteFile(filePath, []byte(jsonString), 777)
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

	err = os.WriteFile(filePath, yt.ResponseSubtitlesBodyData, 777)
	return err
}

// GetDetailsPretty g
func (yt *YouTube) GetDetailsPretty() (details []byte, err error) {
	return json.MarshalIndent(yt.Details, "", "    ")
}

// SaveDetailsPretty g
func (yt *YouTube) SaveDetailsPretty(filePath string) (err error) {
	if filePath == "" {
		filePath = yt.VideoID + ".json"
	}
	details, err := yt.GetDetailsPretty()
	err = os.WriteFile(filePath, details, 777)
	return err
}

// SaveThumbnail g
func (yt *YouTube) SaveThumbnailJPG(filePath string) (err error) {
	if filePath == "" {
		filePath = yt.VideoID + ".jpg"
	}

	thumb := yt.Details.Thumbnail.Thumbnails[len(yt.Details.Thumbnail.Thumbnails)-1]
	resp, err := http.Get(thumb.URL)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("SaveThumbnailJPG %e", err)
		}
	}(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("StatusCode is: " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, []byte(data), 777)
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

	str := fmt.Sprintf("%s %s %s %s %s %s",
		wc.videoID,
		fmt.Sprintf("%3.f%%", percent),
		humanizeContentLength(wc.part),
		fmtDuration(time.Now().Sub(wc.timeStart)),
		wc.humanizedContentLength,
		wc.filePathExtForPrint,
	)

	fmt.Printf("\r%s",
		decorateDownloadLine(
			str[:int(percent)],
			str[int(percent):],
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
func decorateDownloadLine(downloaded, willDownload string) string {
	return fmt.Sprintf("%s%s",
		color.New(color.ReverseVideo).SprintFunc()(downloaded),
		color.New(color.BgWhite, color.Bold).SprintFunc()(willDownload),
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
