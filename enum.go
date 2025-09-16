package ccl

import (
	"fmt"
	"strings"
)

type Enum struct {
	_id      int
	_freezed bool
	_items   map[string]any
}

func CreateEnum() Enum {
	return Enum{
		_id:      1,
		_freezed: false,
		_items:   make(map[string]any),
	}
}

func (e *Enum) Freeze() {
	e._freezed = true
}

func (e *Enum) Add(element string) {
	if !e._freezed {
		_, ok := e._items[element]
		if !ok {
			e._id++
			e._items[element] = e._id
		}
	}
}

func (e *Enum) AddWithValue(element string, value any) {
	if !e._freezed {
		_, ok := e._items[element]
		if !ok {
			e._items[element] = value
			val, ok := value.(int)
			if ok {
				e._id = val
			}
		}
	}
}

func (e Enum) Varient(value any) (string, error) {
	item, ok := value.(string)
	if ok {
		_, ok := e._items[item]
		if ok {
			return item, nil
		} else {
			for k := range e._items {
				val, ok := e._items[k]
				if ok && val == value {
					return k, nil
				}
			}
		}
	} else {
		for k := range e._items {
			val, ok := e._items[k]
			if ok && val == value {
				return k, nil
			}
		}
	}
	return "", fmt.Errorf("Enum varient doesn't exist")
}

/*
Usage:
------

	en := CreateEnum()
	en.AddWithValue("SUCCESS", 201)
	en.AddWithValue("FAILURE", 404)
	en.AddWithValue("PENDING", 312)
	en.AddWithValue("IN_PROGRESS", 240)
	en.Freeze()

	status := 404 // or, "FAILURE"

	state, err := en.Match(status, map[string]func() any{
		"SUCCESS":     func() any { return 1 },
		"FAILURE":     func() any { return 2 },
		"PENDING":     func() any { return 3 },
		"IN_PROGRESS": func() any { return 4 },
		"_DEFAULT_":   func() any { return 0 },
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("State:", state) // State: 2
	}
*/
func (e Enum) Match(item any, matchObj map[string]func() any) any {
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
		panic(fmt.Errorf("%s", strings.TrimSuffix(errmsg, ",")))
	} else {
		varient, _ := e.Varient(item)
		fn, ok := matchObj[varient]
		if ok {
			return fn()
		} else {
			fn, ok := matchObj["_DEFAULT_"]
			if ok {
				return fn()
			}
		}
	}
	return nil
}
