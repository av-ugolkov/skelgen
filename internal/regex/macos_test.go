//go:build darwin

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
			name: "file name with spaces",
			args: args{
				s: "space ok.txt",
			},
			want: true,
		},
		{
			name: "file name with hyphen",
			args: args{
				s: "file-name-123",
			},
			want: true,
		},
		{
			name: "file name with multiple dots",
			args: args{
				s: "file.name.tar.gz",
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
			name: "file name with parentheses and spaces",
			args: args{
				s: "new file (copy)",
			},
			want: true,
		},
		{
			name: "file name with emoji",
			args: args{
				s: "test_✔️_emoji",
			},
			want: true,
		},
		{
			name: "file name with Chinese characters",
			args: args{
				s: "中文文件",
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
			name: "hidden file name",
			args: args{
				s: ".hiddenfile",
			},
			want: true,
		},
		{
			name: "file name with special characters",
			args: args{
				s: "@file$!#%",
			},
			want: true,
		},
		{
			name: "file name with forbidden character ':'",
			args: args{
				s: "bad:name",
			},
			want: false,
		},
		{
			name: "file name with ':' in the middle",
			args: args{
				s: "file:name.txt",
			},
			want: false,
		},
		{
			name: "file name starting with ':'",
			args: args{
				s: ":colonstart",
			},
			want: false,
		},
		{
			name: "file name ending with ':'",
			args: args{
				s: "end:",
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
			name: "file name containing null byte",
			args: args{
				s: "null\000byte",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidate(tt.args.s); got != tt.want {
				t.Errorf("IsValidate(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
