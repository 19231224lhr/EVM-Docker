package main

import (
	"blockchain-crypto/vrf"
	"fmt"
)

func main() {

	//生成公私钥对(priv,pub)
	priv, _ := vrf.Newprivatekey()
	pub, _ := vrf.GeneratePublickey(priv)

	//利用消息message和私钥priv获得vrf结果vrf0以及证明proof
	message := []byte("heloo")
	vrf0, proof, _ := vrf.Prove(priv, message)

	//利用公钥pub，消息message，证明proof 验证vrf结果vrf0的正确性
	res, _ := vrf.Verify(pub, message, vrf0, proof)
	fmt.Println(res)
}
