// Package escape_test is to test the escape package.
package escape_test

import (
	"github.com/gowizzard/trigsnfuncs/internal/escape"
	"reflect"
	"testing"
)

// TestBytes is to test the Bytes function.
func TestBytes(t *testing.T) {

	tests := []struct {
		name     string
		v        []byte
		expected string
	}{
		{
			name:     "add_last_updated.js",
			v:        []byte("#!js name=play api_version=1.0\n\nfunction ping(client, data) {\n    return client.call('ping');\n}\n\nredis.registerFunction('ping', ping);"),
			expected: "\"#!js name=play api_version=1.0\\n\\nfunction ping(client, data) {\\n    return client.call('ping');\\n}\\n\\nredis.registerFunction('ping', ping);\"",
		},
		{
			name:     "ping.js",
			v:        []byte("#!js name=trigger api_version=1.0\n\nfunction setLastUpdate(client, data) {\n    if(data.event == \"hset\") {\n        const time = new Date();\n        client.call(\"hset\", data.key, \"last_update\", time.toISOString());\n    }\n}\n\nredis.registerKeySpaceTrigger(\"addLastUpdated\", \"shop:\", setLastUpdate);"),
			expected: "\"#!js name=trigger api_version=1.0\\n\\nfunction setLastUpdate(client, data) {\\n    if(data.event == \\\"hset\\\") {\\n        const time = new Date();\\n        client.call(\\\"hset\\\", data.key, \\\"last_update\\\", time.toISOString());\\n    }\\n}\\n\\nredis.registerKeySpaceTrigger(\\\"addLastUpdated\\\", \\\"shop:\\\", setLastUpdate);\"",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := escape.Bytes(tt.v)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Bytes() = %v, want %v", result, tt.expected)
			}

		})

	}

}
