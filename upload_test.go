package zdpgo_requests

import (
	"io/ioutil"
	"testing"
)

func TestRequest_Upload(t *testing.T) {
	url := "http://localhost:8888/upload"
	r.Upload(url, "file", "test/test1.txt")
	r.Log.Debug("响应结果", "response", r.Response.Text)
}

func TestRequest_UploadByBytes(t *testing.T) {
	url := "http://localhost:8888/upload"
	fileContent, _ := ioutil.ReadFile("test/test1.txt")
	r.UploadByBytes(url, "file", "test1.txt", fileContent)
	r.Log.Debug("响应结果", "response", r.Response.Text)
}
