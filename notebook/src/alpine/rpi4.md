Installation (Raspberry Pi 4)
===

Assuming the boot media is created (sd card) then the following should,
generally, get alpine up and running

1. Partition disks (as preferred) - leaving the boot partition alone
2. Use `ext4` BUT one will have to manually edit `setup-disk` to allow the fs for the Pi
3. Make sure to `mount` the root partition to `/mnt`
4. And then make sure to `mount -o remount,rw /media/<boot device>` and then `mount --bind /media/<boot device> /mnt/boot`
5. Finally run `setup-disk -m sys /mnt`

_This should get everything installed, after first boot one should edit fstab and
make /boot/boot a bind mount to boot itself (FAT32 doesn't allow symlinking in this case)_

<sub><sup>Updated: 2021-10-21</sup></sub>
