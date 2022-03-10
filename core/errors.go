package core

import (
	"errors"
	"fmt"
)

var (
	ErrTagCannotBeEmpty  = errors.New("tag cannot be empty")
	ErrPathCannotBeEmpty = errors.New("path cannot be empty")
	ErrInvalidTag        = func(tag, regex string) error {
		return errors.New(fmt.Sprintf("invalid tag %v does not match %v", tag, regex))
	}
	ErrUnthinkable = func(tag0, path0, tag1, path1 string) error {
		return errors.New(fmt.Sprintf(
			"the unthinkable has happened, the same ID was generated for combination %v %v and %v %v, please report this on github", tag0, path0, tag1, path1))
	}
	ErrKeyCannotBeEmpty = errors.New("key cannot be empty")
	ErrEntryCannotBeNil = errors.New("entry cannot be nil")
)
