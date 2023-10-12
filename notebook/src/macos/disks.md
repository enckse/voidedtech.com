Disks
===

tools to work with disks within macOS

## iso handling

### mounting

mounting an iso in macOS can be easy or hard depending on the iso file

#### simple

some iso files are easy to mount
```
hdiutil attach <file.iso> -mountpoint /Volumes/mymount
# and then
hdiutil detach /Volumes/mymount
```

#### advanced

and some require a few extra steps
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

### creation

creating different iso files also can require some work

#### cloud ready

To create a cloud-init ready iso on macOS, place "user-data" and "meta-data" in a `configs/` directory (or any name)

```
hdiutil makehybrid -o init.iso -joliet -iso -default-volume-name cidata configs/
```

(make sure to specify `-joliet -iso` because otherwise macOS will try to use `-hfs` which many systems will not have installed/ready)

