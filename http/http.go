package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Do wrap http.Do, and unmarshal.
func Do(req *http.Request, res interface{}) error {
	reply, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer reply.Body.Close()
	b, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, res)
}
