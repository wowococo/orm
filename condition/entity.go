package condition

const (
	KEYWORD_SUFFIX           = "keyword"
	DESENSITIZE_FIELD_SUFFIX = "_desensitize"
	TEXT_TYPE                = "text"

	// number
	BYTE       = "byte"
	LONG       = "long"
	INTEGER    = "integer"
	FLOAT      = "float"
	DOUBLE     = "double"
	SHORT      = "short"
	HALF_FLOAT = "half_float"

	// string
	TEXT    = "text"
	KEYWORD = "keyword"
	BINARY  = "binary"

	// bool
	BOOLEAN = "boolean"

	// date
	DATE = "date"

	// ip
	IP = "ip"

	// geo_point
	GEO_POINT = "geo_point"

	// geo_shape
	GEO_SHAPE = "geo_shape"
)

const (
	OperationAnd = "and"
	OperationOr  = "or"

	OperationEq         = "="
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

type ViewField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Hidden   bool   `json:"hidden"`
	Comment  string `json:"comment"`
	Format   string `json:"format,omitempty"`
	Analyzer string `json:"analyzer,omitempty"`

	Path []string `json:"-"`
}

type CondCfg struct {
	Name        string     `json:"field" mapstructure:"field"`
	Operation   string     `json:"operation" mapstructure:"operation"`
	SubConds    []*CondCfg `json:"sub_conditions" mapstructure:"sub_conditions"`
	ValueOptCfg `mapstructure:",squash"`

	NameField *ViewField `json:"-" mapstructure:"-"`
}

type ValueOptCfg struct {
	ValueFrom string `json:"value_from" mapstructure:"value_from"`
	Value     any    `json:"value" mapstructure:"value"`
}
