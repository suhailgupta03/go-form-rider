package FormRider

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

const (
	methodCannotBeEmpty   = "method cannot be empty"
	incorrectMethodType   = "method can be post or put"
	incorrectURLFormat    = "url must begin with http or https"
	incorrectFieldsFormat = "form fields must be a struct type"
)

// Request sends a multipart/form-data request at the provided url. Accepts
// post or put
func Request(url, method string, fields interface{}) (*http.Response, error) {
	if method == "" {
		return nil, errors.New(methodCannotBeEmpty)
	}

	if strings.ToLower(method) != "post" && strings.ToLower(method) != "put" {
		return nil, errors.New(incorrectMethodType)
	}

	matched, err := regexp.MatchString(`^(http|https)://`, url)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, errors.New(incorrectURLFormat)
	}

	formFields := reflect.ValueOf(fields)
	formFieldsType := formFields.Type()
	if formFields.Kind() != reflect.Struct {
		return nil, errors.New(incorrectFieldsFormat)
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	for i := 0; i < formFields.NumField(); i++ {
		fValue := formFields.Field(i)
		fName := formFieldsType.Field(i).Name
		writer.WriteField(fName, fValue.String())
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the request
	request, err := http.NewRequest(method, url, &requestBody)
	if err != nil {
		return nil, err
	}

	// Setting up of header, this is important!
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Do the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}
