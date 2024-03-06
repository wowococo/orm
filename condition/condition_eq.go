package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type EqCond struct {
	mCfg *CondCfg
	// mValueOpt ValueOptCfg
}

func NewEqCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("eq condition does not support value from type(%s)", cfg.ValueFrom)
	}

	if common.IsSlice(cfg.ValueOptCfg.Value) {
		return nil, fmt.Errorf("eq condition only supports single value")
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, &cfg.ValueOptCfg, cfg.NameField.Type, fieldsMap)
	// if err != nil {
	// 	return nil, err
	// }

	return &EqCond{
		mCfg: cfg,
		// mValueOpt: vOpt,
	}, nil

}

func (cond *EqCond) Convert(ctx context.Context) (string, error) {
	v := cond.mCfg.Value
	vStr, ok := v.(string)
	if ok {
		v = fmt.Sprintf(`"%s"`, vStr)
	}

	dslStr := fmt.Sprintf(`
					{
						"term": {
							"%s": {
								"value": %v
							}
						}
					},`, cond.mCfg.Name, v)

	return dslStr, nil
}
