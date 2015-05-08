package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type ResponseHandler func(http.ResponseWriter, *http.Request, httprouter.Params) int

var assetsDirs = []string{"templates", "static"}

type TemplateRequestHandler struct {
	templates *template.Template
}

func main() {
	SetupRoutes()
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SetupRoutes() {
	log.Println("SetupRoutes()")
	r := httprouter.New()
	assets := findAssetsPath()
	var templates = template.Must(template.ParseGlob(fmt.Sprint(assets, "/templates/*.tmpl")))
	templates.ParseGlob(fmt.Sprint(assets, "/templates/*.html"))
	root := &TemplateRequestHandler{templates: templates}
	r.POST("/upload", LogRequest(apiUpload))

	r.GET("/ui/:ui_type", root.ServeHTTP)
	r.GET("/", root.ServeHTTP)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(fmt.Sprint(assets, "/static/")))))
	http.Handle("/", r)
	log.Println("SetupRoutes() end")
}

func hasAssets(path string) bool {
	for _, dir := range assetsDirs {
		if _, err := os.Stat(fmt.Sprint(path, "/", dir)); os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func findAssetsPath() string {
	if hasAssets(".") {
		return "."
	}

	if hasAssets("..") {
		return ".."
	}

	gopathPath := os.ExpandEnv("$GOPATH/src/github.com/zhangpeihao/how_old_are_you")
	if hasAssets(gopathPath) {
		return gopathPath
	}
	log.Fatalln("Cannot find assets in any of the default search paths. Please run in the same directory, in a Go workspace.")
	panic("cannot reach")
}

func LogRequest(handler ResponseHandler) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		start := time.Now()
		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}
		log.Printf("Started %s %s for %s", req.Method, req.URL.Path, addr)
		code := handler(w, req, params)
		log.Printf("Completed %v %s %s in %v", code, http.StatusText(code), req.URL.Path, time.Since(start))
	}
}

func (h *TemplateRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uiType := params.ByName("ui_type")
	if r.URL.Path == "/" {
		uiType = "main"
	}
	err := h.templates.ExecuteTemplate(w, uiType+".html", h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func apiUpload(w http.ResponseWriter, r *http.Request, params httprouter.Params) int {
	http.Error(w, "敬请期待", 400)
	return 400
}
