package validation

import (
	"auth/pkg/enum"
	"fmt"
	"strings"
)

type RequiredRule struct {
	FieldName string
	Value     interface{}
}

//## Validate Rule = 1
func Required(fields *[]RequiredRule, errs *[]string) int {
	for _, f := range *fields {
		if strings.TrimSpace(fmt.Sprintf("%v", f.Value)) == "" || f.Value == nil {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 1
	} else {
		return enum.Ok
	}
}
