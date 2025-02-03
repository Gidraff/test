// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 */

package openapi




// ApiResponse - Describes the result of uploading an image resource
type ApiResponse struct {

	Code int32 `json:"code,omitempty"`

	Type string `json:"type,omitempty"`

	Message string `json:"message,omitempty"`
}

// AssertApiResponseRequired checks if the required fields are not zero-ed
func AssertApiResponseRequired(obj ApiResponse) error {
	return nil
}

// AssertApiResponseConstraints checks if the values respects the defined constraints
func AssertApiResponseConstraints(obj ApiResponse) error {
	return nil
}
