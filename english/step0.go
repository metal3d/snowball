package english

import (
	"github.com/metal3d/snowball/snowballword"
)

// Step 0 is to strip off apostrophes and "s".
//
func step0(w *snowballword.SnowballWord) bool {
	suffix, suffixRunes := w.FirstSuffix("'s'", "'s", "'")
	if suffix == "" {
		return false
	}
	w.RemoveLastNRunes(len(suffixRunes))
	return true
}
