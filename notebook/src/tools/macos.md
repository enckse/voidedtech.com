macOS
===

## builtin

### vnc

macOS can open vnc connections via `open`

```
open vnc://myvncserver:5900
```

### quarantine

to remove quarantine on a downloaded (trusted) file:
```
xattr -d com.apple.quarantine <file>
```

### keychain

retrieve a password (plaintext) from the keychain
```
security find-generic-password -a <item> -g -w
```

## disks

### iso handling

#### mounting

mounting linux ISO files can be accomplished through a few extra steps
```
hdiutil attach -nomount <image.iso>
```

which will output disk information about the attachment it should make
```
# need somewhere to mount
mkdir -p iso
mount -t cd9660 /dev/disk(see above) ./iso
```

do not forget to umount/detach
```
umount ./iso
hdiutil detach /dev/disk(see above)
```

#### alpine (rpi4) bootable disks

Partition the target
```
diskutil partitionDisk <target disk> MBR "FAT32" ALP 2GB "Free Space" SYS R
```

and then make sure to set it for booting
```
sudo fdisk -e <target disk>
> f 1
> w
> exit
```

Unpack the tar payload in the `/Volumes/ALP` directory, and
```
vim usercfg.txt
---
enable_uart=1
gpu_mem=32
disable_overscan=1
```

#### cloud ready

To create a cloud-init ready iso on macOS, place "user-data" and "meta-data" in a `configs/` directory (or any name)

```
hdiutil makehybrid -o init.iso -joliet -iso -default-volume-name cidata configs/
```

(make sure to specify `-joliet -iso` because otherwise macOS will try to use `-hfs` which many systems will not have installed/ready)

#### mount/unmount

an iso can attached or detached via `hdiutil`
```
hdiutil attach <file.iso> -mountpoint /Volumes/mymount
# and then
hdiutil detach /Volumes/mymount
```

### virtualization

#### dhcp

the leases are stored here
```
cat /var/db/dhcpd_leases
```

_this file can be deleted to reset the leases_
