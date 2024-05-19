# Config

**juicebox_host**: Set this field to the ip of the JuiceBox device.

**juicebox_device_name**: Set this field to the name of the device given in the Enel X Way app.

**dst**: Ip of the host the juicebox normally connects to

**debug**: Enable debug loggin

# Setup

There is some required setup to get some of the config values and ensure the device connects to the proxy.

1. Find the ip of the JuiceBox. This should be available in the router assuming it provides DHCP. Set the ip as static so that the IP doesn't change. This value is used as `juicebox_host` above.
2. Connect to the juicebox via telnet: `telnet <ip> 2000`, and type `list` to get the connect domain/port in the line marked UDPC.
3. Translate the domain found in step 2 to the corresponding ip (`dig +short <host> @8.8.8.8`) and set the config field `dst` to the `ip:port` value.
4. Setup dns override for the domain found in step 2 so that the domain maps to the system running juicepassproxy (home assistant ip)
