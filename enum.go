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

func (e *Enum) Items() map[string]any {
	return e._items
}

func (e *Enum) ID() int {
	return e._id
}

func (e *Enum) Freeze() {
	e._freezed = true
}

// Add method auto increments value under the hood
func (e *Enum) Add(name string) {
	if e._freezed {
		return
	}
	if _, exists := e._items[name]; !exists {
		e._items[name] = e._id
		e._id++
	}
}

// AddWithValue lets user specify unique value for each entry
func (e *Enum) AddWithValue(element string, value any) {
	if !e._freezed {
		if _, ok := e._items[element]; !ok {
			e._items[element] = value
			if val, ok := value.(int); ok {
				e._id = val
			}
		}
	}
}

func _variantHelper[T comparable](mapObj map[string]any, value T) string {
	for k := range mapObj {
		if val, ok := mapObj[k]; ok && val == value {
			return k
		}
	}
	return ""
}

// Returns the Variant(string) for the Enum
func (e Enum) Variant(value any) (string, error) {
	if item, ok := value.(string); ok {
		if _, ok := e._items[item]; ok {
			return item, nil
		} else {
			if val := _variantHelper(e._items, value); val != "" {
				return val, nil
			}
		}
	} else {
		if val := _variantHelper(e._items, value); val != "" {
			return val, nil
		}
	}
	return "", fmt.Errorf("Enum variant doesn't exist")
}

/*
Panics if all Enum variants are not handled.

Usage:
------

	en := CreateEnum()
	en.AddWithValue("SUCCESS", 5) // => 5
	en.Add("FAILURE") // => 6
	en.Add("PENDING") // => 7
	en.Add("IN_PROGRESS") // => 8
	en.Freeze()

	status := 6 // or, status := "FAILURE"

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
		if _, ok := matchObj[k]; !ok {
			unattendedBranches = append(unattendedBranches, k)
		}
	}
	if len(unattendedBranches) > 0 {
		errmsg := "Following branches are unattended: "
		ForEach(unattendedBranches, func(e string) { errmsg += e + "," })
		panic(fmt.Errorf("%s", strings.TrimSuffix(errmsg, ",")))
	} else {
		variant, _ := e.Variant(item)
		if fn, ok := matchObj[variant]; ok {
			return fn()
		} else {
			if fn, ok := matchObj["_DEFAULT_"]; ok {
				return fn()
			}
		}
	}
	return nil
}
