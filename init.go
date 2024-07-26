package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// InitVideo YouTube struct
// make Get request to web page YouTube
// with video ID
func InitVideo(linkOrId string) (yt *YouTube, err error) {
	regex, err := regexp.Compile(`[a-zA-Z0-9-_]{11}`)
	if !regex.MatchString(linkOrId) { //Match on video URL
		return nil, err
	}
	yt = &YouTube{}
	yt.Video.ID = regex.FindString(linkOrId)
	link := fmt.Sprintf("https://www.youtube.com/watch?v=%s%s", yt.Video.ID, "&hl=en")
	if err = yt.requestToVideo(link); err != nil {
		return nil, err
	}
	if err = yt.convertVideoToStruct(); err != nil {
		return nil, err
	}
	return yt, nil
}

func (yt *YouTube) requestToVideo(URL string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error requestToVideo StatusCode is %d", res.StatusCode)
	}
	defer readCloser(res.Body)
	yt.Video.VideoDataByteSlice, err = io.ReadAll(res.Body)
	return err
}

func (yt *YouTube) convertVideoToStruct() (err error) {
	//`\{\"responseContext\".*?\"nanos\"\:[\d]+[\}]{4}`
	playerResponse := regexp.MustCompile(`\{"responseContext".*?"nanos":[\d]+[}]{4}`)
	if !playerResponse.Match(yt.Video.VideoDataByteSlice) {
		return errors.New("error parser: can't find regexInitialPlayerRespons\n" +
			"Could be YouTube change struct page\n" +
			"If it true, need modify github.com/gravitymir/youtube, string 945")

	}
	cs := &VideoStruct{}
	if err = json.Unmarshal(playerResponse.Find(yt.Video.VideoDataByteSlice), cs); err != nil {
		return err
	}

	yt.Video.Status = cs.PlayabilityStatus.Status
	yt.Video.Reason = cs.PlayabilityStatus.ErrorScreen.PlayerErrorMessageRenderer.Reason.SimpleText

	yt.Video.Details = FullDetails{
		VideoDetails: cs.VideoDetails,
		PublishDate:  cs.Microformat.PlayerMicroformatRenderer.PublishDate,
		UploadDate:   cs.Microformat.PlayerMicroformatRenderer.UploadDate,
	}

	yt.Video.VideoAndAudio = cs.StreamingData.Formats

	for _, v := range cs.StreamingData.AdaptiveFormats {
		if v.Fps == 0 {
			yt.Video.AudioFormats = append(yt.Video.AudioFormats, v)
		} else {
			yt.Video.VideoFormats = append(yt.Video.VideoFormats, v)
		}
	}
	yt.Video.CaptionTracks = cs.Captions.PlayerCaptionsTracklistRenderer.CaptionTracks

	return err
}

func readCloser(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
