package convert

import "time"

var timeout time.Duration
var isGoTexEnabled bool
var mathRenderingConfig MathRenderingOptions

func SetTimeout(to time.Duration) {
	timeout = to
}
func EnableGoTex() {
	isGoTexEnabled = true
}
func DisableGoTex() {
	isGoTexEnabled = false
}
func SetMathRenderingOptions(engine MathRenderingEngine, url string) {
	mathRenderingConfig.MathRenderingEngine = engine
	mathRenderingConfig.MathRenderingURL = url
}
