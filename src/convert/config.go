package convert

import "time"

var timeout time.Duration

func SetTimeout(to time.Duration) {
	timeout = to
}
