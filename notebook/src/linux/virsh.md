Virtualization (virsh)
===

Notes about using `virsh` and related commands

# Installation

Simple installation (using a vnc "display")
```
virt-install -n myvm \
    --memory 16384 \
    --cpu host \
    --cdrom /path/to/rhel8.4.iso \
    --disk size=50 \
    --graphics vnc,port=5901,listen=0.0.0.0,password=myvncpass \
    --network default \
    --vcpus 2 \
    --os-variant rhel8.4
```

or rebuild from a disk image with _all the same_ settings, by changing:

```
virt-install -n myvm \
    ... \
    # size=50 becomes...
    --disk /path/to/existing/disk.qcow2 \
    ... \
    # remove --cdrom /path/to/rhel8.4.iso
    --import \
    ...
```

_Which is helpful if the qemu XML file gets removed/deleted_

once configuration/install is done then `console=ttyS0,115200` can be
added to the kernel parameters and one can attach to the machine console
via `virsh console myvm`

# attaching

## disks

A block device can be directly attached
```
virsh attach-disk --domain myvm /dev/sda1 vdb --config --type disk
```

## usb device

Need the code for the usb (e.g. `lsusb -v`) to place into the following file
```
vim usbdevice.xml
---
<hostdev mode='subsystem' type='usb' managed='yes'>
    <source>
        <vendor id='0x1234'/>
        <product id='0x5678'/>
    </source>
</hostdev>
```

Which can be attached to a machine
```
virsh attach-device myvm --file usbdevice.xml --config --persistent
```

_The `detach-device` command can work to detach the usb device too_


<sub><sup>Updated: 2021-10-21</sup></sub>
