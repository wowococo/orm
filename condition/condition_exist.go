package condition

import (
	"context"
	"fmt"
)

type ExistCond struct {
	mCfg       *CondCfg
	mfieldName string
}

func NewExistCond(ctx context.Context) {

}

func (cond *ExistCond) Convert(ctx context.Context) (string, error) {
	dslStr := `
	{
		"exists": {
			"field": "%s"
		}
	}
	`

	return fmt.Sprintf(dslStr, cond.mfieldName), nil
}
