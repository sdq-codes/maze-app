package errors

import (
	"fmt"
	"strings"
)

var (
	listOfValidationErrors []string
)

func FormatValidationMessage(fields map[string]string, errorText string, messageBody string) []string {
	if messageBody == "" {
		messageBody = " already exists"
	}
	for fieldRep, fieldName := range fields {
		if  strings.Contains(errorText, fieldRep){
			tmpMssg := fmt.Sprintf("%s%s", fieldName, messageBody)
			listOfValidationErrors = append(listOfValidationErrors, tmpMssg)
		}
	}
	return listOfValidationErrors
}
