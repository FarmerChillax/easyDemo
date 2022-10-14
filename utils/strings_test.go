package utils

import "testing"

func TestCheckString2Long(t *testing.T) {
	type args struct {
		target string
		len    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "正确校验",
			args: args{target: "1234567", len: 7},
			want: true,
		}, {
			name: "错误校验",
			args: args{target: "1234567", len: 8},
			want: false,
		}, {
			name: "空字符串测试",
			args: args{target: "", len: 0},
			want: true,
		}, {
			name: "中文字符串测试",
			args: args{target: "一二三", len: 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckString2Long(tt.args.target, tt.args.len); got != tt.want {
				t.Errorf("CheckString2Long() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStringInRange(t *testing.T) {
	type args struct {
		target string
		min    int
		max    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "测试最大值",
			args: args{target: "1234567", max: 6, min: 5},
			want: false,
		}, {
			name: "测试最小值",
			args: args{target: "1234567", max: 10, min: 8},
			want: false,
		}, {
			name: "正常测试",
			args: args{target: "1234567", max: 10, min: 5},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringInRange(tt.args.target, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("CheckStringInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckString2Short(t *testing.T) {
	type args struct {
		target string
		len    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "超出长度",
			args: args{target: "1234567", len: 6},
			want: false,
		}, {
			name: "合法中文字符串",
			args: args{target: "一二三", len: 3},
			want: true,
		}, {
			name: "非法中文字符串",
			args: args{target: "一二三", len: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckString2Short(tt.args.target, tt.args.len); got != tt.want {
				t.Errorf("CheckString2Short() = %v, want %v", got, tt.want)
			}
		})
	}
}
