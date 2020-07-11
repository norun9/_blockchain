package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int //X座標
	S *big.Int //transactionのハッシュなどを元に導き出したもの
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
