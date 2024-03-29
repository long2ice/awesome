package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/long2ice/fibers/swagger"
)

func NewSwagger() *swagger.Swagger {
	return swagger.New("Awesome", "Search for awesome github project", "0.1.0",
		swagger.Contact(&openapi3.Contact{
			Name:  "long2ice",
			URL:   "https://github.com/long2ice/awesome",
			Email: "long2ice@gmail.com",
		}),
		swagger.TermsOfService("https://github.com/long2ice"),
	)
}
