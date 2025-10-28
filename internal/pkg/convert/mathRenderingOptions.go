package convert

type MathRenderingEngine int

const (
	Mathjax MathRenderingEngine = iota
	Mathml
	Webtex
	Katex
	Gladtex
)

var MathRenderingEngineName = map[MathRenderingEngine]string{
	Mathjax: "mathjax",
	Mathml:  "mathml",
	Webtex:  "webtex",
	Katex:   "katex",
	Gladtex: "gladtex",
}

type MathRenderingOptions struct {
	MathRenderingEngine MathRenderingEngine `json:"MathRenderingEngine"`
	MathRenderingURL    string              `json:"MathRenderingUrl"`
}
