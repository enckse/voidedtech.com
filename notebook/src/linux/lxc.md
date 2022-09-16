LXC
===

## nesting

To enable nesting (containers in an LXC container) one wants to include
the `/usr/share/lxc/config/nesting.conf` in a machine's configuration. Though if,
for example, something like an "apparmor" setting is in the config and apparmor
is not installed it may require taking the values from the nesting config and
placing them directly into the machine's config (mainly want to make sure things
like `/proc` or `/sys` or `/dev` are setup properly as more than read-only)

_This does have security implications_

## filesystems

lxc can do more than `directory`-based containers (including image files via
loop, btrfs, lvm, etc.) during creation.

<sub><sup>Updated: 2021-09-08</sup></sub>
