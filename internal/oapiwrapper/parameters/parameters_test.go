package parameters_test

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"

	"github.com/mcruzdev/speckify/internal/oapiwrapper/parameters"
)

func TestParameters_ParameterAsInt(t *testing.T) {

	p := openapi3.NewQueryParameter("page")
	p.Example = 10

	ans := parameters.AppendQueryParameter("", p)

	assert.Equal(t, "page=10", ans)

}

func TestParameters_ParameterAsString(t *testing.T) {

	p := openapi3.NewQueryParameter("id")
	p.Example = "b4e2dc44-90cb-11ee-b9d1-0242ac120002"

	ans := parameters.AppendQueryParameter("", p)

	assert.Equal(t, "id=b4e2dc44-90cb-11ee-b9d1-0242ac120002", ans)
}
