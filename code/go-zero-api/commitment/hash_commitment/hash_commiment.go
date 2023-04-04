package hash_commitment

import (
	"crypto/sha256"
)

//对初始值s和混淆因子r的哈希承诺commit。
func Commit(s,r string) (commit [32]byte) {
	str:=s+r
	commit = sha256.Sum256([]byte(str))
	return commit
}

//验证承诺。
func Verify(s,r string, commit [32]byte)(bool){
	return Commit(s,r) == commit
}
