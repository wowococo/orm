package condition

import (
	"context"
	"fmt"
)

type EqCond struct {
	mCfg      *CondCfg
	mValueOpt *ValueOptCfg
}

func NewEqCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]string) (Condition, error) {
	if _, ok := fieldsMap[cfg.Name]; !ok {
		return nil, fmt.Errorf("")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &EqCond{
		mCfg: cfg,
	}, nil

}

func (cond *EqCond) Convert(ctx context.Context) (string, error) {
	
}
