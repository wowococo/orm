package condition

import (
	"context"
	"fmt"

	"orm/common"
)

type AndCond struct {
	mCfg      *CondCfg
	mSubConds []Condition
}

func newAndCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	subConds := []Condition{}

	if len(cfg.SubConds) == 0 {
		return nil, fmt.Errorf("")
	}

	if len(cfg.SubConds) > MaxSubCondition {
		return nil, fmt.Errorf("")
	}

	for _, subCond := range cfg.SubConds {
		cond, err := NewCondition(ctx, subCond, fieldsMap)
		if err != nil {
			return nil, err
		}

		subConds = append(subConds, cond)
	}

	return &AndCond{
		mCfg:      cfg,
		mSubConds: subConds,
	}, nil

}

func (cond *AndCond) Convert(ctx context.Context) (string, error) {
	res := `
	{
		"bool": {
			"filter": [
				%s
			]
		}
	}
	`

	dslStr := ""
	for i, subCond := range cond.mSubConds {
		dsl, err := subCond.Convert(ctx)
		if err != nil {
			return "", err
		}

		if i != len(cond.mSubConds)-1 {
			dsl += ","
		}

		dslStr += dsl

	}

	res = fmt.Sprintf(res, dslStr)
	return res, nil

}
