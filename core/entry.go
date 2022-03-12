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
	Entry struct {
		id   string
		tag  string
		path string
		t    string
	}
)

func NewEntry(tag, path string) (*Entry, error) {
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

	i := &Entry{
		tag:  tag,
		path: path,
	}

	if err = i.validatePathAndGenerateType(); err != nil {
		return nil, err
	}

	i.generateID()

	return i, nil
}

func (e *Entry) GetID() string {
	return e.id
}

func (e *Entry) GetTagPartition() string {
	return e.tag[0:1]
}

func (e *Entry) String() string {
	return fmt.Sprintf("%v %v", e.tag, e.path)
}

func (e *Entry) FileString() string {
	return fmt.Sprintf("%v %v %v %v%v", e.id, e.t, e.tag, e.path, "\r\n")
}

func (e *Entry) generateID() {
	data := md5.Sum([]byte(fmt.Sprintf("%v:%v", e.tag, e.path)))

	e.id = fmt.Sprintf("%x", data)
}

func (e *Entry) validatePathAndGenerateType() error {
	var file *os.File
	var err error

	if file, err = os.Open(e.path); err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	var info os.FileInfo

	if info, err = file.Stat(); err != nil {
		return err
	}

	e.t = "file"

	if info.IsDir() {
		e.t = "dir"
	}

	return nil
}
