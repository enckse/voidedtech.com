DNF
===

## chroots

when building chroot environments in a dnf-based distribution (e.g. RHEL, Alma)
to be able to utilize some `dnf` options to speed-up building/configuring
chroots

For example to create a new chroot, using the host's release, use a "shared" cache, and keep a cache

```
dnf install -y --releasever=/ --installroot /path/to/chroot
--setopt=cachedir=/var/cache/dnf --setopt=keepcache=True package1 package2
```
