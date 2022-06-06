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
	response, err := r.AnyText(Request{
		Method: "POST",
		Url:    targetUrl,
		Text:   base64.StdEncoding.EncodeToString(encryptData),
	})
	if err != nil {
		r.Log.Error("发送JSON请求失败", "error", err)
		return nil, err
	}

	// 返回响应
	return response, nil
}
