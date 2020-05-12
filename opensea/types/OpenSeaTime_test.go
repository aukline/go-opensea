package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOpenSeaTime_UnmarshallJSON(t *testing.T) {
	type Foo struct {
		Time OpenSeaTime `json:"created_at"`
	}
	var f Foo
	js := `{"created_at": "2019-09-05T07:39:47.495758"}`
	expectedDate := time.Date(2019, time.September, 05, 07, 39, 47, 495758000, time.UTC)

	err := json.Unmarshal([]byte(js), &f)
	assert.Nil(t, err)
	assert.Equal(t, expectedDate, f.Time.Time)
}
