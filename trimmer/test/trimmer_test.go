package test

import (
	"reflect"
	"testing"

	"github.com/requiemofthesouls/logger/trimmer"
)

func Test_trimmer_Trim(t1 *testing.T) {
	type args struct {
		handler string
		text    []byte
	}
	tests := []struct {
		name          string
		trimmedFields trimmer.TrimmedFields
		args          args
		want          []byte
	}{
		{
			name:          "empty handlers",
			trimmedFields: map[string][]string{},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "empty trimmed fields in handler",
			trimmedFields: map[string][]string{
				"handler": {},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "int field",
			trimmedFields: map[string][]string{
				"handler": {"a"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": "TRIMMED_CONTENT", "b": "2", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "string field",
			trimmedFields: map[string][]string{
				"handler": {"b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "TRIMMED_CONTENT", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "bool field",
			trimmedFields: map[string][]string{
				"handler": {"c"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c": "TRIMMED_CONTENT", "d": [1, 2, 3]}`),
		},
		{
			name: "slice field",
			trimmedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "composite type slice field",
			trimmedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [{"aa": 1}, {"aa": 33}]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "several fields",
			trimmedFields: map[string][]string{
				"handler": {"a", "d", "b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": "TRIMMED_CONTENT", "b": "TRIMMED_CONTENT", "c":true, "d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "several handlers",
			trimmedFields: map[string][]string{
				"handler1": {"c", "a", "d", "b"},
				"handler2": {"d"},
			},
			args: args{
				handler: "handler2",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1,"b": "2","c":true,"d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "camel_case handler",
			trimmedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "Handler",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1,"b": "2","c":true,"d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "same fields in text",
			trimmedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1,"b": "2","body": {"d":true},"d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1,"b": "2","body": {"d": "TRIMMED_CONTENT"},"d": "TRIMMED_CONTENT"}`),
		},
		{
			name: "empty text",
			trimmedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte{},
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := trimmer.New(tt.trimmedFields)
			if got := t.Trim(tt.args.handler, tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
