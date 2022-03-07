package core

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func ValidateTags(regex string, tags ...string) error {
	if len(tags) == 0 {
		return errors.New("requires at least one tag")
	}

	for _, tag := range tags {
		if tag == "" {
			return errors.New("tags cannot be empty")
		}

		if match, err := regexp.MatchString(regex, tag); err != nil {
			return err
		} else if !match {
			return errors.New(fmt.Sprintf("tag %v does not match %v", tag, regex))
		}
	}

	return nil
}

func ValidatePaths(paths ...string) error {
	if len(paths) == 0 {
		return errors.New("requires at least one file or directory to tag")
	}

	for _, path := range paths {
		if path == "" {
			return errors.New("path cannot be empty")
		}

		var file *os.File
		var err error

		if file, err = os.Open(path); err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}
	}

	return nil
}
