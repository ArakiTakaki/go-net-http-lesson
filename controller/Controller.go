package controller

import (
	"fmt"
	"text/template"
)

type Attributes interface{}

func View(file string, data Attributes) {
	directory := "resouces" + file + ".tmpl.html"
	tmpl, err := template.ParseFiles(directory)
	if err != nil {
		fmt.Print(file)
		fmt.Println("でエラーが発生しました")
		panic(err)
	}
	//err = tmpl.Execute(w, data)
}
