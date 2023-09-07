package oapiwrapper

import "github.com/getkin/kin-openapi/openapi3"

func init() {

}

func Get(location string) *openapi3.T {
	loader := openapi3.NewLoader()
	file, err := loader.LoadFromFile(location)
	if err != nil {
		panic(any(err))
	}
	return file
}
