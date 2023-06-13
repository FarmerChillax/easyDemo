package time

import (
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	type args struct {
		fn func()
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
		{
			name: "exec time",
			args: args{
				func() {
					time.Sleep(time.Second)
				},
			},
			want: time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exec(tt.args.fn); got <= tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
