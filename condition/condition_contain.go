package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type ContainCond struct {
	mCfg         *CondCfg
	IsSliceValue bool
	mValue       any
	mSliceValue  []any
}

func NewContainCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("contain condition does not support value from type(%s)", cfg.ValueFrom)
	}

	containCond := &ContainCond{
		mCfg: cfg,
	}

	if common.IsSlice(cfg.Value) {
		if len(cfg.Value.([]any)) == 0 {
			return nil, fmt.Errorf("contain condition right value is an empty array")
		}

		containCond.IsSliceValue = true
		containCond.mSliceValue = cfg.Value.([]any)

	} else {
		containCond.IsSliceValue = false
		containCond.mValue = cfg.Value
	}

	return containCond, nil

}

func (cond *ContainCond) Convert(ctx context.Context) (string, error) {
	
}
