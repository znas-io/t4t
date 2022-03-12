package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func MapFileEntries(f *os.File) (map[string]*Entry, error) {
	reader := bufio.NewReader(f)
	m := make(map[string]*Entry)

	for {
		var data []byte
		var isPrefix bool
		var err error

		if data, isPrefix, err = reader.ReadLine(); data == nil {
			break
		} else if isPrefix {
			continue
		} else if err != nil {
			return nil, err
		}

		split := strings.Split(string(data), " ")

		if len(split) != 4 {
			continue
		}

		var entry *Entry

		// WARNING: If the data in a line is invalid we ignore it and move on
		if entry, err = NewEntry(split[2], split[3]); err != nil {
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

	if file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
		return nil, err
	}

	return file, nil
}
