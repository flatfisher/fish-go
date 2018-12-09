package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string]int
	}{
		{
			name: "simple",
			input: []string{
				"foo",
				"bar",
				"foo",
			},
			want: map[string]int{
				"foo": 2,
				"bar": 1,
			},
		},
		{
			name:  "empty",
			input: []string{},
			want:  map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordCount(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %#v, but got %#v", tt.want, got)
			}
		})
	}

}
