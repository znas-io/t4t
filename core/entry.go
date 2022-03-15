package core

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
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
	if path == "" {
		return nil, ErrPathCannotBeEmpty
	}

	var err error

	if tag, err = ValidateTag(tag); err != nil {
		return nil, err
	}

	var abs string

	if abs, err = filepath.Abs(path); err != nil {
		return nil, err
	}

	path = abs

	entry := &Entry{
		tag:  tag,
		path: path,
	}

	if err = entry.validatePathAndGenerateType(); err != nil {
		return nil, err
	}

	entry.generateID()

	return entry, nil
}

func (e *Entry) GetID() string {
	return e.id
}

func (e *Entry) GetPath() string {
	return e.path
}

func (e *Entry) GetTagPartition() string {
	return GetTagPartition(e.tag)
}

func (e *Entry) String() string {
	return fmt.Sprintf("%v %v\r\n", e.tag, e.path)
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
