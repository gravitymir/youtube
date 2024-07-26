package youtube

import (
	"fmt"
	"io"
	"net/http"

	kkdai "github.com/kkdai/youtube/v2"
)

type LineFromPlayList struct {
	ID    string
	Title string
}
type List []LineFromPlayList

func InitPlaylist(playListID string) (yt *YouTube, err error) {
	client := kkdai.Client{}
	playList, err := client.GetPlaylist(playListID) //"PLcihjUVySO7oupvXPjcvLQZoTX_QnFE52"
	if err != nil {
		return yt, err
	}
	yt = &YouTube{}
	yt.PlayList.List = List{}
	for _, v := range playList.Videos {
		fmt.Println(v.ID)
		yt.PlayList.List = append(yt.PlayList.List, LineFromPlayList{ID: v.ID, Title: v.Title})
	}

	return yt, err
}

func (yt *YouTube) RequestToPlayList(URL string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error requestToPlayList StatusCode is %d", res.StatusCode)
	}
	defer readCloser(res.Body)
	yt.PlayList.PlayListDataByteSlice, err = io.ReadAll(res.Body)
	return err
}
