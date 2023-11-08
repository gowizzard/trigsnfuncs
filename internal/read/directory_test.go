// Package read_test is to test the read package.
package read_test

import (
	"github.com/gowizzard/trigsnfuncs/internal/read"
	"os"
	"testing"
)

// TestDirectory is to test the Directory function.
func TestDirectory(t *testing.T) {

	files := []struct {
		name    string
		content string
	}{
		{
			name:    "ping.js",
			content: "#!js name=trigger api_version=1.0\n\nfunction setLastUpdate(client, data) {\n    if(data.event == \"hset\") {\n        const time = new Date();\n        client.call(\"hset\", data.key, \"last_update\", time.toISOString());\n    }\n}\n\nredis.registerKeySpaceTrigger(\"addLastUpdated\", \"shop:\", setLastUpdate);",
		},
		{
			name:    "add_last_updated.js",
			content: "#!js name=play api_version=1.0\n\nfunction ping(client, data) {\n    return client.call('ping');\n}\n\nredis.registerFunction('ping', ping);",
		},
	}

	path := os.TempDir()
	for _, value := range files {
		err := os.WriteFile(path+"/"+value.name, []byte(value.content), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	_, err := read.Directory(path)
	if err != nil {
		t.Error(err)
	}

}
