package network

import (
	"fmt"
	"net/http"
	"reflect"
)

var (
	ErrInvalidData = fmt.Errorf("data must be a struct or a map")
)

type RequestBuilder struct{}

func (rb RequestBuilder) BuildGetRequest(url string, data interface{}) (req *http.Request, err error) {
	t := reflect.TypeOf(data)
	if kind := t.Kind(); kind != reflect.Struct && kind != reflect.Map {
		err = ErrInvalidData
		return
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	v := reflect.ValueOf(data)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	numFields := reflect.TypeOf(data).NumField()
	typeOfData := reflect.TypeOf(data)
	for i := 0; i < numFields; i++ {
		var key string
		key = typeOfData.Field(i).Name
		if tag, ok := typeOfData.Field(i).Tag.Lookup("json"); ok {
			key = tag
		}
		field := v.Field(i)

		kind := field.Type().Kind()
		if kind == reflect.Struct {
			err = ErrInvalidData
			return
		}

		if kind == reflect.Slice {
			switch field.Interface().(type) {
			case []string:
				s := field.Interface().([]string)
				for _, elem := range s {
					if elem == "" {
						continue // exclude empty string params
					}
					q.Add(key, elem)
				}
				break

			default:
				panic(fmt.Sprintf("unsupported type %s given to build request", field.Type().String()))
			}
			continue
		}

		val := fmt.Sprintf("%v", field.Interface())
		if val == "" {
			continue // exclude empty params
		}

		q.Add(key, val)
	}

	req.URL.RawQuery = q.Encode()
	return
}
