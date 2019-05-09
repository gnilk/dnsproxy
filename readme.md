A DNS Proxy with parental blocking abilities.
Devices/IP (or names) can be blocked on high-level or on a per domain basis.

Can talk to Netgear routers and Unifi Controllers to obtain device names which can be used by the block rules.
A positive block can be redirected to specific IP number (defaults to 0.0.0.0).

New from 2019-05-09
* Unifi Controller support (reading device names)
* Added device name to log instead of client-address (client IP is still used if device name is not found)

see: "config.json" for an example
