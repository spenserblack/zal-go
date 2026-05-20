package cmd

import "errors"

// errInvalidMinMax is raised when the maximum is less than the minimum.
var errInvalidMinMax = errors.New("max cannot be less than min")

// assertMinMax returns an error if the maximum is less than the minimum.
func assertMinMax(min, max int) error {
	if max < min {
		return errInvalidMinMax
	}
	return nil
}
