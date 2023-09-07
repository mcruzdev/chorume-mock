package model

import "github.com/wiremock/go-wiremock"

type T struct {
	Mappings []*wiremock.StubRule `json:"mappings"`
}

func NewT(mappings []*wiremock.StubRule) T {
	return T{
		Mappings: mappings,
	}
}
