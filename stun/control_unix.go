//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd

package stun

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func Control(_ string, _ string, raw syscall.RawConn) (err error) {
	if controlErr := raw.Control(func(fd uintptr) {
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		if err != nil {
			return
		}
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
	}); controlErr != nil {
		return controlErr
	}
	return err
}

func validateLocalAddr(_ string) error { return nil }
