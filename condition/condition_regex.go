package condition

import (
	"context"
	"fmt"

	"github.com/dlclark/regexp2"

	"orm/common"
	"orm/value_opt"
)

type RegexCond struct {
	mCfg    *CondCfg
	mValue  string
	mRegexp *regexp2.Regexp
}

func NewRegexCond(ctx context.Context, cfg *CondCfg, fieldsMap map[string]*common.ViewField) (Condition, error) {
	if cfg.ValueOptCfg.ValueFrom != value_opt.ValueFrom_Const {
		return nil, fmt.Errorf("regex condition does not support value from type(%s)", cfg.ValueFrom)
	}

	// if !DataType_IsString(cfg.NameField.Type) {
	// 	return nil, fmt.Errorf("like condition left field is not a string field: %s:%s", cfg.NameField.Name, cfg.NameField.Type)
	// }

	val, ok := cfg.ValueOptCfg.Value.(string)
	if !ok {
		return nil, fmt.Errorf("regex condition right value is not a string value: %v", cfg.Value)
	}

	// vOpt, err := value_opt.NewValueOpt(ctx, (*value_opt.ValueOptCfg)(&cfg.ValueOptCfg), fieldsMap)
	// if err != nil {

	// }

	regexp, err := regexp2.Compile(val, regexp2.RE2)
	if err != nil {
		return nil, fmt.Errorf("regular expression error: %s", err.Error())
	}

	return &RegexCond{
		mCfg:    cfg,
		mValue:  val,
		mRegexp: regexp,
	}, nil
}

func (cond *RegexCond) Convert(ctx context.Context) (string, error) {
	v := cond.mValue
	dslStr := fmt.Sprintf(`
					{
						"regexp": {
							"%s": %v
						}
					}`, cond.mCfg.Name, v)

	return dslStr, nil
}
