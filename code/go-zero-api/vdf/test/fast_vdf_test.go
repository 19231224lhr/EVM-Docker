package main

import (
	"blockchain-crypto/vdf/wesolowski_go"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2047(t *testing.T) {
	input, _ := hex.DecodeString("ed8c9fa1c04979a2d39e4c91b97fafa7325b66fda3d06a14af4500d113d2fec0")
	iteration := 10       /* any value */
	int_size_bits := 2049 /* most of value not power of two */
	y, proof := wesolowski_go.GenerateVDF(input[:], iteration, int_size_bits)
	output := append(y, proof...)
	valid := wesolowski_go.VerifyVDF(input[:], output, iteration, int_size_bits)
	assert.Equal(t, true, valid, "must be true")
}

func Test1000(t *testing.T) {
	seed := []byte{0xde, 0xad, 0xbe, 0xef}
	ref_output := "0011c26e62c608dba629ce37953b2c7765f6c3c48f58ae5dc6ebc19206ca3135f8a240538a42f989d990488185f2d10a2504838f4f2e4dd933119088aa0e5b506bfd835d03147b03d5111e6ca135bf297435faf27a8ccbdb0c7598934fdccde6c9afbdc0488662618fc3934eaa9913f97559fb119ff959fc5f35a71da783c64af0000461c617ca4fd50ab15bf8c62963b043e1920b619402aa11a7fb82b793e9fca643bb8b8026e09493e6ed0f69ad7dafef7938f7c78d7067247d43ce2cf73174ffd78d2d4107a0421cf16a7fb118978b4903425bb84dcc4d0102267103494b798247cabc65caff373c368530fb7d869317d86a279eb55facf75a430109b5343875003fc63ce964ad0fc804687fb21b9322d672299cf0eeb53f5f426a4123e44db2ca593b50c026e54c079cd79634cd3969941aba18edae5fb51792776a2ed9076c79a456bb783ad87cdad013ce8a933c0c1a787a0232205dfa34b8ab65c1bd06f4004a3ffd5aec0c9cabedd081228c0b8c59e2bc2487f1fa2344288a8e9d7eefd169003fb7e55b707e9c5d76c84fa510647ec6c392f19f1a4ce98e71a601c1ee2479a93e5b9e4512b7c4cffc18b3498a36334e49db29aa56d487dec7b9dcc3128f722888903f10fd468a62ebb599eaced4114e36df647cc60b16c15f33b2f1c96bfe7c33d274ca57a456448ac7ac5539d10d71f72c0561d0460b9ab8f338898c8b8203"

	y_buf, proof_buf := wesolowski_go.GenerateVDF(seed, 1000, 2048)
	generated := hex.EncodeToString(append(y_buf, proof_buf...))

	assert.Equal(t, ref_output, generated, "they should be equal")
}

func Test5000(t *testing.T) {
	seed := []byte{0xde, 0xad, 0xbe, 0xef}
	ref_output := "00624b507eae90e9ac51a8ee427a0195e843642992cc34fa2c42a43ab75ec098a4d85e64b45fa5e48f3241495761a8af35e6e17d1259ea6d2b173c58c925c47a64eedfadccdff41841e20dd8a50c8a8dd87f0a0160e4f68f548f2a8ec6994bc826de9fee7e3fcd9d526808667327bf5f4cefd905117932b1cb737d74853f89206cffe2d0bd37520dcced47fd18a0e894d200c7d583ee7d55d84be5332a9c21bb6726e6ae64f294726a6e88be77ccd10e4139ed3c6309cfa2763b958149ff948233db76fc05a7abdfb08411e3f5a6567ab9247ebdbe310b0bb453340d385bdf6a4d67d454b1cdc2fbd82b12ae474bfe3038fd8fa8552f692a62df48dc208c18ac638d00748ead55bfc86ac0bd40e3038ea2d6e13cea20d3df4abf8923ada9927bacc531f35fcbf74080fdf5bb25f22b4186e941b8585f5ecc533b588dcf46b18e0871781a6cfe3b1f7f8dbc47369334ff7b32294453162a20d6c2d5711558cce36a9db0507e665922a8b99fab12145098f930b1c11afaf6f1e78256b536de1513b64d90ff8ba3b7b91d5989bb1aaf557063938df474e73d0568a51b2a5ad89b058f3adfaea137c9c1925feeb80b68aac4cb4335224efc365c81db004826c69b9659cd0a3e46bd4af5ccfa55bd53f3b7ddbae2a0802ada1076397cb6c1d7c9b6a4239db430c66fd90a24c21602723d0b895708d49cb92e8aeefc9a3760997b5b6760378f0b"

	y_buf, proof_buf := wesolowski_go.GenerateVDF(seed, 5000, 2048)
	generated := hex.EncodeToString(append(y_buf, proof_buf...))

	assert.Equal(t, ref_output, generated, "they should be equal")
}

func Test5000Verify(t *testing.T) {
	seed := []byte{0xde, 0xad, 0xbe, 0xef}
	ref_output := "00624b507eae90e9ac51a8ee427a0195e843642992cc34fa2c42a43ab75ec098a4d85e64b45fa5e48f3241495761a8af35e6e17d1259ea6d2b173c58c925c47a64eedfadccdff41841e20dd8a50c8a8dd87f0a0160e4f68f548f2a8ec6994bc826de9fee7e3fcd9d526808667327bf5f4cefd905117932b1cb737d74853f89206cffe2d0bd37520dcced47fd18a0e894d200c7d583ee7d55d84be5332a9c21bb6726e6ae64f294726a6e88be77ccd10e4139ed3c6309cfa2763b958149ff948233db76fc05a7abdfb08411e3f5a6567ab9247ebdbe310b0bb453340d385bdf6a4d67d454b1cdc2fbd82b12ae474bfe3038fd8fa8552f692a62df48dc208c18ac638d00748ead55bfc86ac0bd40e3038ea2d6e13cea20d3df4abf8923ada9927bacc531f35fcbf74080fdf5bb25f22b4186e941b8585f5ecc533b588dcf46b18e0871781a6cfe3b1f7f8dbc47369334ff7b32294453162a20d6c2d5711558cce36a9db0507e665922a8b99fab12145098f930b1c11afaf6f1e78256b536de1513b64d90ff8ba3b7b91d5989bb1aaf557063938df474e73d0568a51b2a5ad89b058f3adfaea137c9c1925feeb80b68aac4cb4335224efc365c81db004826c69b9659cd0a3e46bd4af5ccfa55bd53f3b7ddbae2a0802ada1076397cb6c1d7c9b6a4239db430c66fd90a24c21602723d0b895708d49cb92e8aeefc9a3760997b5b6760378f0b"

	buf, _ := hex.DecodeString(ref_output)
	assert.Equal(t, true, wesolowski_go.VerifyVDF(seed, buf, 5000, 2048), "must be true")
}
