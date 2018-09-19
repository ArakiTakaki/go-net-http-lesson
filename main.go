package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ArakiTakaki/go-net-http-lesson/lib"

	"github.com/ArakiTakaki/go-net-http-lesson/get"
	"github.com/ArakiTakaki/go-net-http-lesson/routes"
)

func main() {
	get := get.HttpUtilGet("http://araki.mars.jp/")

	res := get.ResponseObj()
	green("# Response")
	fmt.Println("ResponseDetail")

	next()

	green("## ACCESS STATUS")
	fmt.Print(" * ")
	fmt.Println(res.Request.URL)
	fmt.Print(" * ")
	fmt.Println(get.Status())
	fmt.Print(" * ")
	fmt.Println(res.Proto)
	fmt.Print(" * ")
	fmt.Println(res.Request.Method)

	next()

	green("## RESPONSE HEADER")
	data := get.ToArrayHeader()
	for i := range data {
		fmt.Println(data[i])
	}

	next()

	green("## RESPONSE CONTENT ")
	// bodyのリソースを解放するため
	defer res.Body.Close()
	// bodyの char[] 型を取得する
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("```")
	// char[] => string に変更し出力を行う。
	fmt.Println(string(body))
	fmt.Println("```")

	router := lib.NewRoute()
	routes.Routing(router)
	router.Start(":3000")

}

func green(data string) {
	fmt.Printf("\x1b[32m%s\x1b[0m", data)
	fmt.Print("\n\n")
}
func next() {
	fmt.Print("\n\n")
}
