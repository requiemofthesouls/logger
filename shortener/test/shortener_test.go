package test

import (
	"reflect"
	"testing"

	"github.com/requiemofthesouls/logger/shortener"
)

func Test_shortener_Shorten(t *testing.T) {
	type args struct {
		handler string
		text    []byte
	}
	tests := []struct {
		name         string
		loggedFields shortener.LoggedFields
		args         args
		want         []byte
	}{
		{
			name:         "empty handlers",
			loggedFields: map[string][]string{},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "empty loggedFields in handler",
			loggedFields: map[string][]string{
				"handler": {},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte{},
		},
		{
			name: "empty loggedFields in another handler",
			loggedFields: map[string][]string{
				"another_handler": {},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
		},
		{
			name: "int field",
			loggedFields: map[string][]string{
				"handler": {"a"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1}`),
		},
		{
			name: "string field",
			loggedFields: map[string][]string{
				"handler": {"b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"b": "2"}`),
		},
		{
			name: "bool field",
			loggedFields: map[string][]string{
				"handler": {"c"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"c":true}`),
		},
		{
			name: "slice field",
			loggedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [1, 2, 3]}`),
			},
			want: []byte(`{"d": [1, 2, 3]}`),
		},
		{
			name: "composite type slice field",
			loggedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1, "b": "2", "c":true, "d": [{"aa": 1}, {"aa": 33}]}`),
			},
			want: []byte(`{"d": [{"aa": 1}, {"aa": 33}]}`),
		},
		{
			name: "several fields",
			loggedFields: map[string][]string{
				"handler": {"a", "d", "b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"a": 1, "d": [1, 2, 3], "b": "2"}`),
		},
		{
			name: "all fields in other order",
			loggedFields: map[string][]string{
				"handler": {"c", "a", "d", "b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"c":true, "a": 1, "d": [1, 2, 3], "b": "2"}`),
		},
		{
			name: "another fields in text",
			loggedFields: map[string][]string{
				"handler": {"c", "a", "d", "b"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"aaa": 1,"bbb": "2","ccc":true,"ddd": [1, 2, 3]}`),
			},
			want: []byte(`{"aaa": 1,"bbb": "2","ccc":true,"ddd": [1, 2, 3]}`),
		},
		{
			name: "several handlers",
			loggedFields: map[string][]string{
				"handler1": {"c", "a", "d", "b"},
				"handler2": {"d"},
			},
			args: args{
				handler: "handler2",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"d": [1, 2, 3]}`),
		},
		{
			name: "camel_case handler",
			loggedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "Handler",
				text:    []byte(`{"a": 1,"b": "2","c":true,"d": [1, 2, 3]}`),
			},
			want: []byte(`{"d": [1, 2, 3]}`),
		},
		{
			name: "same fields in text",
			loggedFields: map[string][]string{
				"handler": {"d"},
			},
			args: args{
				handler: "handler",
				text:    []byte(`{"a": 1,"b": "2","body": {"d":true},"d": [1, 2, 3]}`),
			},
			want: []byte(`{"d":true, "d": [1, 2, 3]}`),
		},
		{
			name: "empty text",
			loggedFields: map[string][]string{
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
		t.Run(tt.name, func(t *testing.T) {
			s := shortener.New(tt.loggedFields)
			if got := s.Shorten(tt.args.handler, tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
