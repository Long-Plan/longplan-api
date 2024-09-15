package requestor

import (
	"fmt"
	"strings"
)

func BuildQueryParams(params map[string]interface{}) string {
	var queries []string
	for key, val := range params {
		query := fmt.Sprintf("%v=%v", key, val)
		queries = append(queries, query)
	}

	if len(queries) > 0 {
		return "?" + strings.Join(queries, "&")
	}
	return ""
}
