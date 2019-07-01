A DNS Proxy with parental blocking abilities.
Devices/IP (or names) can be blocked on a per device level or on a per domain basis.

Can talk to Netgear routers and Unifi Controllers to obtain device names which can be used by the block rules.
A positive block can be redirected to specific IP number (defaults to 0.0.0.0).

## New from 2019-07-01
* All rules can now have TimeSpan definitions
* Possible to test configuration before deploying it (./dnsproxy -t)
* Timeout to router configuration added (default is 10 seconds if not specified)

The new timespan handling makes it possible to specify an allowed interval where the default rule is blocking.
Making it easier to handle screentime rules for kids.
```
	{
		"Name":"Kids-iPad",
		"Rules": [
			{
				"Type": "ActionTypePass",
                            	"TimeSpan": "16:30-17:30"

			},
			{
				"Type":"ActionTypeBlockedDevice"
			}
		]
	},
```

## New from 2019-05-09
* Unifi Controller support (reading device names)
* Added device name to log instead of client-address (client IP is still used if device name is not found)

see: "config.json" for an example
