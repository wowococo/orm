package condition

import (
	"context"
	"fmt"

	"orm/common"
)

type OrCond struct {
	mCfg      *CondCfg
	mSubConds []Condition
}

func newOrCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	subConds := []Condition{}

	if len(cfg.SubConds) == 0 {
		return nil, fmt.Errorf("sub condition size is 0")
	}

	if len(cfg.SubConds) > MaxSubCondition {
		return nil, fmt.Errorf("sub condition size limit 5 but %d", len(cfg.SubConds))
	}

	for _, subCond := range cfg.SubConds {
		cond, err := NewCondition(ctx, subCond, fieldsMap)
		if err != nil {
			return nil, err
		}

		subConds = append(subConds, cond)
	}

	return &OrCond{
		mCfg:      cfg,
		mSubConds: subConds,
	}, nil

}

func (cond *OrCond) Convert(ctx context.Context) (string, error) {
	res := `
	{
		"bool": {
			"should": [
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
