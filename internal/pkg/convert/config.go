package convert

import "time"

var timeout time.Duration
var isGoTexEnabled bool

func SetTimeout(to time.Duration) {
	timeout = to
}
func EnableGoTex() {
	isGoTexEnabled = true
}
func DisableGoTex() {
	isGoTexEnabled = false
}
