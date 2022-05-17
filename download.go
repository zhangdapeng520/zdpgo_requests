package zdpgo_requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

// DownloadToBytes 下载文件，返回文件流
func (r *Requests) DownloadToBytes(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		r.Log.Error("获取下载数据失败", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error("读取下载数据失败", "error", err)
		return nil, err
	}
	return data, nil
}

// Download 下载文件并保存到指定路径
func (r *Requests) Download(urlPath, saveDir string) {
	resp, err := http.Get(urlPath)
	if err != nil {
		r.Log.Error("获取下载数据失败", "error", err, "urlPath", urlPath)
		return
	}
	defer resp.Body.Close()

	// 创建保存文件
	fileName := r.File.GetFileName(urlPath)
	savePath := path.Join(saveDir, fileName)
	out, err := os.Create(savePath)
	if err != nil {
		r.Log.Error("创建保存文件失败", "error", err, "savePath", savePath)
		return
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		r.Log.Error("保存下载数据失败", "error", err)
		return
	}
}

// DownloadMany 批量下载数据
func (r *Requests) DownloadMany(urlPath []string, saveDir string) {
	var wg = new(sync.WaitGroup)
	for _, url := range urlPath {
		r.Log.Debug("正在下载", "url", url)
		savePath := r.File.GetFileName(url)
		if savePath == "" {
			r.Log.Error("获取文件名失败", "savePath", savePath, "url", url)
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, url, savePath string) {
			r.Download(url, saveDir)
			wg.Done()
		}(wg, url, savePath)
	}
	wg.Wait()
}

// DownloadToTmpAndReturnIsDeleted 下载文件到临时目录，等待指定时间（秒）以后判断依然存在并删除
// @param urlPath 下载文件路径
// @param waitSeconds 等待时间，单位秒
// @return bool 是否被删除
func (r *Requests) DownloadToTmpAndReturnIsDeleted(urlPath string, waitSeconds int) bool {
	tmpFileName := r.DownloadToTmp(urlPath)

	// 等待一段时间
	time.Sleep(time.Duration(waitSeconds) * time.Second)

	// 判断数据是否存在
	flag := r.File.IsDirContainsFile("tmp", tmpFileName)

	// 删除临时文件
	r.File.DeleteDirFile("tmp", tmpFileName)
	return flag
}

// DownloadToTmp 下载文件到临时目录
func (r *Requests) DownloadToTmp(urlPath string) string {
	// 获取文件数据
	resp, err := http.Get(urlPath)
	if err != nil {
		r.Log.Error("获取文件数据失败", "error", err, "url", urlPath)
		return ""
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// 获取文件后缀
	suffix := r.File.GetFileSuffix(urlPath)

	// 读取文件数据
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Log.Error("读取文件数据失败", "error", err, "resp", resp, "body", resp.Body)
		return ""
	}

	// 将数据写入临时文件
	if !r.File.IsExists("tmp") {
		r.File.CreateMultiDir("tmp")
	}
	tmpFileName := fmt.Sprintf("%s%s", r.Random.Str(32), suffix)
	r.File.CreateFile(path.Join("tmp", tmpFileName), string(data))

	// 返回临时文件名
	return tmpFileName
}
