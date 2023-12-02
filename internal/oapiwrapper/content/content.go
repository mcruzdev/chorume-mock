package content

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mcruzdev/speckify/internal/model"
)

func GetMediaType(c openapi3.Content) (*openapi3.MediaType, error) {
	mimeTypes := []string{
		"application/json",
		"text/plain",
	}

	for _, mime := range mimeTypes {

		mt := c.Get(mime)

		if mt != nil {
			return mt, nil
		}
	}

	return nil, model.ErrMimeTypeNotAvailable
}
