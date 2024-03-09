package FormRider

import (
	"testing"
)

func TestRequest(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		method   string
		fields   interface{}
		wantErr  bool
		errorMsg string
	}{
		{"Empty Method", "http://example.com", "", nil, true, methodCannotBeEmpty},
		{"Invalid Method", "http://example.com", "GET", nil, true, incorrectMethodType},
		{"Correct Method Case Sensitivity", "http://example.com", "post", struct{}{}, false, ""},
		{"Invalid URL Format", "example.com", "POST", struct{}{}, true, incorrectURLFormat},
		{"Valid URL but Incorrect Protocol", "ftp://example.com", "POST", struct{}{}, true, incorrectURLFormat},
		{"Non-Struct Fields", "http://example.com", "POST", "string", true, incorrectFieldsFormat},
		{"Empty Struct Fields", "http://example.com", "POST", struct{}{}, false, ""},
		{"Valid Request", "http://example.com", "POST", struct{ Name string }{"John"}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Request(tt.url, tt.method, tt.fields)
			if (err != nil) != tt.wantErr || (err != nil && err.Error() != tt.errorMsg) {
				t.Errorf("Request() error = %v, wantErr %v, errorMsg %v", err, tt.wantErr, tt.errorMsg)
			}
		})
	}
}
