package parameters

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func AppendQueryParameter(query string, parameter *openapi3.Parameter) string {
	if len(query) == 0 {
		return fmt.Sprintf("%s=%v", parameter.Name, parameter.Example)
	}
	return fmt.Sprintf("%s&%s=%v", query, parameter.Name, &parameter.Example)
}
