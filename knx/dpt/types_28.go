// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_28001 represents DPT 28.001 / Var String UTF-8.
type DPT_28001 string

func (d DPT_28001) Pack() []byte {
	// len(d) is gives us the number of bytes in d
	var buf = make([]byte, 1, len(d)+2)

	buf = append(buf, d...)
	buf = append(buf, 0x00)

	return buf
}

func (d *DPT_28001) Unpack(data []byte) error {
	var buf = data[1 : len(data)-1]

	*d = DPT_28001(buf)

	return nil
}

func (d DPT_28001) Unit() string {
	return ""
}

func (d DPT_28001) String() string {
	return string(d)
}

func (d DPT_28001) AsFloat32() float32 {
	return float32(len(d))
}

func (d *DPT_28001) SetFloat32(val float32) {
	*d = DPT_28001(fmt.Sprintf("%v", val))
}
