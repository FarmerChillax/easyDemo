package utils

import "testing"

func TestCheckURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "http-case",
			args:    args{url: "http://blog.farmer233.top"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "https-case",
			args:    args{url: "https://blog.farmer233.top"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "false-domain-case",
			args:    args{url: "https://test-url"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "false-case",
			args:    args{"testString"},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
