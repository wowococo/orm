package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type GteCond struct {
	mCfg *CondCfg
	// mValueOpt value_opt.ValueOptCfg
}

func NewGteCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("gte condition does not support value from type(%s)", cfg.ValueFrom)
	}

	if common.IsSlice(cfg.ValueOptCfg.Value) {
		return nil, fmt.Errorf("gte condition only supports single value")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &GteCond{
		mCfg: cfg,
	}, nil

}

func (cond *GteCond) Convert(ctx context.Context) (string, error) {
	v := cond.mCfg.Value
	vStr, ok := v.(string)
	if ok {
		v = fmt.Sprintf(`"%s"`, vStr)
	}

	dslStr := fmt.Sprintf(`
				{
					"range": {
						"%s": {
							"gte": %v
						}
					}
				}`, cond.mCfg.Name, v)

	return dslStr, nil

}
