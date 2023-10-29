//go:build development

package database_test

import (
	"testing"
)

// TestTFunctionLoad is to test the TFunctionLoad function.
func TestTFunctionLoad(t *testing.T) {

	response, err := c.TFunctionLoad("\"#!js api_version=1.0 name=lib\\n redis.registerFunction('hello', ()=>{return 'Hello world'})\"")
	if err != nil {
		t.Error(err)
	}

	t.Log(response)

}
