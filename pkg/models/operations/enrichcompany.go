// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EnrichCompanyQueryParams struct {
	// Country code of the company
	CountryCode *string `queryParam:"style=form,explode=true,name=country_code"`
	// The headquarters of the company
	Headquarters *string `queryParam:"style=form,explode=true,name=headquarters"`
	// ID of the company
	ID *string `queryParam:"style=form,explode=true,name=id"`
	// Game of the company
	Name *string `queryParam:"style=form,explode=true,name=name"`
}

type EnrichCompanyRequest struct {
	QueryParams EnrichCompanyQueryParams
}

type EnrichCompanyResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
