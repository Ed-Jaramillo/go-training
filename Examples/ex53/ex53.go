package main

import (
	"errors"
)

func validateStatus(status string) error {
	err1 := errors.New("status cannot be empty")
	err2 := errors.New("status exceerds 140 characters")
	if status == "" {
		return err1
	}
	
	if len(status) > 140 {
		return err2
	}
	return nil
}