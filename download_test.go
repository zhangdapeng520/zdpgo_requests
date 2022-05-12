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
	r.Download(imgUrl, "tmp")
}

func TestRequests_DownloadToTmp(t *testing.T) {
	r := getRequests()
	data := []struct {
		Url       string
		NotResult string
	}{
		{"https://www.twle.cn/static/i/img1.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
		{"https://images3.alphacoders.com/120/1205462.jpg", ""},
	}

	for _, url := range data {
		result := r.DownloadToTmp(url.Url)
		if result == url.NotResult {
			panic("下载错误：不是期望的值")
		}
	}
}

func TestRequests_DownloadToTmpAndReturnIsDeleted(t *testing.T) {
	r := getRequests()
	data := []struct {
		Url       string
		IsDeleted bool
	}{
		{"https://www.twle.cn/static/i/img1.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
		{"https://images3.alphacoders.com/120/1205462.jpg", false},
	}

	for _, url := range data {
		result := r.DownloadToTmpAndReturnIsDeleted(url.Url, 10)
		if result == url.IsDeleted {
			panic("下载错误：不是期望的值")
		}
	}
}

func TestRequests_DownloadMany(t *testing.T) {
	r := getRequests()
	data := [][]string{
		{"https://www.twle.cn/static/i/img1.jpg", "https://images3.alphacoders.com/120/1205462.jpg"},
	}

	for _, urls := range data {
		r.DownloadMany(urls, "tmp")
	}
}
