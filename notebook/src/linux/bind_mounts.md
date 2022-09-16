Bind Mounts
===

A very useful tool in the Linux toolbox is a "bind mount" which can be used, for
example, to mount a directory from one location to another (e.g. no symlinks, allows
for more options like _read only_).

## Example

For example to bind mount a "new /var" over the old:
```
mount --bind /mnt/data/var /var
```

_These can also be placed into `/etc/fstab`_
