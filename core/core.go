package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	tagIndex  = 0
	pathIndex = 1
)

func MapFileEntries(f *os.File) (map[string]*Entry, error) {
	reader := bufio.NewReader(f)
	m := make(map[string]*Entry)

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

		m[entry.id] = entry
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
