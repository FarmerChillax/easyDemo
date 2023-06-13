package time

import "testing"

func TestGetLastMonthTimestamp(t *testing.T) {
	tests := []struct {
		name  string
		want  int64
		want1 int64
	}{
		// TODO: Add test cases.
		{
			name:  "test",
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetLastMonthTimestamp()
			t.Log(got)
			t.Log(got1)
			// if got != tt.want {
			// 	t.Errorf("GetLastMonthTimestamp() got = %v, want %v", got, tt.want)
			// }
			// if got1 != tt.want1 {
			// 	t.Errorf("GetLastMonthTimestamp() got1 = %v, want %v", got1, tt.want1)
			// }
		})
	}
}
