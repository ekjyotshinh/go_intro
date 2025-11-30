package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJSONBody(r *http.Request, dst interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}
	return nil
}