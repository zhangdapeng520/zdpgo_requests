package zdpgo_requests

import (
	"encoding/base64"

	"github.com/zhangdapeng520/zdpgo_password"
)

/*
@Time : 2022/6/6 14:07
@Author : 张大鹏
@File : post.go
@Software: Goland2021.3.1
@Description:
*/

var (
	ecc *zdpgo_password.Ecc
)

// PostEcc 发送POST请求的ECC加密数据
// @param targetUrl 目标地址
// @param jsonStr JSON格式的字符串
func (r *Requests) PostEcc(targetUrl, jsonStr string) (*Response, error) {
	// 获取ECC加密对象
	if ecc == nil {
		ecc = r.Password.GetEcc()
		privateKey, publicKey, err := ecc.GetKey()
		if err != nil {
			r.Log.Error("获取ECC加密私钥和公钥失败", "error", err)
			return nil, err
		}
		r.Config.Ecc.PrivateKey = privateKey
		r.Config.Ecc.PublicKey = publicKey
	}

	// 加密数据
	encryptData, err := ecc.EncryptByPublicKey([]byte(jsonStr), r.Config.Ecc.PublicKey)
	if err != nil {
		r.Log.Error("ECC加密数据失败", "error", err)
		return nil, err
	}

	// 发送请求
	response, err := r.Any("POST", targetUrl, r.GetText(base64.StdEncoding.EncodeToString(encryptData)))
	if err != nil {
		r.Log.Error("发送JSON请求失败", "error", err)
		return nil, err
	}

	// 返回响应
	return response, nil
}

func (r *Requests) PostAes(targetUrl, jsonStr string) (*Response, error) {
	// 加密数据
	encryptData, err := r.Password.Aes.Encrypt([]byte(jsonStr))
	if err != nil {
		r.Log.Error("AES加密数据失败", "error", err)
		return nil, err
	}

	// 发送请求
	response, err := r.Any("POST", targetUrl, string(encryptData))
	if err != nil {
		r.Log.Error("发送JSON请求失败", "error", err)
		return nil, err
	}

	// AES解密响应数据
	decryptBytes, err := r.Password.Aes.Decrypt(response.Content)
	if err != nil {
		r.Log.Error("AES解密响应数据失败", "error", err, "status", response.StatusCode, "content", response.Text)
		return nil, err
	}
	response.Content = decryptBytes
	response.Text = string(decryptBytes)

	// 返回响应
	return response, nil
}

// PostEccText POST提交ECC加密的纯文本
func (r *Requests) PostEccText(targetUrl, data string) (*Response, error) {
	// 获取ECC加密对象
	if ecc == nil {
		ecc = r.Password.GetEcc()
		privateKey, publicKey, err := ecc.GetKey()
		if err != nil {
			r.Log.Error("获取ECC加密私钥和公钥失败", "error", err)
			return nil, err
		}
		r.Config.Ecc.PrivateKey = privateKey
		r.Config.Ecc.PublicKey = publicKey
	}

	// 加密数据
	encryptData, err := ecc.EncryptByPublicKey([]byte(data), r.Config.Ecc.PublicKey)
	if err != nil {
		r.Log.Error("ECC加密数据失败", "error", err)
		return nil, err
	}

	// 发送请求
	response, err := r.Any("POST", targetUrl, string(encryptData))
	if err != nil {
		r.Log.Error("发送POST请求失败", "error", err)
		return nil, err
	}

	// ecc解密响应数据
	decryptBytest, err := ecc.Decrypt([]byte(response.Text))
	if err != nil {
		r.Log.Error("ECC解密响应数据失败", "error", err)
		return nil, err
	}
	response.Content = decryptBytest
	response.Text = string(decryptBytest)

	// 返回响应
	return response, nil
}
