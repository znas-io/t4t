package core

import "sort"

type (
	Entries []*Entry

	SortedEntriesMap struct {
		entriesMap map[string]*Entry
		entries    Entries
	}
)

func NewSortedEntriesMap() *SortedEntriesMap {
	return &SortedEntriesMap{
		entriesMap: make(map[string]*Entry),
		entries:    make([]*Entry, 0),
	}
}

func (s *SortedEntriesMap) Add(entry *Entry) error {
	if entry == nil {
		return ErrEntryCannotBeNil
	}

	var e *Entry
	var ok bool

	if e, ok = s.entriesMap[entry.id]; ok {
		if e.tag != entry.tag || e.path != entry.path {
			return ErrUnthinkable(e.tag, e.path, entry.tag, entry.path)
		}
	} else {
		s.entriesMap[entry.id] = entry
		s.entries = append(s.entries, entry)
	}

	return nil
}

func (s *SortedEntriesMap) GetEntries() []*Entry {
	sort.Sort(s.entries)
	return s.entries
}

func (t Entries) Len() int {
	return len(t)
}

func (t Entries) Less(i int, j int) bool {
	return t[i].tag < t[j].tag
}

func (t Entries) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}
