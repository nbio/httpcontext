// Forked from the Gorilla context test:
// https://github.com/gorilla/context/blob/master/context_test.go
// © 2012 The Gorilla Authors

package httpcontext

import (
	"net/http"
	"testing"

	"github.com/nbio/st"
)

type keyType int

const (
	key1 keyType = iota
	key2
)

func TestContext(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	empty, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	crc := getContextReadCloser(req)

	// Get()
	st.Expect(t, Get(req, key1), nil)

	// Set()
	Set(req, key1, "1")
	st.Expect(t, Get(req, key1), "1")
	st.Expect(t, len(crc.Context()), 1)

	Set(req, key2, "2")
	st.Expect(t, Get(req, key2), "2")
	st.Expect(t, len(crc.Context()), 2)

	// GetOk()
	value, ok := GetOk(req, key1)
	st.Expect(t, value, "1")
	st.Expect(t, ok, true)

	value, ok = GetOk(req, "not exists")
	st.Expect(t, value, nil)
	st.Expect(t, ok, false)

	Set(req, "nil value", nil)
	value, ok = GetOk(req, "nil value")
	st.Expect(t, value, nil)
	st.Expect(t, ok, true)

	// GetString()
	Set(req, "int value", 13)
	Set(req, "string value", "hello")
	str := GetString(req, "int value")
	st.Expect(t, str, "")
	str = GetString(req, "string value")
	st.Expect(t, str, "hello")

	// GetAll()
	values := GetAll(req)
	st.Expect(t, len(values), 5)

	// GetAll() for empty request
	values = GetAll(empty)
	st.Expect(t, len(values), 0)

	// Delete()
	Delete(req, key1)
	st.Expect(t, Get(req, key1), nil)
	st.Expect(t, len(crc.Context()), 4)

	Delete(req, key2)
	st.Expect(t, Get(req, key2), nil)
	st.Expect(t, len(crc.Context()), 3)

	// Clear()
	Set(req, key1, true)
	values = GetAll(req)
	Clear(req)
	st.Expect(t, len(crc.Context()), 0)
	val, _ := values["int value"].(int)
	st.Expect(t, val, 13) // Clear shouldn't delete values grabbed before
}
