raspberry pi (4)
===

## bootable disk

partition the disk
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
