package httputils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func DecodeJSON(r *http.Request, v interface{}) error {
	defer io.Copy(ioutil.Discard, r.Body)
	return json.NewDecoder(r.Body).Decode(v)
}

func IsJsonRequest(r *http.Request) bool {
	return strings.TrimSpace(strings.Split(r.Header.Get("Content-Type"), ";")[0]) == "application/json"
}
