package spanish

import (
	"github.com/metal3d/snowball/snowballword"
)

// Applies transformations necessary after
// a word has been completely processed.
//
func postprocess(word *snowballword.SnowballWord) {

	removeAccuteAccents(word)
}
