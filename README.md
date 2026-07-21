# TCP NAT Type Test

This toolkit uses TCP protocol to detect NAT types. 

It is divided into a server and a client, which work together to detect the specific NAT type of the client's network. 

Currently, it can identify the following 5 types of NAT:
- NAT0: Public IP, no NAT
- NAT1: Full Cone NAT
- NAT2: Restricted Cone NAT
- NAT3: Port Restricted Cone NAT
- NAT4: Symmetric NAT

Note: This toolkit confirms to no RFC, the server and client must be used together.

## Server

The server requires two public IPv4 addresses and two TCP ports. Existing
deployments can continue running the server on Linux.

## Client

Build the client on Windows:

```powershell
Set-Location .\client
.\build.ps1
```

Run it with an explicit server address. The local address is optional; when it
is omitted, the client selects the outbound IPv4 address:

```powershell
.\tcpstunc_windows_amd64.exe -H <server-ip> -P 3478
```

TCP NAT discovery requires the server to connect back to the client's
temporary listening port. Windows Firewall must therefore have an inbound TCP
allow rule for the client executable, preferably scoped to the TCP STUN server
addresses. The client does not change firewall policy itself.

On Windows, the client uses `SO_REUSEADDR` to let its listener and outbound
connection share the same local endpoint. This is required by the discovery
protocol. Bind the client to a specific local IPv4 address rather than a
wildcard address.
