package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ErrNonSuccessResponse struct {
	Response http.Response
}

func (e ErrNonSuccessResponse) Error() string {
	bodyBytes, _ := ioutil.ReadAll(e.Response.Body)
	msg := fmt.Sprintf("Non-success status code with body %s", string(bodyBytes))
	return msg
}
