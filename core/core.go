package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	tagIndex      = 0
	pathIndex     = 1
	validTagRegex = "^[a-z0-9][a-z0-9-]+[a-z0-9]$"
)

func MapFileEntriesByID(f *os.File) (map[string]*Entry, error) {
	var entries Entries
	var err error

	if entries, err = getFileEntries(f); err != nil {
		return nil, err
	}

	m := make(map[string]*Entry)

	for _, entry := range entries {
		m[entry.id] = entry
	}

	return m, nil
}

func MapFileEntriesByTag(f *os.File) (map[string]Entries, error) {
	var entries Entries
	var err error

	if entries, err = getFileEntries(f); err != nil {
		return nil, err
	}

	m := make(map[string]Entries)

	for _, entry := range entries {
		if _, ok := m[entry.tag]; !ok {
			m[entry.tag] = make(Entries, 0)
		}

		m[entry.tag] = append(m[entry.tag], entry)
	}

	return m, nil
}

func GetDataFile(partition string) (*os.File, error) {
	var home string
	var err error

	if home, err = os.UserHomeDir(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%v/.t4t", home)

	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	path = fmt.Sprintf("%v/.t4t/%v", home, partition)

	var file *os.File

	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if file, err = os.Create(path); err != nil {
			return nil, err
		}
		return file, nil
	} else if err != nil {
		return nil, err
	}

	if file, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, os.ModeAppend); err != nil {
		return nil, err
	}

	return file, nil
}

func ValidateTag(tag string) (string, error) {
	if tag == "" {
		return tag, ErrTagCannotBeEmpty
	}

	tag = strings.ToLower(tag)

	var matched bool
	var err error

	if matched, err = regexp.MatchString(validTagRegex, tag); err != nil {
		return tag, err
	} else if !matched {
		return tag, ErrInvalidTag(tag, validTagRegex)
	}

	return tag, nil
}

func GetTagPartition(tag string) string {
	return tag[0:1]
}

func getFileEntries(f *os.File) (Entries, error) {
	reader := bufio.NewReader(f)
	r := make(Entries, 0)

	for {
		var data []byte
		var isPrefix bool
		var err error

		if data, isPrefix, err = reader.ReadLine(); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else if data == nil || isPrefix {
			continue
		}

		split := strings.Split(string(data), " ")

		if len(split) != 2 {
			continue
		}

		var entry *Entry

		// WARNING: If the data in a line is invalid we ignore it and move on
		if entry, err = NewEntry(split[tagIndex], split[pathIndex]); err != nil {
			continue
		}

		r = append(r, entry)
	}

	return r, nil
}
