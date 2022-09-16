Disks
===

# iso handling

## alpine (rpi4) bootable disks

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

## cloud ready

To create a cloud-init ready iso on macOS, place "user-data" and "meta-data" in a `configs/` directory (or any name)

```
hdiutil makehybrid -o init.iso -joliet -iso -default-volume-name cidata configs/
```

(make sure to specify `-joliet -iso` because otherwise macOS will try to use `-hfs` which many systems will not have installed/ready)

## mount/unmount

an iso can attached or detached via `hdiutil`
```
hdiutil attach <file.iso> -mountpoint /Volumes/mymount
# and then
hdiutil detach /Volumes/mymount
```
