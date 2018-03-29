# SafeMap #
Provides a map that can be used in a protected and concurrent fashion.  Key must
be a string, but data can be anything.

[![GoDoc](https://godoc.org/github.com/emperorcow/safemap?status.svg)](http://godoc.org/github.com/emperorcow/safemap)
[![Coverage Status](https://coveralls.io/repos/emperorcow/safemap/badge.svg?branch=master)](https://coveralls.io/r/emperorcow/safemap?branch=master)

# Usage #
You'll first need to make a new map and add some data.  Create a new map using the New() function and then add some data in using the Add(Key, Val) function, which takes a key as a string, and any data type as the value:
```
sm := safemap.New()
sm.Set("one", TestData{ID: 1, Name: "one"})
sm.Set("two", TestData{ID: 2, Name: "two"})
sm.Set("three", TestData{ID: 3, Name: "three"})
```

You can get data from the map using Get.   

```
datakey, ok := om.Get("two")

```

There are also many other things, you can do, like delete by key, get the size, etc.  See the godoc for more information

