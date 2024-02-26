package util

import (
	"context"
	"errors"
	"fmt"
	"strings"
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

func ConvertToDSL(ctx context.Context, cond CondCfg, fieldsMap map[string]string) (string, error) {
	filterStr := ""

	filterStri := ""

	// 从视图的字段信息中获取字段的类型，如果是 text 就给字段名加 .keyword；
	// 如果是脱敏字段，text 类型的加上 _desensitize.keyword, 其余类型的字段加上 _desensitize
	filterField := cond.Name
	desensitizeFiled := filterField + DESENSITIZE_FIELD_SUFFIX

	fieldType, ok1 := fieldsMap[filterField]
	_, ok2 := fieldsMap[desensitizeFiled]
	if ok1 && ok2 {
		// 脱敏字段
		filterField = desensitizeFiled
	}
	if fieldType == TEXT_TYPE {
		filterField = wrapKeyWordFieldName(filterField)
	}

	switch cond.Operation {
	case OperationIn:
		value := cond.Value.([]interface{})
		for i := 0; i < len(value); i++ {
			v := value[i]
			_, ok := value[i].(string)
			if ok {
				v = fmt.Sprintf(`"%s"`, value[i])
			}
			filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
						{
							"term": {
								"%s": {
									"value": %v
								}
							}
						}`, filterField, v))
			if i != len(value)-1 {
				filterStri = fmt.Sprintf("%s,", filterStri)
			}
		}
		filterStri = fmt.Sprintf(`
					{
						"bool": {
							"should": [%s
							]
						}
					}, `, filterStri)

	case OperationEq:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"term": {
							"%s": {
								"value": %v
							}
						}
					},`, filterField, v))
	case OperationNotEq:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}

		filterStri = fmt.Sprintf(`
					{
						"bool": {
							"must_not": [
								{
									"term": {
										"%s": {
											"value": %v
										}
									}
								}
							]
						}
					}, `, filterField, v)

	case OperationRange:
		value := cond.Value.([]interface{})
		if len(value) != 2 {
			return "", errors.New("When filter's operation is range, the value should be an array with length is 2. ")
		}
		gte := value[0]
		lte := value[1]

		gteStr, ok := gte.(string)
		if ok {
			gte = fmt.Sprintf(`"%s"`, gteStr)
		}
		lteStr, ok := lte.(string)
		if ok {
			lte = fmt.Sprintf(`"%s"`, lteStr)
		}
		// range 左闭右开区间
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"range": {
							"%s": {
								"gte": %v,
								"lt": %v
							}
						}
					},`, filterField, gte, lte))

	case OperationOutRange:
		value := cond.Value.([]interface{})
		if len(value) != 2 {
			return "", errors.New("When filter's operation is out_range, the value should be an array with length is 2. ")
		}

		lt := value[0]
		gte := value[1]

		ltStr, ok := lt.(string)
		if ok {
			lt = fmt.Sprintf(`"%s"`, ltStr)
		}
		gteStr, ok := gte.(string)
		if ok {
			gte = fmt.Sprintf(`"%s"`, gteStr)
		}
		// out_range  (-inf, value[0]] || [value[1], +inf)
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"bool": {
							"should": [
								{
									"range": {
										"%s": {
											"lt": %v
										}
									}
								},
								{
									"range": {
										"%s": {
											"gte":  %v
										}
									}
								}
							]
						}
					}, `, filterField, lt, filterField, gte))

	case OperationLike:
		v := cond.Value
		v = fmt.Sprintf(`".*%v.*"`, v)
		// vStr, ok := filter.Value.(string)
		// if ok {
		// 	v = fmt.Sprintf(`".*%s.*"`, vStr)
		// }
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"regexp": {
							"%s": %v
						}
					},`, filterField, v))

	case OperationNotLike:
		v := cond.Value
		v = fmt.Sprintf(`".*%v.*"`, v)
		// vStr, ok := filter.Value.(string)
		// if ok {
		// 	v = fmt.Sprintf(`".*%s.*"`, vStr)
		// }

		filterStri = fmt.Sprintf(`
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
					}, `, filterField, v)

	case OperationGt:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"range": {
							"%s": {
								"gt": %v
							}
						}
					},`, filterField, v))

	case OperationGte:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"range": {
							"%s": {
								"gte": %v
							}
						}
					},`, filterField, v))

	case OperationLt:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"range": {
							"%s": {
								"lt": %v
							}
						}
					},`, filterField, v))

	case OperationLte:
		v := cond.Value
		vStr, ok := cond.Value.(string)
		if ok {
			v = fmt.Sprintf(`"%s"`, vStr)
		}
		filterStri = fmt.Sprintf("%s%s", filterStri, fmt.Sprintf(`
					{
						"range": {
							"%s": {
								"lte": %v
							}
						}
					},`, filterField, v))

	default:
		return "", errors.New("Unsupport filter operation type")
	}

	filterStr = fmt.Sprintf("%s%s", filterStr, filterStri)

	return filterStr, nil

}

// 将过滤条件拼接到 dsl 请求的 query 部分
func AppendCondition(ctx context.Context, cond *CondCfg, fieldsMap map[string]string) (dsl string, err error) {
	if cond == nil {
		return "", nil
	}
	switch cond.Operation {
	case OperationAnd:
		cond, err = newAndCond(ctx, cond, fieldsMap)
	case OperationOr:
		cond, err = newOrCond(ctx, cond, fieldsMap)
	default:
		cond, err = NewCondWithOpr(ctx, cond, fieldsMap)
	}
	if err != nil {
		return "", err
	}

}

func newAndCond(ctx context.Context, cond *CondCfg, fieldsMap map[string]string) {
	subConds := []
}

// 转换成 keyword
func wrapKeyWordFieldName(fields ...string) string {
	for _, field := range fields {
		if field == "" {
			return ""
		}
	}

	return strings.Join(fields, ".") + "." + KEYWORD_SUFFIX
}
