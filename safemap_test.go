// Package orderedmap provides a map where the order of items is maintained.
// Furthermore, access to contained data is done in a way that is protected and
// concurrent.
package safemap

import (
	"reflect"
	"testing"
)

type TestData struct {
	ID   int
	Name string
}

func TestNewOrderedMap(t *testing.T) {
	sm := New()
	if reflect.TypeOf(sm).Name() != "SafeMap" {

		t.Error("Map is not the correct type")
	}

	if sm.Count() != 0 {
		t.Error("New map is not empty")
	}
}

func TestAdd(t *testing.T) {
	sm := New()
	one := TestData{ID: 1, Name: "one"}
	two := TestData{ID: 2, Name: "two"}

	sm.Add("one", one)
	sm.Add("two", two)

	if sm.Count() != 2 {
		t.Error("Map does not contain two items")
	}
}

func TestGet(t *testing.T) {
	sm := New()

	sm.Add("one", TestData{ID: 1, Name: "one"})
	sm.Add("two", TestData{ID: 2, Name: "two"})
	sm.Add("three", TestData{ID: 3, Name: "three"})

	test, ok := sm.Get("two")
	gotten := test.(TestData)

	if !ok {
		t.Error("Unable to get item frsm map by key")
	}

	if gotten.ID != 2 || gotten.Name != "two" {
		t.Error("Wrong item was returned frsm map")
	}
}

func TestDelete(t *testing.T) {
	sm := New()
	sm.Add("one", TestData{ID: 1, Name: "one"})
	sm.Add("two", TestData{ID: 2, Name: "two"})
	sm.Add("three", TestData{ID: 3, Name: "three"})

	sm.Delete("two")
	_, ok := sm.Get("two")
	if ok {
		t.Error("Deleted key still exists")
	}
	if sm.Count() != 2 {
		t.Error("Size of ordered map was wrong")
	}
}

func TestCount(t *testing.T) {
	sm := New()
	sm.Add("one", TestData{ID: 1, Name: "one"})
	if sm.Count() != 1 {
		t.Error("First count was wrong")
	}
	sm.Add("two", TestData{ID: 2, Name: "two"})
	if sm.Count() != 2 {
		t.Error("Second count was wrong")
	}
	sm.Add("three", TestData{ID: 3, Name: "three"})
	if sm.Count() != 3 {
		t.Error("Third count was wrong")
	}
}

func TestRange(t *testing.T) {
	sm := New()
	sm.Add("one", TestData{ID: 1, Name: "one"})
	sm.Add("two", TestData{ID: 2, Name: "two"})
	sm.Add("three", TestData{ID: 3, Name: "three"})

	for k, v := range sm.Range() {
		gotten := v.(TestData)
		if k == "one" {
			if gotten.ID != 1 || gotten.Name != "one" {
				t.Error("Range item for 'one' was wrong")
			}
		}
		if k == "two" {
			if gotten.ID != 2 || gotten.Name != "two" {
				t.Error("Range item for 'one' was wrong")
			}
		}
		if k == "three" {
			if gotten.ID != 3 || gotten.Name != "three" {
				t.Error("Range item for 'one' was wrong")
			}
		}
	}
}
