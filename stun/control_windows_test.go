//go:build windows

package stun

import "testing"

func TestValidateLocalAddrOnWindows(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{name: "specific IPv4", address: "192.0.2.10:0"},
		{name: "wildcard IPv4", address: "0.0.0.0:0", wantErr: true},
		{name: "wildcard IPv6", address: "[::]:0", wantErr: true},
		{name: "IPv6", address: "[2001:db8::1]:0", wantErr: true},
		{name: "hostname", address: "localhost:0", wantErr: true},
		{name: "missing port", address: "192.0.2.10", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateLocalAddr(test.address)
			if test.wantErr && err == nil {
				t.Fatal("expected an error")
			}
			if !test.wantErr && err != nil {
				t.Fatal(err)
			}
		})
	}
}
