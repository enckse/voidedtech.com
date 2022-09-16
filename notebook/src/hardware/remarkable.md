Remarkable2
===

The remarkable2 table is an eink notebook.

# Using

It's entirely possible to use the remarkable2 without using connect/cloud services

## Connecting

Connect via USB and configure the usb device with the following

```
auto usb0
iface usb0 inet static
    address 10.11.99.3
    netmask 255.255.255.249
```

## Import/Export

To export notebooks or import PDFs:

- Connect
- _Make sure, under storage settings on the remarkable, to enable the web interface_
- Navigate to http://10.11.99.1 and notebooks can be exported
- Drag-and-drop PDFs onto the web interface to add PDFs to the remarkable

## Templates

To make a new notebook template

- Connect
- Prepare to `ssh root@10.11.99.1` using the password configured on the remarkable
- Make sure the image is 1404x1872 and a PNG.
- Place the image in `/usr/share/remarkable/templates/`
- Edit `/usr/share/remarkable/templates/templates.json` and define the template (may want to back this up first)
- Restart the UI: `systemctl restart xochitl`
