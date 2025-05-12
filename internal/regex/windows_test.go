//go:build windows

package regex

import "testing"

func TestIsValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ordinary file name",
			args: args{
				s: "file.txt",
			},
			want: true,
		},
		{
			name: "file name with underscore and numbers",
			args: args{
				s: "example_file123",
			},
			want: true,
		},
		{
			name: "file name with dot",
			args: args{
				s: "readme.md",
			},
			want: true,
		},
		{
			name: "file name with Cyrillic characters",
			args: args{
				s: "Файл_тест.doc",
			},
			want: true,
		},
		{
			name: "file name with space inside",
			args: args{
				s: "space ok.txt",
			},
			want: true,
		},
		{
			name: "file name with hyphen",
			args: args{
				s: "my-file-01",
			},
			want: true,
		},
		{
			name: "file name with parentheses",
			args: args{
				s: "test(1).txt",
			},
			want: true,
		},
		{
			name: "file name with Chinese characters",
			args: args{
				s: "新建文档.txt",
			},
			want: true,
		},
		{
			name: "file name with Arabic characters",
			args: args{
				s: "صورة.doc",
			},
			want: true,
		},
		{
			name: "file name with multiple dots",
			args: args{
				s: "file.name.with.dots",
			},
			want: true,
		},
		{
			name: "single digit file name",
			args: args{
				s: "1",
			},
			want: true,
		},
		{
			name: "single letter file name",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "file name with Unicode characters",
			args: args{
				s: "тестовый_документ",
			},
			want: true,
		},
		{
			name: "file name with allowed special characters",
			args: args{
				s: "!test!",
			},
			want: true,
		},
		{
			name: "file name with underscores and numbers",
			args: args{
				s: "data_2025_05_12",
			},
			want: true,
		},
		{
			name: "reserved name 'con'",
			args: args{
				s: "con",
			},
			want: false,
		},
		{
			name: "reserved name 'PRN'",
			args: args{
				s: "PRN",
			},
			want: false,
		},
		{
			name: "reserved name with extension",
			args: args{
				s: "nul.txt",
			},
			want: false,
		},
		{
			name: "reserved name 'COM1'",
			args: args{
				s: "COM1",
			},
			want: false,
		},
		{
			name: "reserved name 'LPT9.txt'",
			args: args{
				s: "LPT9.txt",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '?'",
			args: args{
				s: "file?.txt",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '|'",
			args: args{
				s: "bad|name",
			},
			want: false,
		},
		{
			name: "file name with trailing space",
			args: args{
				s: "name.txt ",
			},
			want: false,
		},
		{
			name: "file name with trailing dot",
			args: args{
				s: "trailingdot.",
			},
			want: false,
		},
		{
			name: "reserved name 'aux'",
			args: args{
				s: "aux",
			},
			want: false,
		},
		{
			name: "file name with forbidden characters '<' and '>'",
			args: args{
				s: "bad<name>.txt",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '\"'",
			args: args{
				s: "quote\".txt",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '*'",
			args: args{
				s: "file*.txt",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '/'",
			args: args{
				s: "slash/name",
			},
			want: false,
		},
		{
			name: "file name with forbidden character '\\'",
			args: args{
				s: "back\\slash",
			},
			want: false,
		},
		{
			name: "file name with forbidden character ':'",
			args: args{
				s: ":colon:",
			},
			want: false,
		},
		{
			name: "empty file name",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "file name with null byte",
			args: args{
				s: string([]byte{0}),
			},
			want: false,
		},
		{
			name: "file name with control character",
			args: args{
				s: string([]byte{31}),
			},
			want: false,
		},
		{
			name: "file name with newline character",
			args: args{
				s: "file\nname.txt",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidate(tt.args.s); got != tt.want {
				t.Errorf("IsValidate(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
