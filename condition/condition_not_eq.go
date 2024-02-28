package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type NotEqCond struct {
	mCfg *CondCfg
	// mValueOpt value_opt.ValueOptCfg
}

func NewNotEqCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("not_eq condition does not support value from type(%s)", cfg.ValueFrom)
	}

	if common.IsSlice(cfg.ValueOptCfg.Value) {
		return nil, fmt.Errorf("not_eq condition only supports single value")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &NotEqCond{
		mCfg: cfg,
	}, nil

}

func (cond *NotEqCond) Convert(ctx context.Context) (string, error) {
	v := cond.mCfg.Value
	vStr, ok := cond.mCfg.Value.(string)
	if ok {
		v = fmt.Sprintf(`"%s"`, vStr)
	}

	dslStr := fmt.Sprintf(`
					{
						"bool": {
							"must_not": [
								{
									"term": {
										"%s": %v
									}
								}
							]
						}
					}`, cond.mCfg.Name, v)

	return dslStr, nil
}
