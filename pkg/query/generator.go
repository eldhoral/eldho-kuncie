package generator

import (
	"strings"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

func DynamicUpdateStatement(column []string, params data.Params) string {
	var columns []string

	for row, dataColumn := range column {
		value := params.GetString(dataColumn)
		totalRow := len(column)
		if value != "" {
			if row+1 == totalRow {
				columns = append(columns, dataColumn+" = :"+dataColumn+" ")
			} else {
				columns = append(columns, dataColumn+" = :"+dataColumn+", ")
			}
		}
	}
	queryColumn := strings.Join(columns, "")
	return queryColumn
}
