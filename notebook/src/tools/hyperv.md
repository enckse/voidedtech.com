Hyper-V 
===

## Static Guest IP

Enabling static IPs in guests

### Host

In order to get a static guest IP, one can **NOT** use the default switch and instead must create "their own"

Make the switch
```
New-VMSwitch -SwitchName "MySwitchName" -SwitchType Internal
```

Get the necessary index
```
Get-NetAdapter
```

Set the IP settings
```
New-NetIPAddress -IPAddress 192.168.3.1 -PrefixLength 24 -InterfaceIndex {INDEX}
```

Setup NAT (internet access)
```
New-NetNat -Name MyNAT -InternalIPInterfaceAddressPrefix 192.168.3.0/24
```

### Guest

Setup the guest IP (manually, no DHCP server exists)
```
vim /etc/network/interfaces
---
auto eth0
iface eth0 inet static
	address 192.168.3.2 
	netmask 255.255.255.0
	gateway 192.168.3.1
```

_don't forget to also configure a DNS resolver_
