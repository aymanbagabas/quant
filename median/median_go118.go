//go:build go1.18
// +build go1.18

package median

import "github.com/soniakeys/quant/internal"

func sortSlice(s chValues) {
	internal.SortFunc(s, s.Compare)
}
