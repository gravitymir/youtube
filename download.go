package youtube

import (
	"fmt"
	kkdai "github.com/kkdai/youtube/v2"
	"os"
	"os/exec"
)

type Formats = kkdai.FormatList

func GetFormats(videoID string) (rormats Formats, err error) {
	client := kkdai.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return nil, err
	}

	return video.Formats, nil
}

func download(ID, folderName, fileName, qualityOrTAG string) error {
	fmt.Printf("⬇ %s %s %s\n", qualityOrTAG, folderName, fileName)
	//fmt.Printf("youtubedr download -q %s -o %s %s",
	//	qualityOrTAG,
	//	fmt.Sprintf("../media/%s/%s", folderName, fileName),
	//	fmt.Sprintf("\"%s\"", ID))
	cmd := exec.Command("youtubedr", "download",
		"-q", qualityOrTAG,
		"-o", fmt.Sprintf("../media/%s/%s", folderName, fileName),
		fmt.Sprintf("\"%s\"", ID))

	//  fmt.Printf("youtubedr download -q %s -o %s %s",
	//	fmt.Sprintf("\"%s\"", qualityOrTAG),
	//	fmt.Sprintf("'../media/%s/%s'", folderName, fileName),
	//	fmt.Sprintf("\"%s\"", ID),
	//)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s %s download is done\n\n",
		qualityOrTAG,
		folderName,
		fileName,
	)
	return nil
}

// DownloadVideo first choice
func DownloadVideo(ID, folderName string) (err error) {
	tags := []string{
		"313", // 2160x3840 video/webm;
		"271", // 1440x2560 video/webm;
		"248", // 1080x1920 video/webm;
		"303", // 1080x1920 video/webm;
		"247", // 720x1280 video/webm;
		"302", // 720x1280 video/webm;
		"244", // 480x854 video/webm;
		"243", // 360x640 video/webm;
		"242", // 240x426 video/webm;
		"278", // 144x256 video/webm;
	}
	for k := range tags {
		if err = download(
			ID,
			folderName,
			fmt.Sprintf("%s.webm", tags[k]),
			tags[k],
		); err == nil {
			break
		}
	}
	return err
}

// DownloadAudio first choice
func DownloadAudio(ID, folderName string) (err error) {

	tags := []string{
		"251", //audio/webm; codecs="opus"
		"250", //audio/webm; codecs="opus"
		"249", //audio/webm; codecs="opus"
	}
	for k := range tags {
		if err = download(
			ID,
			folderName,
			fmt.Sprintf("%s.webm", tags[k]),
			tags[k],
		); err == nil {
			break
		}
	}
	return err
}

//MERGE AUDIO and Video
//ffmpeg -i 278.webm -i 249.webm -c:v copy -map 0:v  -map 1:a   278and249.webm
