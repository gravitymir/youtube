package youtube

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type CaptionTrack struct {
	BaseURL string `json:"baseUrl"`
	Name    struct {
		SimpleText string `json:"simpleText"`
	} `json:"name"`
	VssID          string `json:"vssId"`
	LanguageCode   string `json:"languageCode"`
	Kind           string `json:"kind"`
	IsTranslatable bool   `json:"isTranslatable"`
}

func (yt *YouTube) getCaptions(track CaptionTrack) (err error) {
	if len(yt.Video.CaptionTracks) == 0 {
		return errors.New("error: GetCaptions() CaptionTracks is empty")
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
	if yt.CaptionsDataByteSlice, err = io.ReadAll(resp.Body); err != nil {
		return err
	}
	return xml.Unmarshal(yt.CaptionsDataByteSlice, &yt.Captions)
}

// SaveCaptionsXML save Captions to file in XML format, plain format
func (yt *YouTube) SaveCaptionsXML(track CaptionTrack, filePath string) (err error) {
	err = yt.getCaptions(track)
	if err != nil {
		return err
	}
	if filePath == "" {
		filePath = yt.Video.ID
	}

	err = os.WriteFile(filePath, yt.CaptionsDataByteSlice, 0777)
	return err
}

// SaveCaptionsJsonPretty save Captions to file in pretty format
func (yt *YouTube) SaveCaptionsJsonPretty(track CaptionTrack, filePath string) (err error) {

	jsonString, err := yt.GetCaptionsJsonPretty(track)
	if err != nil {
		return err
	}
	if filePath == "" {
		return errors.New("error: SaveCaptionsJsonPretty filePath is empty")
	}

	err = os.WriteFile(filePath, []byte(jsonString), 0777)
	return err
}

// GetCaptionsJsonPretty return Captions in string JSON with 4 spaces for better humans reading
// pretty style JSON print
func (yt *YouTube) GetCaptionsJsonPretty(track CaptionTrack) (Captions string, err error) {
	if len(yt.Video.CaptionTracks) == 0 {
		return "", errors.New("error: GetCaptions() CaptionTracks is empty")
	}
	if err = yt.getCaptions(track); err != nil {
		return "", err
	}
	s, err := json.MarshalIndent(yt.Captions, "", "    ")
	if err != nil {
		return "", err
	}
	return string(s), err
}

// GetCaptionsJson Captions in string JSON
func (yt *YouTube) GetCaptionsJson(track CaptionTrack) (Captions string, err error) {
	if len(yt.Video.CaptionTracks) == 0 {
		return "", errors.New("error: GetCaptions() CaptionTracks is empty")
	}
	if err = yt.getCaptions(track); err != nil {
		return "", err
	}
	if res, err := json.Marshal(yt.Captions); err != nil {
		return string(res), err
	}
	return Captions, err
}

// GetCaptionsPlainText get Captions into only plain text, without timestamps
func (yt *YouTube) GetCaptionsPlainText(track CaptionTrack) (Captions string, err error) {
	if len(yt.Video.CaptionTracks) == 0 {
		return "", errors.New("error: GetCaptions() CaptionTracks is empty")
	}
	if err = yt.getCaptions(track); err != nil {
		return "", err
	}
	for _, v := range yt.Captions.Text {
		Captions += fmt.Sprintf("%s\n", v.Text)
	}
	return strings.TrimRight(Captions, "\n"), err
}

// GetCaptionsXML get Captions into XML format
func (yt *YouTube) GetCaptionsXML(track CaptionTrack) (Captions string, err error) {
	if len(yt.Video.CaptionTracks) == 0 {
		return "", errors.New("error: GetCaptionsXML() CaptionTracks is empty")
	}
	if err = yt.getCaptions(track); err != nil {
		return "", err
	}
	return string(yt.CaptionsDataByteSlice), err
}
