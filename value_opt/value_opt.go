package value_opt

import (
	"context"
	"fmt"
)

const (
	ValueFrom_Const = "const"
	ValueFrom_Field = "field"
	ValueFrom_User = "user" 
)

type ValueOpt interface {
	GetValueType() string
}

type ValueOptCfg struct {
	ValueFrom string `json:"value_from" mapstructure:"value_from"`
	Value any `json:"value" mapstructure:"value"`
}

func NewValueOpt(ctx context.Context, cfg *ValueOptCfg, keyType string, fieldsMap map[string]string) (vOpt ValueOpt, err error) {
	if cfg == nil {
		return nil, nil
	}

	switch cfg.ValueFrom{
	case ValueFrom_Const:
		vOpt, err = NewConst(ctx, cfg, keyType)
	case ValueFrom_Field:
		vOpt, err = NewField(ctx, cfg, fieldsMap)
	case ValueFrom_User:
	default:
		return nil , fmt.Errorf("invalid value from type: %s", cfg.ValueFrom)
	}
	if err != nil {
		return nil, err
	}

	return vOpt, nil
}