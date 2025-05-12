package regex

import "testing"

func TestIsValidateLinux(t *testing.T) {
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
			name: "hidden file",
			args: args{
				s: ".hidden",
			},
			want: true,
		},
		{
			name: "file name with Cyrillic characters",
			args: args{
				s: "файл_с_именем",
			},
			want: true,
		},
		{
			name: "file name with Japanese characters",
			args: args{
				s: "テストファイル",
			},
			want: true,
		},
		{
			name: "file name with emoji",
			args: args{
				s: "emoji_📁_test",
			},
			want: true,
		},
		{
			name: "file name with spaces",
			args: args{
				s: "file name with space",
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
			name: "file name with special character",
			args: args{
				s: "!",
			},
			want: true,
		},
		{
			name: "file name with multiple special characters",
			args: args{
				s: "@#$%^&()[]{}",
			},
			want: true,
		},
		{
			name: "file name with additional special characters",
			args: args{
				s: "=+,.~`",
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
				s: "صورة",
			},
			want: true,
		},
		{
			name: "file name with slash",
			args: args{
				s: "bad/name",
			},
			want: false,
		},
		{
			name: "file path instead of name",
			args: args{
				s: "folder/file",
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
			name: "file name with null character",
			args: args{
				s: "null\000byte",
			},
			want: false,
		},
		{
			name: "file name with single slash",
			args: args{
				s: "/",
			},
			want: false,
		},
		{
			name: "file name ending with slash",
			args: args{
				s: "slash_at_end/",
			},
			want: false,
		},
		{
			name: "file name starting with slash",
			args: args{
				s: "/starts_with",
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
			if got := linuxIsValidate(tt.args.s); got != tt.want {
				t.Errorf("IsValidate(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
