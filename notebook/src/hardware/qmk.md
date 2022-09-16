QMK
===

There is a lot of documentation (by QMK itself) about how to use their tooling to flash systems. The problem one runs into is trying to
do this from Linux directly (therefore no QMK Toolbox).

QMK is an extensive project but much of the documentation/tooling is so generic that finding out how to "do things by hand" requires a bit of digging
(looking through QMK Toolbox to extract commands to manually flash a keyboard). It's possible to fully compile the firmware and flash a keyboard
on Linux from the command line.

## keebio/iris/rev6

For example in order to flash a keebio iris (rev6) from (Alpine) Linux:

- Install: `dfu-programmer`, `qmk-cli`, `gcc-avr`, and `avr-libc`
- Expect `qmk setup` to complain about all the missing tools _for other keyboards that don't matter for this ONE keyboard_
- Run `qmk json2c the.json > keyboards/keebio/iris/keymaps/enckse/keymap.c` if the online configurator was used and the json was downloaded
- Run `qmk compile -kb keebio/iris/rev6 -km enckse`

Now, at this point, it's probably easier to run `dfu-programmer` against each half of the keyboard (remember to set it for flashing):

```
sudo dfu-programmer ATmega32U4 erase --force
sudo dfu-programmer ATmega32U4 flash --force enckse.hex
sudo dfu-programmer ATmega32U4 reset
```

Of course different keyboards will require different, specific tools but QMK itself is definitely trying to cover a LOT more than
a single keyboard needs day-to-day

<sub><sup>Updated: 2022-02-25</sup></sub>
