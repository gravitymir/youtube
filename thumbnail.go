package youtube

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func (yt *YouTube) SaveThumbnail(folderPath string) (err error) {
	thumb := yt.Video.Details.Thumbnail.Thumbnails[len(yt.Video.Details.Thumbnail.Thumbnails)-1]
	resp, err := http.Get(thumb.URL)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("SaveThumbnailJPG %s", err)
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
	err = os.WriteFile(path.Join(
		folderPath,
		"thumbnail.jpg",
	), data, 0777)
	return err
}
