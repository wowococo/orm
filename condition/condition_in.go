package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type InCond struct {
	mCfg   *CondCfg
	mValue []any
	// mValueOpt value_opt.ValueOptCfg
}

func NewInCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("in condition does not support value from type(%s)", cfg.ValueFrom)
	}

	if !common.IsSlice(cfg.ValueOptCfg.Value) {
		return nil, fmt.Errorf("in condition right value should be an array")
	}

	if !common.IsSameType(cfg.ValueOptCfg.Value.([]any)) {
		return nil, fmt.Errorf("in condition right value should be an array composed of elements of same type")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &InCond{
		mCfg:   cfg,
		mValue: cfg.ValueOptCfg.Value.([]any),
	}, nil
}

func (cond *InCond) Convert(ctx context.Context) (string, error) {
	value := cond.mValue

	dslStr := fmt.Sprintf(`
					{
						"terms": {
							"%s": %v
						}
					}`, cond.mCfg.Name, value)

	return dslStr, nil

}

// func (cond *InCond) Convert(ctx context.Context) (string, error) {
// 	value := cond.mValue

// 	var dslStr string
// 	for i := 0; i < len(value); i++ {
// 		v := value[i]
// 		_, ok := value[i].(string)
// 		if ok {
// 			v = fmt.Sprintf(`"%s"`, value[i])
// 		}

// 		dslStr = fmt.Sprintf("%s%s", dslStr, fmt.Sprintf(`
// 						{
// 							"term": {
// 								"%s": {
// 									"value": %v
// 								}
// 							}
// 						}`, cond.mCfg.Name, v))
// 		if i != len(value)-1 {
// 			dslStr = fmt.Sprintf("%s,", dslStr)
// 		}
// 	}

// 	dslStr = fmt.Sprintf(`
// 					{
// 						"bool": {
// 							"should": [%s
// 							]
// 						}
// 					}, `, dslStr)

// 	return dslStr, nil

// }
