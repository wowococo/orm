package main

import (
	"context"
	"encoding/json"
	"fmt"
	"orm/common"
	"orm/condition"
)

var condStr = `
{
	  "operation": "and",
	  "sub_conditions": [
		{
		  "operation": "or",
		  "sub_conditions": [
			{
			  "operation": "==",
			  "field": "f1",
			  "value_from": "const",
			  "value": "123"
			},
			{
			  "operation": "!=",
			  "field": "f2",
			  "value_from": "const",
			  "value": "hhhh"
			}
		  ]
		},
		{
		  "operation": "==",
		  "field": "f4",
		  "value_from": "const",
		  "value": "group"
		}
	  ]
}
`

func initFields() map[string]*common.ViewField {
	fieldsMap := make(map[string]*common.ViewField)

	fieldsMap["f1"] = &common.ViewField{
		Name: "f1",
		Type: "keyword",
	}
	fieldsMap["f2"] = &common.ViewField{
		Name: "f2",
		Type: "double",
	}
	fieldsMap["f4"] = &common.ViewField{
		Name: "f4",
		Type: "long",
	}

	return fieldsMap
}

func main() {
	var cfg condition.CondCfg
	err := json.Unmarshal([]byte(condStr), &cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	cond, err := condition.NewCondition(ctx, &cfg, initFields())
	if err != nil {
		panic(err)
	}

	dsl, err := cond.Convert(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(dsl)
}
