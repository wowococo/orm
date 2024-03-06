package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type OutRangeCond struct {
	mCfg   *CondCfg
	mValue []any

	// mValueOpt value_opt.ValueOptCfg
}

func NewOutRangeCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("out_range condition does not support value from type(%s)", cfg.ValueFrom)
	}

	val, ok := cfg.ValueOptCfg.Value.([]any)
	if !ok {
		return nil, fmt.Errorf("out_range condition right value should be an array of length 2")
	}

	if len(val) != 2 {
		return nil, fmt.Errorf("out_range condition right value should be an array of length 2")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &OutRangeCond{
		mCfg:   cfg,
		mValue: val,
		// mValueOpt: vOpt,
	}, nil
}

func (cond *OutRangeCond) Convert(ctx context.Context) (string, error) {
	lt := cond.mValue[0]
	gte := cond.mValue[1]

	ltStr, ok := lt.(string)
	if ok {
		lt = fmt.Sprintf(`"%s"`, ltStr)
	}
	gteStr, ok := gte.(string)
	if ok {
		gte = fmt.Sprintf(`"%s"`, gteStr)
	}
	// out_range  (-inf, value[0]) || [value[1], +inf)
	dslStr := fmt.Sprintf(`
					{
						"bool": {
							"should": [
								{
									"range": {
										"%s": {
											"lt": %v
										}
									}
								},
								{
									"range": {
										"%s": {
											"gte":  %v
										}
									}
								}
							]
						}
					}`, cond.mCfg.Name, lt, cond.mCfg.Name, gte)

	return dslStr, nil
}
