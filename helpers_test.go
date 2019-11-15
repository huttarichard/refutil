package refutil

import "errors"

type TestError struct{}

func (err *TestError) Error() string { return "test" }

func emptyError() error {
	// Create new error of error interface
	err := errors.New("retyped")
	err = (*TestError)(nil)
	return err
}
