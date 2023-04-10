package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestYamlDeepSetKey(t *testing.T) {
	type args struct {
		field string
		data  any
	}
	tests := []struct {
		name string
		args args
		want any
		err  error
	}{
		{
			name: "ok",
			args: args{
				field: "name->duc",
				data:  "duc",
			},
			want: map[string]any{
				"name": map[string]any{
					"duc": "duc",
				},
			},
			err: nil,
		},
		{
			name: "ok2",
			args: args{
				field: "name",
				data:  "duc",
			},
			want: map[string]any{
				"name": "duc",
			},
			err: nil,
		},
		{
			name: "ok2",
			args: args{
				field: "name->duc->a->b",
				data:  "duc",
			},
			want: map[string]any{
				"name": map[string]any{
					"duc": map[string]any{
						"a": map[string]any{
							"b": "duc",
						},
					},
				},
			},
			err: nil,
		},
		{
			name: "fail",
			args: args{
				field: "name->duc->aaaa->",
				data:  "duc",
			},
			err: ErrorInvalidSeparator,
		},
		{
			name: "fail",
			args: args{
				field: "->name->duc->aaaa",
				data:  "duc",
			},
			err: ErrorInvalidSeparator,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			out, _ := yaml.Marshal(tt.want)
			want := string(out)
			got, err := YamlDeepSetKey(tt.args.field, tt.args.data)
			assert.ErrorIs(t, err, tt.err)
			if err == nil {
				assert.Equal(t, want, string(got))
			}
		})
	}
}

func Test_deepGet(t *testing.T) {
	var tests = []struct {
		input   map[string]any
		key     string
		wants   bool
		wantRes any
	}{
		{
			input: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": map[string]any{
							"d": "d",
						},
					},
				},
			},
			key:   "a->b->c",
			wants: true,
			wantRes: map[string]any{
				"d": "d",
			},
		},
		{
			input: map[string]any{
				"a": map[string]any{
					"b": map[string]any{},
				},
			},
			key:     "a->b->c",
			wants:   false,
			wantRes: nil,
		},
		{
			input: map[string]any{
				"a": map[string]any{
					"b": map[string]any{},
				},
			},
			key:     "",
			wants:   false,
			wantRes: nil,
		},
		{
			input: map[string]any{
				"a": map[any]any{
					"b": map[string]any{},
				},
			},
			key:     "",
			wants:   false,
			wantRes: nil,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.key, func(t *testing.T) {
			t.Parallel()
			res, b := deepGet(tt.key, tt.input)
			assert.Equal(t, tt.wants, b)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
