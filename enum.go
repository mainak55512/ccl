package ccl

import (
	"fmt"
	"strings"
)

type Enum struct {
	_id    int
	_items map[string]int
}

func CreateEnum() Enum {
	return Enum{
		_id:    1,
		_items: make(map[string]int),
	}
}

func (e *Enum) Add(element string) {
	_, ok := e._items[element]
	if !ok {
		e._items[element] = e._id
	}
	e._id++
}

// USAGE:
//
//	en := CreateEnum()
//	en.Add("SUCCESS")
//	en.Add("FAILURE")
//	en.Add("PENDING")
//	en.Add("IN_PROGRESS")
//
//	status := "UNKNOWN"
//
//	state, err := en.Match(status, map[string]func() any{
//		"SUCCESS": func() any { return 1 },
//		"FAILURE": func() any { return 2 },
//		// "PENDING":     func() any { return 3 },
//		"IN_PROGRESS": func() any { return 4 },
//		"_DEFAULT_":   func() any { return 0 },
//	})
//
//	if err != nil {
//		fmt.Println(err) // => Following branches are unattended: PENDING
//	} else {
//
//		fmt.Println("State:", state)
//	}
func (e *Enum) Match(item string, matchObj map[string]func() any) (any, error) {
	var unattendedBranches []string
	for k := range e._items {
		_, ok := matchObj[k]
		if !ok {
			unattendedBranches = append(unattendedBranches, k)
		}
	}
	if len(unattendedBranches) > 0 {
		errmsg := "Following branches are unattended: "
		ForEach(unattendedBranches, func(e string) { errmsg += e + "," })
		return nil, fmt.Errorf("%s", strings.TrimSuffix(errmsg, ","))
	} else {
		fn, ok := matchObj[item]
		if ok {
			return fn(), nil
		} else {
			fn, ok := matchObj["_DEFAULT_"]
			if ok {
				return fn(), nil
			}
		}
	}
	return nil, nil
}
