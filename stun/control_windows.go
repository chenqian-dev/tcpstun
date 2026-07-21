//go:build windows

package stun

import (
	"fmt"
	"net"
	"syscall"

	"golang.org/x/sys/windows"
)

func Control(_ string, _ string, raw syscall.RawConn) (err error) {
	if controlErr := raw.Control(func(fd uintptr) {
		err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
	}); controlErr != nil {
		return controlErr
	}
	return err
}

func validateLocalAddr(address string) error {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("invalid Windows local address %q: %w", address, err)
	}
	ip := net.ParseIP(host)
	if ip == nil || ip.To4() == nil || ip.IsUnspecified() {
		return fmt.Errorf("Windows local address must use a specific IPv4 address: %q", address)
	}
	return nil
}
