// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CompanyEmployeesPathParams struct {
	// ID of the company
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type CompanyEmployeesQueryParams struct {
	// The offset number to start at
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// Number of results per page (0-100)
	PerPage *string `queryParam:"style=form,explode=true,name=per_page"`
}

type CompanyEmployeesRequest struct {
	PathParams  CompanyEmployeesPathParams
	QueryParams CompanyEmployeesQueryParams
}

type CompanyEmployeesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
