package valid_parentheses

import "container/list"

var m = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	l := list.New()
	for _, c := range s {
		if v, ok := m[c]; !ok {
			l.PushBack(c)
		} else {
			if l.Back() == nil || l.Back().Value != v {
				return false
			}
			l.Remove(l.Back())
		}
	}
	return l.Len() == 0
}
