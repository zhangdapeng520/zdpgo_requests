package zdpgo_requests

import "testing"

// 测试文件上传
func TestRequests_Upload(t *testing.T) {
	targetUrl := "http://localhost:8888/upload"
	filename := "README.md"

	r := getRequests()
	err := r.Upload(targetUrl, filename)
	if err != nil {
		t.Error(err)
	}
}

func TestRequests_UploadToBytes(t *testing.T) {
	targetUrl := "http://localhost:8888/upload"
	filename := "README.md"

	r := getRequests()
	respBytes, err := r.UploadToBytes(targetUrl, filename)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(respBytes))
}
