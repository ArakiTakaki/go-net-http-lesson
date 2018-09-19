package server

import (
	"fmt"
	"net/http"
	"text/template"
)

// type Server struct {};
type Page struct {
	Title string
	Count int
}

func GetService() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/info", infoHeader)
	// 静的ファイルを記載
	mux.HandleFunc("/static/", statics)
	// リクエストを受付可能な状態にする

	// httpからのリクエストを監視する
	// これがないと永遠と受付状態が続くっぽい

	fmt.Println()
	return mux
}

//静的ファイル
func statics(w http.ResponseWriter, req *http.Request) {
	// URL.Pathを行うことにより、ディレクトリトラバーサルを防止する。
	http.ServeFile(w, req, req.URL.Path[1:])
}

func infoHeader(w http.ResponseWriter, req *http.Request) {
	// handleは基本的に実行される
	page := Page{"Hello", 1}
	pageString := "resouces/index.tmpl.html"
	tmpl, err := template.ParseFiles(pageString)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
	fmt.Println("test")
}
