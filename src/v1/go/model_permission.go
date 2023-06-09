/*
 * cloudwpss23-openapi-cyan
 *
 * OpenAPI Reference für das CloudWP der HAW Hamburg für das SommerSemster 2023
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Permission struct {

	Owner string `json:"owner,omitempty"`

	Operations []string `json:"operations,omitempty"`
}

// AssertPermissionRequired checks if the required fields are not zero-ed
func AssertPermissionRequired(obj Permission) error {
	return nil
}

// AssertRecursePermissionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Permission (e.g. [][]Permission), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePermissionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPermission, ok := obj.(Permission)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPermissionRequired(aPermission)
	})
}
