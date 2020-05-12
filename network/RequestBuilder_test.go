package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Param1 string `json:"param_1"`
	Param2 int
	Param3 []string
}

type Bar struct {
	Param1 Foo `json:"foo"`
	Param2 string
	Param3 int
}

func TestRequestBuilder_BuildGetRequest(t *testing.T) {
	f := Foo{"lorem ipsum", 120, []string{"item1", "item2"}}

	rb := RequestBuilder{}
	req, err := rb.BuildGetRequest("https://example.com", f)
	expectedQuery := "https://example.com?Param2=120&Param3=item1&Param3=item2&param_1=lorem+ipsum"
	assert.Nil(t, err)
	assert.Equal(t, expectedQuery, req.URL.String())
}

func TestRequestBuilder_BuildGetRequest_NoEmptyParams(t *testing.T) {
	f := Foo{Param2: 120, Param3: []string{"item1", "item2"}}

	rb := RequestBuilder{}
	req, err := rb.BuildGetRequest("https://example.com", f)
	expectedQuery := "https://example.com?Param2=120&Param3=item1&Param3=item2"
	assert.Nil(t, err)
	assert.Equal(t, expectedQuery, req.URL.String())
}

func TestRequestBuilder_BuildGetRequest_InvalidData(t *testing.T) {
	f := Foo{"lorem ipsum", 120, []string{"item1", "item2"}}
	b := Bar{
		Param1: f,
		Param2: "bar",
		Param3: 0,
	}

	rb := RequestBuilder{}
	_, err := rb.BuildGetRequest("blah", b)
	assert.Equal(t, ErrInvalidData, err)
}
