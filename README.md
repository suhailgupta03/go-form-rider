# go-form-rider

Simple multipart form request creator. 

To use with google sheets, set using the steps explained [html-form-to-google-sheet](https://github.com/levinunnink/html-form-to-google-sheet)
and then can be integrated with [thunderbyte](https://github.com/suhailgupta03/thunderbyte) or used as a separate utility to upload the form data

```go
	type formData struct {
		Email string
		Name  string
	}

	// Replace with your own action URL
	resp, err := FormRider.Request(
		"https://script.google.com/macros/s/AXXXcbwLvbgyuD99IFi-4MSCMLL_oyAUamJKlc_Q1sU7Sl8b4OfGoX_RVwVcDx7SKCHZIYRXpk/exec",
		"post",
		formData{
			Email: "foo@bar.foobar",
			Name:  "foobar",
		},
	)

	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf(resp.Status)
	}
```