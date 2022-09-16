Kinesis Advantage 2
===

Using a kinesis advantage 2 (on Linux) is not any different than any other operating system.
A few useful notes, though, about working with it.

## Export

To "export" the current configuration:

1. Go into Power User Mode (progm+shift+esc)
2. Enable vdrive (progm+F1)
3. `mount -t vfat /dev/sda /mnt/vdrive` (or whichever device gets mapped, notice the _lack_ of the partition)
4. Poke around in the mounted area, mainly in the `active/` directory
5. Make sure to `umount /mnt/vdrive`
6. Close Power User Mode (same key press as 1)
