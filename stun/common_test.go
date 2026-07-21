package stun

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestListenAndDialCanShareLocalEndpoint(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	remote, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer remote.Close()
	setListenerDeadline(t, remote)

	local, err := ListenTcp(ctx, "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer local.Close()
	setListenerDeadline(t, local)

	outbound, err := DialTcp(local.Addr().String(), remote.Addr().String(), 2)
	if err != nil {
		t.Fatalf("dial from listener endpoint: %v", err)
	}
	defer outbound.Close()

	acceptedOutbound, err := remote.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer acceptedOutbound.Close()

	callback, err := net.DialTimeout("tcp4", local.Addr().String(), 2*time.Second)
	if err != nil {
		t.Fatalf("callback dial: %v", err)
	}
	defer callback.Close()

	acceptedCallback, err := local.Accept()
	if err != nil {
		t.Fatalf("accept callback: %v", err)
	}
	acceptedCallback.Close()
}

func setListenerDeadline(t *testing.T, listener net.Listener) {
	t.Helper()
	tcpListener, ok := listener.(*net.TCPListener)
	if !ok {
		t.Fatalf("listener type is %T, want *net.TCPListener", listener)
	}
	if err := tcpListener.SetDeadline(time.Now().Add(3 * time.Second)); err != nil {
		t.Fatal(err)
	}
}
