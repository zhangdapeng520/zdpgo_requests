package zdpgo_requests

import (
	"io/ioutil"
	"testing"
)

func TestRequest_Upload(t *testing.T) {
	url := "http://localhost:8888/upload"
	response, err := r.Upload(url, "file", "test/test1.txt")
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}
}

func TestRequest_UploadByBytes(t *testing.T) {
	url := "http://localhost:8888/upload"
	fileContent, _ := ioutil.ReadFile("test/test1.txt")
	response, err := r.UploadByBytes(url, "file", "test1.txt", fileContent)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}
}
