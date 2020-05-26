package main

import (
	"encoding/json"
	"fmt"
	"goTestForm/models"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {

	blogName := r.PostFormValue("blogName")
	// fmt.Fprintf(w, blogName)

	blog := models.Blog{}
	blog.Name = blogName
	err := blog.Add()
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/list", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var blog models.Blog
	json.Unmarshal(body, &blog)
	rows := blog.Delete()
	if rows < 1 {
		ResponseCode(w, 201)
		return
	}
	ResponseCode(w, 200)
}

type ResCode struct {
	Code int `json:"code"`
}

func ResponseCode(w http.ResponseWriter, code int) {
	res := ResCode{
		Code: code,
	}
	output, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w, string(output))
}
