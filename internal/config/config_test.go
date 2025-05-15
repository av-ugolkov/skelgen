package config

import (
	"testing"
)

func Test_replacePlaceholders(t *testing.T) {
	type args struct {
		input        string
		replaceValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "replace placeholder",
			args: args{
				input:        "Hello, ${name}!",
				replaceValue: "World",
			},
			want: "Hello, World!",
		},
		{
			name: "no placeholder",
			args: args{
				input:        "Hello, World!",
				replaceValue: "World",
			},
			want: "Hello, World!",
		},
		{
			name: "no replace value",
			args: args{
				input:        "Hello, ${name}!",
				replaceValue: "",
			},
			want: "Hello, !",
		},
		{
			name: "no input",
			args: args{
				input:        "",
				replaceValue: "World",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replacePlaceholders(tt.args.input, tt.args.replaceValue); got != tt.want {
				t.Errorf("replacePlaceholders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractOnePlaceholder(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "extract placeholder",
			args: args{
				input: "Hello, ${name}!",
			},
			want: "name",
		},
		{
			name: "no placeholder",
			args: args{
				input: "Hello, World!",
			},
			want: "",
		},
		{
			name: "no input",
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "invalid input",
			args: args{
				input: "Hello, ${name",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractOnePlaceholder(tt.args.input); got != tt.want {
				t.Errorf("extractOnePlaceholder() = %v, want %v", got, tt.want)
			}
		})
	}
}
