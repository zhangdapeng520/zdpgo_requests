package zdpgo_requests

import "testing"

func TestRequests_DownloadToBytes(t *testing.T) {
	r := getRequests()
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	dataBytes, err := r.DownloadToBytes(imgUrl)
	if err != nil {
		t.Error(err)
	}
	t.Log(dataBytes)
}

func TestRequests_Download(t *testing.T) {
	r := getRequests()
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	err := r.Download(imgUrl, "test.jpg")
	if err != nil {
		t.Error(err)
	}
}
