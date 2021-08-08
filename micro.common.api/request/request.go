package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

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

	err = validate.Struct(reqPointer)
	if err != nil {
		return err
	}

	return nil
}
