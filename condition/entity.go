package condition

import (
	"orm/common"
	"orm/value_opt"
)

const (
	KEYWORD_SUFFIX           = "keyword"
	DESENSITIZE_FIELD_SUFFIX = "_desensitize"
	TEXT_TYPE                = "text"

	DataType_Byte      = "byte"
	DataType_Long      = "long"
	DataType_Integer   = "integer"
	DataType_Float     = "float"
	DataType_Double    = "double"
	DataType_Short     = "short"
	DataType_HalfFloat = "half_float"

	DataType_Text    = "text"
	DataType_Keyword = "keyword"
	DataType_Binary  = "binary"

	DataType_Boolean = "boolean"

	DataType_Date = "date"

	DataType_Ip = "ip"

	DataType_GeoPoint = "geo_point"

	DataType_GeoShape = "geo_shape"
)

func DataType_IsString(t string) bool {
	return (t == DataType_Text || t == DataType_Keyword)
}

func DataType_IsNumber(t string) bool {
	return (t == DataType_Short || t == DataType_Integer || t == DataType_Long ||
		t == DataType_Float || t == DataType_Double)
}

const (
	OperationAnd = "and"
	OperationOr  = "or"

	OperationEq         = "=="
	OperationNotEq      = "!="
	OperationGt         = ">"
	OperationGte        = ">="
	OperationLt         = "<"
	OperationLte        = "<="
	OperationIn         = "in"
	OperarionNotIn      = "not_in"
	OperationLike       = "like"
	OperationNotLike    = "not_like"
	OperationContain    = "contain"
	OperationNotContain = "not_contain"
	OperationRange      = "range"
	OperationOutRange   = "out_range"
	OperationExist      = "exist"
	operationNotExist   = "not_exist"
	OperationRegex      = "regex"
)

type Filter struct {
	Name      string `json:"name"`
	Operation string `json:"operation"`
	Value     any    `json:"value"`
}

type CondCfg struct {
	Name                  string     `json:"field" mapstructure:"field"`
	Operation             string     `json:"operation" mapstructure:"operation"`
	SubConds              []*CondCfg `json:"sub_conditions" mapstructure:"sub_conditions"`
	value_opt.ValueOptCfg `mapstructure:",squash"`

	NameField *common.ViewField `json:"-" mapstructure:"-"`
}
