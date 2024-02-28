package condition

import (
	"context"
	"fmt"

	"orm/common"
	"orm/value_opt"
)

type NotLikeCond struct {
	mCfg      *CondCfg
	mValue    string
	// mValueOpt value_opt.ValueOptCfg
}

func NewNotLikeCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("not_like condition does not support value from type(%s)", cfg.ValueFrom)
	}

	// if !DataType_IsString(cfg.NameField.Type) {
	// 	return nil, fmt.Errorf("not_like condition left field is not a string field: %s:%s", cfg.NameField.Name, cfg.NameField.Type)
	// }

	val, ok := cfg.ValueOptCfg.Value.(string)
	if !ok {
		return nil, fmt.Errorf("not_like condition right value is not a string value: %v", cfg.Value)
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	return &NotLikeCond{
		mCfg:   cfg,
		mValue: val,
	}, nil
}

func (cond *NotLikeCond) Convert(ctx context.Context) (string, error) {
	v := cond.mCfg.Value
	v = fmt.Sprintf(`".*%v.*"`, v)

	dslStr := fmt.Sprintf(`
					{
						"bool": {
							"must_not": [
								{
									"regexp": {
										"%s": %v
									}
								}
							]
						}
					}`, cond.mCfg.Name, v)

	return dslStr, nil
}
