package condition

import (
	"context"
	"fmt"
	"strings"

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
					"filter": [
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
							"term": {
								"%s": {
									"value": %v
								}
							}
						}`, cond.mCfg.Name, val)
	}

	return dslStr, nil
}
