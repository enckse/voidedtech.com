Builtin
===

various builtin functionality available in macOS

## transcode

`avconvert` and `sips` can be used to convert media between different types
(e.g. MOV to MP4 or HEIC to JPEG

## vnc

open vnc connections via `open`

```
open vnc://myvncserver:5900
```

## scheduling

use `pmset` to schedule power settings

e.g. wake up at midnight every night of the week at midnight
```
pmset repeat wake MTWRFSU 00:00:00
```
