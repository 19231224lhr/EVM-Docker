# 数字签名

|       名称        | 已实现 | 来源                                                               | 描述              |
|:---------------:|:--:|:-----------------------------------------------------------------|-----------------|
|      ECDSA      |  √ | `github.com/btcsuite/btcd/btcec/ecdsa`                           | btc的go实现中的ecdsa |
|    ECSchnorr    |  √ | `github.com/BintaSong/go-ethereum-fork/crypto/ecschnorr`         | 网友实现            |
|       BLS       |  √ | `github.com/dfinity-side-projects/go-dfinity-crypto/bls`         | dfinity的相关项目    |
|      EdDSA      |  √ | `github.com/teserakt-io/golang-ed25519`                          | Teserakt AG的项目  |
|       SM2       |  √ | `github.com/tjfoc/gmsm/sm2`                                      | 同济区块链研究院项目      |
|     BLS多重签名   |  √ | `github.com/dfinity-side-projects/go-dfinity-crypto/groupsig`    | dfinity的相关项目    |
|  BLS门限签名      | √ | `github.com/dfinity-side-projects/go-dfinity-crypto/`    |  dfinity的相关项目    |
|  BLS(纯go)      |  √ | `github.com/chuwt/chia-bls-go`    |  来自chia-bls-go项目，引用了kilic/bls12-381库（曲线库）  |
| PS随机短签名 |  × | `github.com/Zhiyi-Zhang/PS-Signature-and-EL-PASSO`    | 网友实现            |
|      CLSAG      |  × | `github.com/monero-project/monero/src/ringct`    | monero的项目       |
|  BLS(较新)      |  × | `github.com/relab/hotstuff/crypto/bls12`    |  ResilientSystemsLab的项目    |
