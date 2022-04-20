package zdpgo_requests

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// DownloadToBytes 下载文件，返回文件流
func (r *Requests) DownloadToBytes(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Download 下载文件并保存到指定路径
func (r *Requests) Download(urlPath string, savePath string) error {
	resp, err := http.Get(urlPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
