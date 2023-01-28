Cloud Init
===

Notes about using cloud-init to bootstrap a small system (mostly virtualized in
the following use cases)

## networking

It can be easier to set a static IP via kernel parameters (e.g. `ip=` then it
is to find the documentation for cloud-init networking which will tell you to
use `ip=` or some other arcane method)

### post-boot

a cloud-init system can end up no longer using `ip=` in some cases (e.g.
Fedora cloud images) which means creating an `/etc/sysconfig/network-scripts/ifcfg-Wired_connection_1`
may be required to handle networking with contents like (for Fedora)

```
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
IPADDR=192.168.64.2
PREFIX=24
GATEWAY=192.168.64.1
DNS1=192.168.1.1
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
NAME="Wired connection 1"
DEVICE=enp0s1
ONBOOT=yes
AUTOCONNECT_PRIORITY=-999
```

## vmlinuz/initramfs

These can be pulled out of most ISOs though the ISO needs to have been built
with the `virtio` module if it is going to be used for virtualization (e.g.
passing a kernel/initram to a process to start as a VM)

## disks

It is easier to use a "cloud ready" image that is a disk image that can be
booted directly (and not booting an installation media/iso). Then one can set
the kernel parameter `root=/dev/path2` as necessary to boot the rootfs object properly.
