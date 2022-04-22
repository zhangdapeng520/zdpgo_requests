package zdpgo_requests

import (
	"embed"
	"testing"
)

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

//go:embed examples/test/*
var fsObj embed.FS

func TestRequests_UploadFsToBytes(t *testing.T) {
	targetUrl := "http://localhost:8888/upload"
	filename := "examples/test/main.go"

	r := getRequests()
	respBytes, err := r.UploadFsToBytes(targetUrl, fsObj, "file", filename)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(respBytes))
}

func TestRequests_UploadFsToString(t *testing.T) {
	targetUrl := "http://localhost:8888/upload"
	filename := "examples/test/main.go"

	r := getRequests()
	respBytes, err := r.UploadFsToString(targetUrl, fsObj, "file", filename)
	if err != nil {
		t.Error(err)
	}
	t.Log(respBytes)
}
