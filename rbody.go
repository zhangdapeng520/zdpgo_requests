package zdpgo_requests

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

// 设置表单的字段
func (req *Request) setBodyBytes(Forms url.Values) {
	data := Forms.Encode()
	req.httpReq.Body = ioutil.NopCloser(strings.NewReader(data))
	req.httpReq.ContentLength = int64(len(data))
}

// 设置表单的二进制输入流
func (req *Request) setBodyRawBytes(read io.ReadCloser) {
	req.httpReq.Body = read
}
