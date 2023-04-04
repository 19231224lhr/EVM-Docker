package logic

import (
	"blockchain-crypto/crypto_api/vdf/internal/svc"
	"blockchain-crypto/crypto_api/vdf/internal/types"
	"blockchain-crypto/vdf/wesolowski_go"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Vdf_wgoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVdf_wgoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Vdf_wgoLogic {
	return &Vdf_wgoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Vdf_wgoLogic) Vdf_wgo(req *types.WgoReq) (resp *types.WgoRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		// input
		input_string := req.Input
		input, err := base64.StdEncoding.DecodeString(input_string)
		if err != nil {
			fmt.Println("err = ", err)
		}
		// 创建新的 [32]byte 数组，并从切片复制数据
		var array_input [32]byte
		copy(array_input[:], input)

		vdf := wesolowski_go.New(req.Difficulty, array_input)
		outputChannel := vdf.GetOutputChannel()
		vdf.Execute()
		output := <-outputChannel

		// json
		output_string := base64.StdEncoding.EncodeToString(output[:])

		return &types.WgoRes{
			Output: output_string,
		}, nil
	case "2":
		// input
		input_string := req.Input
		input, err := base64.StdEncoding.DecodeString(input_string)
		if err != nil {
			fmt.Println("err = ", err)
		}
		// 创建新的 [32]byte 数组，并从切片复制数据
		var array_input [32]byte
		copy(array_input[:], input)

		vdf := wesolowski_go.New(req.Difficulty, array_input)

		// output
		output_string := req.InputProof
		output, err := base64.StdEncoding.DecodeString(output_string)
		if err != nil {
			fmt.Println("err = ", err)
		}
		// 创建新的 [516]byte 数组，并从切片复制数据
		var array_output [516]byte
		copy(array_output[:], output)

		res := vdf.Verify(array_output)

		return &types.WgoRes{
			Res: res,
		}, nil
	default:
		fmt.Println("wrong num")
	}
	return
}
