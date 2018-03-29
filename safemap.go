/*
Provides a map that can be used in a protected and concurrent fashion.
Map key must be a string, but the data can be anything.

Let's start by first getting a new safemap and adding a few new things to it.
You must get a map by using the New() function, and can add things to it using
Add(Key, Value):

	sm := safemap.New()
	sm.Add("one", 1)
	sm.Add("two", 2)

Now we can access our data through the Get function:

	data := sm.Get("one")

Furthermore, we can iterate on the data safely if we would like as well, knowing
that we will only get a copy of the data in the state it was in when we started



*/
package safemap

import (
	"sync"
)

// A map structure that stores data within a protected fashion.
type SafeMap struct {
	data map[string]interface{}
	lock sync.RWMutex
}

// Create a new safemap struct
func New() SafeMap {
	return SafeMap{
		data: make(map[string]interface{}),
	}
}

// Add an object onto the end of the map
func (m *SafeMap) Add(key string, value interface{}) {
	m.lock.Lock()
	m.data[key] = value
	m.lock.Unlock()
}

// Get an item out of the map based on its key.  In the event the
// key does not exist or the data is out of range, the function will have a
// second return of false.
//
// Get key:
// 	data, ok := om.Get("mykey")
//
// Test if key exists:
// 	if _, ok := om.Get("mykey"); ok {
// 		... DO SOMETHING HERE ...
// 	}
func (m SafeMap) Get(key string) (interface{}, bool) {
	m.lock.RLock()
	data, ok := m.data[key]
	m.lock.RUnlock()
	return data, ok
}

// Delete a specific key and all associated data from the map
func (m *SafeMap) Delete(key string) {
	m.lock.Lock()
	delete(m.data, key)
	m.lock.Unlock()
}

// Get the total size of the map
func (m SafeMap) Count() int {
	m.lock.RLock()
	cnt := len(m.data)
	m.lock.RUnlock()
	return cnt
}

// Get a copy of the map for use in a range
func (m SafeMap) Range() map[string]interface{} {
	tmp := make(map[string]interface{})
	m.lock.RLock()
	// We have to loop through our original to ensure we copy the data
	for k, v := range m.data {
		tmp[k] = v
	}
	m.lock.RUnlock()
	return tmp
}
