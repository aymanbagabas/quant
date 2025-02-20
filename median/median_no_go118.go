//go:build !go1.18
// +build !go1.18

package median

import "sort"

func sortSlice(s chValues) {
	sort.Sort(s)
}
