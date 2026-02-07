package util

type Set struct {
	item map[any]struct{}
}

func SetNew() Set {
	return Set{
		item: make(map[any]struct{}),
	}
}

func (s *Set) Add(entry []any) {
	for _, x := range entry {
		s.item[x] = struct{}{}
	}
}

func (s *Set) Del(entry any) {
	delete(s.item, entry)
}

func (s *Set) Check(entry any) bool {
	_, isOk := s.item[entry]
	return isOk
}
