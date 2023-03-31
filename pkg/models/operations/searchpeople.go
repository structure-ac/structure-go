// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type SearchPeopleApplicationJSON struct {
	// Filter for searching
	Filter *string `json:"filter,omitempty"`
	// Number of results per page (0-100)
	Limit *string `json:"limit,omitempty"`
	// The offset number to start at
	Page *string `json:"page,omitempty"`
	// Query for searching
	Query *string `json:"query,omitempty"`
}

type SearchPeopleResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
