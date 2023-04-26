/*
 * cloudwpss23-openapi-cyan
 *
 * OpenAPI Reference für das CloudWP der HAW Hamburg für das SommerSemster 2023
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetFiles200Response struct {

	Bucket string `json:"bucket,omitempty"`

	Files []FileInfo `json:"files,omitempty"`
}

// AssertGetFiles200ResponseRequired checks if the required fields are not zero-ed
func AssertGetFiles200ResponseRequired(obj GetFiles200Response) error {
	for _, el := range obj.Files {
		if err := AssertFileInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGetFiles200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetFiles200Response (e.g. [][]GetFiles200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetFiles200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetFiles200Response, ok := obj.(GetFiles200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetFiles200ResponseRequired(aGetFiles200Response)
	})
}
