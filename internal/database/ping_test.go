//go:build development

package database_test

import "testing"

// TestPing is to test the Ping function.
func TestPing(t *testing.T) {

	result, err := c.Ping()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)

}
