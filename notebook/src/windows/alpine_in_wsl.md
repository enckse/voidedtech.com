Alpine in WSL
===

Configure Alpine in WSL (assumes WSL has been configured/setup)

# Create WSL Image

Import Alpine into WSL (make sure the directory structure exists)
```
wsl --import AlpineX.YY C:\Path\To\Directory\For\Alpine\WSL .\alpine.version.tar.gz
```

Test getting into it
```
wsl -d AlpineX.YY
```

## Configure the Image

May want to make sure any Terminals call as a user, but make sure to `adduser` that user first
```
wsl -d AlpineX.YY --user myuser
```

May want to install `alpine-conf` to run things like timezone setup (e.g. `setup-timezone`)

## Troubleshooting

Mounts like `C:\` are in `/mnt/c` but also `\\wsl$` can access WSL from the Windows host

<sub><sup>Updated: 2022-06-25</sup></sub>
