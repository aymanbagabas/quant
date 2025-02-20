//go:build go1.18
// +build go1.18

package median

import "github.com/aymanbagabas/quant/internal"

func sortSlice(s chValues) {
	internal.SortFunc(s, s.Compare)
}
