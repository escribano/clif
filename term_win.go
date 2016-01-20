// +build windows

package clif

import (
	"github.com/escribano/ts"
)

func init() {
	TermWidthCall = func() (int, error) {
		size, err := ts.GetSize()
		if err != nil {
			return TERM_DEFAULT_WIDTH, err
		}
		return size.Col(), nil
	}

	TermWidthCurrent, _ = TermWidthCall()
}
