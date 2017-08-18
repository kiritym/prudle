package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var httpReq HttpReq

type APIHandlerStr struct {
	path string
}

func (hf *APIHandlerStr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := connection()
	defer db.Close()
	req, _ := find(db, r.URL.Path)
	w.Header().Set("Content-Type", req.ContentType+";charset="+req.CharSet)
	w.WriteHeader(strToInt(req.StatusCode))
	w.Write([]byte(req.Payload))
}

type HandlerFactory struct {
	handler_path string
}

func (hf *HandlerFactory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		httpReq = HttpReq{ApiPath: r.FormValue("api_path"),
			ContentType: r.FormValue("content_type"),
			CharSet:     r.FormValue("char_set"),
			StatusCode:  r.FormValue("status_code"),
			Payload:     r.FormValue("payload"),
		}
		db := connection()
		defer db.Close()
		// req, _ := find(db, httpReq.ApiPath)
		// if req.ApiPath != "" {
		// 	w.Write([]byte("duplicate"))
		// 	return
		// }
		err1 := httpReq.save(db)
		if err1 != nil {
			w.Write([]byte("error!"))
			return
		}

		hf.handler_path = httpReq.ApiPath
		handler := APIHandlerStr{hf.handler_path}
		handle := fmt.Sprintf("%s", hf.handler_path)
		http.Handle(handle, &handler)
		w.Write([]byte(r.FormValue("api_path")))
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, HttpStatus)
}
