package core

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	validTagRegex = "^[a-z0-9][a-z0-9-]+[a-z0-9]$"
)

type (
	Tag struct {
		id   string
		tag  string
		path string
		t    string
	}
)

func NewTag(tag, path string) (*Tag, error) {
	if tag == "" {
		return nil, ErrTagCannotBeEmpty
	}

	if path == "" {
		return nil, ErrPathCannotBeEmpty
	}

	tag = strings.ToLower(tag)

	var matched bool
	var err error

	if matched, err = regexp.MatchString(validTagRegex, tag); err != nil {
		return nil, err
	} else if !matched {
		return nil, ErrInvalidTag(tag, validTagRegex)
	}

	var abs string

	if abs, err = filepath.Abs(path); err != nil {
		return nil, err
	}

	path = abs

	i := &Tag{
		tag:  tag,
		path: path,
	}

	if err = i.validatePathAndGenerateType(); err != nil {
		return nil, err
	}

	i.generateID()

	return i, nil
}

func (t *Tag) GetID() string {
	return t.id
}

func (t *Tag) GetTag() string {
	return t.tag
}

func (t *Tag) GetPath() string {
	return t.path
}

func (t *Tag) GetType() string {
	return t.t
}

func (t *Tag) String() string {
	return fmt.Sprintf("%v %v", t.tag, t.path)
}

func (t *Tag) FileString() string {
	return fmt.Sprintf("%v %v %v %v", t.id, t.t, t.tag, t.path)
}

func (t *Tag) generateID() {
	data := md5.Sum([]byte(fmt.Sprintf("%v:%v", t.tag, t.path)))

	t.id = fmt.Sprintf("%x", data)
}

func (t *Tag) validatePathAndGenerateType() error {
	var file *os.File
	var err error

	if file, err = os.Open(t.path); err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	var info os.FileInfo

	if info, err = file.Stat(); err != nil {
		return err
	}

	t.t = "file"

	if info.IsDir() {
		t.t = "dir"
	}

	return nil
}
