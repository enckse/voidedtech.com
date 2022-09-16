Image to TOTP Token
===

Given a relatively "clean" QR code image (without the underlying base32 code),
one should be able to extract via decoding, e.g.:

```
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/liyue201/goqr"
)

func main() {
	file := flag.String("image", "qrcode.png", "image to decode")
	flag.Parse()
	path := *file
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}
	for _, qrCode := range qrCodes {
		fmt.Printf("qrCode text: %s\n", qrCode.Payload)
	}
}
```

<sub><sup>Updated: 2021-09-07</sup></sub>
