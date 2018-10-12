package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Render(response http.ResponseWriter, request *http.Request, view string, object map[string]interface{}) {
	// device := request.UserAgent()
	// expression := regexp.MustCompile("(Mobi(le|/xyz)|Tablet)")
	// if !expression.MatchString(device) {
	// 	response.Header().Set("Connection", "keep-alive")
	// }
	// response.Header().Set("Vary", "Accept-Encoding")
	// response.Header().Set("Cache-Control", "private, max-age=7776000")
	// response.Header().Set("Transfer-Encoding", "gzip, chunked")
	// var goHTML *template.Template
	templateFilePath := filepath.Join("./app/views/", view+".gohtml")
	fmt.Println(templateFilePath)
	info, err := os.Stat(templateFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(response, request)
			return
		}
	}
	if info.IsDir() {
		http.NotFound(response, request)
		return
	}
	goHTML := template.Must(template.ParseFiles("./app/views/" + view + ".gohtml"))
	err = goHTML.Execute(response, object)
	if err != nil {
		log.Println(err.Error())
		http.Error(response, http.StatusText(500), 500)
	}

}
