package handler

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func NewIndex() *app.Handler {
	return &app.Handler{
		Author:          "murasame29",
		BackgroundColor: "#fff",
		Icon: app.Icon{
			Default: "/web/icon/logo.png",
		},
		Keywords: []string{
			"murasame29",
			"Golang",
			"Go",
			"Go-app",
			"Go言語",
			"Golang frontend",
			"Go frontend",
			"Docker",
			"名刺",
			"bussiness card",
		},
		LoadingLabel: "Delivering Docker to the port ... ʕ◔ϖ◔ʔ",
		Name:         "Docker Card",
		// TODO:画像を用意する
		Image: "/web/icon/logo.png",
		// CSS適用可能
		Styles: []string{
			"/web/style/reset.css",
			"/web/style/index.css",
			"https://cdn.jsdelivr.net/npm/xterm@4.17.0/css/xterm.css",
		},
		Scripts: []string{
			"https://cdn.jsdelivr.net/npm/xterm@4.17.0/lib/xterm.min.js",
			"https://cdn.jsdelivr.net/npm/xterm-pty@0.9.4/index.js",
		},
		ThemeColor: "#000000",
		Title:      "Docker Card",
	}
}
