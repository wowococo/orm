package value_opt

import (
	"context"
	"fmt"
	
	"orm/common"
)

type Const struct {
	mValueType  string
	mConstValue []any
}

func NewConst(ctx context.Context, cfg *ValueOptCfg, keyType string) (ValueOpt, error) {
	vOpt := &Const{}
	if common.IsSlice(cfg.Value) {
		if len(cfg.Value.([]any)) == 0 {
			return nil, fmt.Errorf("const value is empty")
		}

		arrV := cfg.Value.([]any)

		vOpt.mValueType = keyType
		vOpt.mConstValue = arrV
	} else {
		vOpt.mValueType = keyType
		vOpt.mConstValue = []any{cfg.Value}
	}

	return vOpt, nil

}

func (vOpt *Const) GetValueType() string {
	return vOpt.mValueType
}

func (vOpt *Const) GetData() ([]any, error) {
	return vOpt.mConstValue, nil
}

func (vOpt *Const) GetSingleData() (any, error) {
	if len(vOpt.mConstValue) == 0 {
		return nil, nil
	} else if len(vOpt.mConstValue) > 1 {
		return nil, fmt.Errorf("only support single data: %v", vOpt.mConstValue...)
	}

	return vOpt.mConstValue[0], nil
}
