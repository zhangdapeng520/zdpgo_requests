package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_requests"
	"reflect"
)

func main() {
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
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		Ecc: zdpgo_requests.EccConfig{
			PrivateKey: []byte(privateKey),
			PublicKey:  []byte(publicKey),
		},
	}, zdpgo_log.Tmp)
	target := "http://127.0.0.1:3333/ecctext"
	jsonStr := "{\"age\":22,\"username\":\"zhangdapeng\"}"
	response, err := requests.PostEccText(target, jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text, reflect.TypeOf(response.Text))
}
