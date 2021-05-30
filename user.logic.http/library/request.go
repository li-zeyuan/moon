package library

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

func ParseBody(r *http.Request, reqPointer interface{}) error {
	if reflect.ValueOf(reqPointer).Kind() != reflect.Ptr {
		return errors.New("request params must pointer")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, reqPointer)
	if err != nil {
		return err
	}

	return nil
}
