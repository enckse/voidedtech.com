Boot (EFI)
===

Using EFI in Alpine is supported but the documentation is not consistently available
in the wiki where one would expect.

## environment

Setup the necessary environment variables

```
export BOOTLOADER=grub
export USE_EFI=1
```

## packages

Drop `syslinux` and add grub/friends

```
apk del syslinux
apk add grub grub-efi efibootmgr
```
