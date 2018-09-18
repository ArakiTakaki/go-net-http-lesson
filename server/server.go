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

func Start(){
	// リクエストを受付可能な状態にする
	http.HandleFunc("/info",infoHeader)
	// httpからのリクエストを監視する
	// これがないと永遠と受付状態が続くっぽい

	//静的ファイルを記載
	http.HandleFunc("/static/",statics)
	http.ListenAndServe(":3000", nil)
}

//静的ファイル
func statics(w http.ResponseWriter, req *http.Request){
	// URL.Pathを行うことにより、ディレクトリトラバーサルを防止する。
	http.ServeFile(w, req, req.URL.Path[1:])
}


func infoHeader(w http.ResponseWriter, req *http.Request){
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
	return
}
