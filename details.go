package youtube

import (
	"encoding/json"
	"os"
	"path"
)

type VideoDetails struct {
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
	AllowRatings           bool   `json:"allowRatings"`
	ViewCount              string `json:"viewCount"`
	Author                 string `json:"author"`
	IsLowLatencyLiveStream bool   `json:"isLowLatencyLiveStream"`
	IsPrivate              bool   `json:"isPrivate"`
	IsUnpluggedCorpus      bool   `json:"isUnpluggedCorpus"`
	LatencyClass           string `json:"latencyClass"`
	IsLiveContent          bool   `json:"isLiveContent"`
}
type FullDetails struct {
	VideoDetails
	PublishDate string `json:"publishDate"`
	UploadDate  string `json:"uploadDate"`
}

// GetDetailsPretty g
func (yt *YouTube) GetDetailsPretty() (details []byte, err error) {
	return json.MarshalIndent(yt.Video.Details, "", "    ")
}

// SaveDetailsPretty g
func (yt *YouTube) SaveDetailsPretty(filePath string) (err error) {
	details, err := yt.GetDetailsPretty()
	if err != nil {
		return err
	}
	err = os.WriteFile(
		path.Join(filePath, "details.json"),
		details,
		0777,
	)
	return err
}
