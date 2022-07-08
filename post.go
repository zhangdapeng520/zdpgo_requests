package zdpgo_requests

/*
@Time : 2022/6/6 14:07
@Author : 张大鹏
@File : post.go
@Software: Goland2021.3.1
@Description:
*/

func (r *Requests) PostAes(targetUrl, jsonStr string) (*Response, error) {
	// 加密数据
	encryptData, err := r.Password.Aes.Encrypt([]byte(jsonStr))
	if err != nil {
		return nil, err
	}

	// 发送请求
	response, err := r.Any("POST", targetUrl, string(encryptData))
	if err != nil {
		return nil, err
	}

	// AES解密响应数据
	decryptBytes, err := r.Password.Aes.Decrypt(response.Content)
	if err != nil {
		return nil, err
	}
	response.Content = decryptBytes
	response.Text = string(decryptBytes)

	// 返回响应
	return response, nil
}
