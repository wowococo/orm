package condition

import (
	"context"
	"fmt"
	"strings"

	"orm/common"
	"orm/value_opt"
)

type NotContainCond struct {
	mCfg         *CondCfg
	IsSliceValue bool
	mValue       any
	mSliceValue  []any
}

func NewNotContainCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("not_contain condition does not support value from type(%s)", cfg.ValueFrom)
	}

	notContainCond := &NotContainCond{
		mCfg: cfg,
	}

	if common.IsSlice(cfg.Value) {
		if len(cfg.Value.([]any)) == 0 {
			return nil, fmt.Errorf("not_contain condition right value is an empty array")
		}

		notContainCond.IsSliceValue = true
		notContainCond.mSliceValue = cfg.Value.([]any)

	} else {
		notContainCond.IsSliceValue = false
		notContainCond.mValue = cfg.Value
	}

	return notContainCond, nil

}

func (cond *NotContainCond) Convert(ctx context.Context) (string, error) {
	var dslStr string
	if cond.IsSliceValue {
		subStrs := []string{}
		for _, val := range cond.mSliceValue {
			vStr, ok := val.(string)
			if ok {
				val = fmt.Sprintf(`"%s"`, vStr)
			}

			subStr := fmt.Sprintf(`
						{
							"term": {
								"%s": {
									"value": %v
								}
							}
						}`, cond.mCfg.Name, val)

			subStrs = append(subStrs, subStr)

		}

		dslStr = fmt.Sprintf(`
			{
				"bool": {
					"must_not": [
						%s
					]
				}
			}
		`, strings.Join(subStrs, ","))

	} else {
		val := cond.mValue
		vStr, ok := val.(string)
		if ok {
			val = fmt.Sprintf(`"%s"`, vStr)
		}

		dslStr = fmt.Sprintf(`
						{
							"bool": {
								"must_not": {
									"term": {
										"%s": {
											"value": %v
										}
									}
								}
							}
						}`, cond.mCfg.Name, val)
	}

	return dslStr, nil
}
