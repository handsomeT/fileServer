package writer

import (
	"fileServer/lib/util"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
)

func File(w http.ResponseWriter, name string, data []byte) {
	// 依据文件后缀打开文件
	switch util.GetFileFormat(name) {
	case ".md":
		write(w, markdown(data))
	default:
		write(w, plain(data))
	}
}

func plain(data []byte) []byte {
	return data
}

// writer markdown
func markdown(markdown []byte) []byte {
	html := blackfriday.Run(markdown, blackfriday.WithRenderer(
		blackfriday.NewHTMLRenderer(
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags | blackfriday.CompletePage,
			})))
	return html
}
