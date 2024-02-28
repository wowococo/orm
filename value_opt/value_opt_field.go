package value_opt

import (
	"context"
	"fmt"
)

type Field struct {
	// mValueField *common.ViewField
	mValueField string
}

func NewField(ctx context.Context, cfg *ValueOptCfg, fieldsMap map[string]string) (ValueOpt, error) {
	if cfg == nil {
		return nil, nil
	}

	fieldName := cfg.Value.(string)
	field, ok := fieldsMap[fieldName]
	if !ok {
		return nil, fmt.Errorf("failed to find input field")
	}

	vOpt := &Field{
		mValueField: field,
	}

	return vOpt, nil
}

func (vOpt *Field) GetValueType() string {
	// return vOpt.mValueField.Type
	return ""
}

func (vOpt *Field) GetData() ([]any, error) {
	return nil, nil
}

func (vOpt *Field) GetSingleData() (any, error) {
	return nil, nil
}
