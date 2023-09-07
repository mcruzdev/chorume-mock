package generator

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/wiremock/go-wiremock"
	"strconv"
)

func GenerateStubRules(openapi3 *openapi3.T) []*wiremock.StubRule {

	var stubRules []*wiremock.StubRule
	for path, pathItem := range openapi3.Paths {
		if pathItem.Get != nil {
			stubRule := wiremock.NewStubRule("GET", wiremock.URLMatching(path))

			for statusCode := range pathItem.Get.Responses {
				statusCodeInt, err := strconv.Atoi(statusCode)
				if err != nil {
					panic(any(err))
				}

				if Is2xx(statusCodeInt) {
					stubRule.WillReturnResponse(wiremock.NewResponse().WithStatus(int64(statusCodeInt)))
				}
			}

			stubRules = append(stubRules, stubRule)
		}
	}
	return stubRules
}

func Is2xx(status int) bool {
	return status >= 200 && status < 300
}
