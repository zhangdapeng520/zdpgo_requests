package zdpgo_requests

import "testing"

/*
@Time : 2022/6/6 15:07
@Author : 张大鹏
@File : post_test.go
@Software: Goland2021.3.1
@Description:
*/

func TestRequests_PostEcc(t *testing.T) {
	privateKey := `-----BEGIN  ZDPGO_PASSWORD ECC PRIVATE KEY -----
MHcCAQEEIKyfOnD7NdXudekftRtH2mBuOPf/UTzJ1Ulo2Hiu22XvoAoGCCqGSM49
AwEHoUQDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58cc9ybrYOo7t6bs818HMybbahM
Qylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PRIVATE KEY -----
`
	publicKey := `-----BEGIN  ZDPGO_PASSWORD ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58c
c9ybrYOo7t6bs818HMybbahMQylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PUBLIC KEY -----
`
	requests := NewWithConfig(&Config{
		Debug: true,
		Ecc: EccConfig{
			PrivateKey: []byte(privateKey),
			PublicKey:  []byte(publicKey),
		},
	})
	target := "http://127.0.0.1:3333/ecc"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"
	_, err := requests.PostEcc(target, jsonStr)
	if err != nil {
		panic(err)
	}
}

func TestRequests_PostEccText(t *testing.T) {
	privateKey := `-----BEGIN  ZDPGO_PASSWORD ECC PRIVATE KEY -----
MHcCAQEEIKyfOnD7NdXudekftRtH2mBuOPf/UTzJ1Ulo2Hiu22XvoAoGCCqGSM49
AwEHoUQDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58cc9ybrYOo7t6bs818HMybbahM
Qylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PRIVATE KEY -----
`
	publicKey := `-----BEGIN  ZDPGO_PASSWORD ECC PUBLIC KEY -----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXClGdjDvOFSHJzs2LtSfGcVzP58c
c9ybrYOo7t6bs818HMybbahMQylb+qB4aTtHV0JPqZAr8MChRmvze7nNFw==
-----END  ZDPGO_PASSWORD ECC PUBLIC KEY -----
`
	requests := NewWithConfig(&Config{
		Debug: true,
		Ecc: EccConfig{
			PrivateKey: []byte(privateKey),
			PublicKey:  []byte(publicKey),
		},
	})
	target := "http://127.0.0.1:3333/ecctext"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"
	_, err := requests.PostEccText(target, jsonStr)
	if err != nil {
		panic(err)
	}
}
