package condition

import (
	"context"
	"fmt"
)

type NotExistCond struct {
	mCfg       *CondCfg
	mfieldName string
}

func NewNotExistCond(ctx context.Context, cfg *CondCfg) (Condition, error) {
	return &NotExistCond{
		mCfg:       cfg,
		mfieldName: cfg.Name,
	}, nil
}

func (cond *NotExistCond) Convert(ctx context.Context) (string, error) {
	dslStr := `
	{
		"bool": {
			"must_not": [
				{
					"exists": {
						"field": "%s"
					}
				}
			]
		}
	}
	`

	return fmt.Sprintf(dslStr, cond.mfieldName), nil
}
