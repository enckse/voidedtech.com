Traffic Control (tc)
===

Any useful notes about using `tc`, one can use `bmon` to interactively review the impact of playing with `tc`

# basic traffic shaping

An example, on Alpine Linux running as a router, of using `ifb` with `tc` to try and shape some traffic (in both directions)

```
#!/bin/bash
# the internal "lan" interface LAN devices are "attached" to
LAN="lan0"
# the ifb device to create/use
IFB="ifb0"

# make sure the ifb device exists
ip link add $IFB type ifb >/dev/null 2>&1
ip link set $IFB up

# using ifb, via lan, to handle interface redirection
tc qdisc add dev $LAN handle ffff: ingress
tc filter add dev $LAN parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0

# using testing numbers
tc qdisc add dev $LAN root handle 1: htb default 10
tc class add dev $LAN parent 1: classid 1:1 htb rate 900mbit
tc class add dev $LAN parent 1:1 classid 1:10 htb rate 850mbit
tc class add dev $LAN parent 1:1 classid 1:20 htb rate 500mbit
tc class add dev $LAN parent 1:1 classid 1:30 htb rate 250mbit

# ideally these would be different as upload != download
tc qdisc add dev $IFB root handle 1: htb default 10
tc class add dev $IFB parent 1: classid 1:1 htb rate 900mbit
tc class add dev $IFB parent 1:1 classid 1:10 htb rate 850mbit
tc class add dev $IFB parent 1:1 classid 1:20 htb rate 500mbit
tc class add dev $IFB parent 1:1 classid 1:30 htb rate 250mbit

# filtering by IP/subnets
tc filter add dev $LAN parent 1:0 protocol ip prio 1 u32 match ip src 192.168.1.10/32 flowid 1:20
tc filter add dev $IFB parent 1:0 protocol ip prio 1 u32 match ip src 10.0.0.0/24 flowid 1:30
```
