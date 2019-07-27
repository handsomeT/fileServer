package render

import "gopkg.in/russross/blackfriday.v2"

// render markdown
func Markdown(markdown []byte) []byte {
	html := blackfriday.Run(markdown, blackfriday.WithRenderer(
		blackfriday.NewHTMLRenderer(
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags | blackfriday.CompletePage,
			})))
	return html
}