package main

import (
	"encoding/json"
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
		req, _ := find(db, httpReq.ApiPath)
		if req.ApiPath != "" {
			result := struct {
				Duplicate bool
				Path      string
			}{
				true,
				req.ApiPath,
			}
			data, _ := json.Marshal(result)
			w.Write(data)
			err1 := httpReq.save(db)
			if err1 != nil {
				w.Write([]byte("error!"))
				return
			}
			return
		}
		err1 := httpReq.save(db)
		if err1 != nil {
			w.Write([]byte("error!"))
			return
		}

		hf.handler_path = httpReq.ApiPath
		handler := APIHandlerStr{hf.handler_path}
		handle := fmt.Sprintf("%s", hf.handler_path)
		http.Handle(handle, &handler)
		result := struct {
			Duplicate bool
			Path      string
		}{
			false,
			r.FormValue("api_path"),
		}
		data, _ := json.Marshal(result)
		w.Write(data)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	db := connection()
	defer db.Close()
	a, _ := findAll(db)

	data := struct {
		Status  map[int]string
		ApiList []string
	}{
		HttpStatusCodes,
		a,
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}
