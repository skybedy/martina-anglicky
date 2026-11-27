package main

import "testing"

func TestServerAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
	}{
		{
			name: "default",
			want: defaultAddress,
		},
		{
			name:    "configured",
			address: "127.0.0.1:18081",
			want:    "127.0.0.1:18081",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ADDR", tt.address)

			if got := serverAddress(); got != tt.want {
				t.Errorf("serverAddress() = %q, want %q", got, tt.want)
			}
		})
	}
}
