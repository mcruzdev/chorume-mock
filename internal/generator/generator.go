package generator

import (
	"fmt"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mcruzdev/speckify/internal/model"
	"github.com/mcruzdev/speckify/internal/oapiwrapper/content"
	"github.com/mcruzdev/speckify/internal/oapiwrapper/operation"
	"github.com/mcruzdev/speckify/internal/oapiwrapper/parameters"
	"github.com/wiremock/go-wiremock"
)

type httpMethodHandler func(method string) func(*openapi3.Operation)

func operationHandler(path string, pathItem *openapi3.PathItem) *wiremock.StubRule {
	for key, op := range pathItem.Operations() {
		switch key {
		case "GET":
			return generateStubGet(path, op)
		default:
			panic(any(model.ErrMethodNotFound))
		}
	}
	return nil
}

func generateStubGet(path string, op *openapi3.Operation) *wiremock.StubRule {

	if !operation.Validate(op) {
		return nil
	}

	sr := wiremock.NewStubRule("GET", wiremock.URLEqualTo(path))
	r := wiremock.NewResponse()

	params := op.Parameters

	var query string = ""

	for _, p := range params {
		switch p.Value.In {
		case openapi3.ParameterInQuery:
			query += parameters.AppendQueryParameter(query, p.Value)
		}
	}

	for status, res := range op.Responses {

		statusCode, err := strconv.Atoi(status)
		if err != nil {
			return nil
		}

		if Is2xx(statusCode) {
			r = r.WithStatus(int64(statusCode))
		}

		mt, err := content.GetMediaType(res.Value.Content)
		if err != nil {
			return nil
		}

		i := mt.Example

		switch t := i.(type) {
		case string:
			r = r.WithBody(fmt.Sprintf("%v", t))
		}
	}

	return sr.WillReturnResponse(r)
}

func GenerateStubRules(openapi3 *openapi3.T) []*wiremock.StubRule {

	var stubRules []*wiremock.StubRule

	for path, pathItem := range openapi3.Paths {
		sr := operationHandler(path, pathItem)
		stubRules = append(stubRules, sr)
	}
	return stubRules
}

func Is2xx(status int) bool {
	return status >= 200 && status < 300
}
