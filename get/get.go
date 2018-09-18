package get

import (
	"log"
	"net/http"
	"sort"
	"strings"
)

type Get struct{ response *http.Response }

func HttpUtilGet(url string) *Get {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return &Get{response}
}

func (g *Get) Status() int {
	return g.response.StatusCode
}

func (g *Get) ResponseObj() *http.Response {
	return g.response
}

func (g *Get) ToArrayHeader() []string {
	var work []string
	for key, values := range g.response.Header {
		tmp := " * " + key + " : "
		tmp += strings.Join(values, ", ")
		work = append(work, tmp)
	}
	sort.Strings(work)
	return work
}
