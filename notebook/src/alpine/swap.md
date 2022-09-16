Swap
===

To enable swap in alpine (for manual disk setup):

during install
```
mkswap /dev/<device>
swapon /dev/<device>
```

after install, booted into machine
```
vim /etc/fstab
---
UUID=<diskuuid>   swap    swap    defaults    0 0
```

enable swap
```
rc-update add swap
```

<sub><sup>Updated: 2021-10-21</sup></sub>
