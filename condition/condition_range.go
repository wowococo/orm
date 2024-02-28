package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type RangeCond struct {
	mCfg   *CondCfg
	mValue []any
	// mValueOpt value_opt.ValueOptCfg
}

func NewRangeCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("range condition does not support value from type(%s)", cfg.ValueFrom)
	}

	val, ok := cfg.ValueOptCfg.Value.([]any)
	if !ok {
		return nil, fmt.Errorf("range condition right value should be an array of length 2")
	}

	if len(val) != 2 {
		return nil, fmt.Errorf("range condition right value should be an array of length 2")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &RangeCond{
		mCfg:   cfg,
		mValue: val,
		// mValueOpt: vOpt,
	}, nil
}

func (cond *RangeCond) Convert(ctx context.Context) (string, error) {
	gte := cond.mValue[0]
	lte := cond.mValue[1]

	gteStr, ok := gte.(string)
	if ok {
		gte = fmt.Sprintf(`"%s"`, gteStr)
	}
	lteStr, ok := lte.(string)
	if ok {
		lte = fmt.Sprintf(`"%s"`, lteStr)
	}
	// range 左闭右开区间
	dslStr := fmt.Sprintf(`
				{
					"range": {
						"%s": {
							"gte": %v,
							"lt": %v
						}
					}
				}`, cond.mCfg.Name, gte, lte)

	return dslStr, nil
}
