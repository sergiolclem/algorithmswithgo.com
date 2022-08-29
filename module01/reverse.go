package module01

import (
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	wl := strings.Split(word, "")
	for i := range wl[0:int(len(wl)/2)] {
		wl[i], wl[len(wl)-1-i] = wl[len(wl)-1-i], wl[i]
	}

	return strings.Join(wl, "")
}
