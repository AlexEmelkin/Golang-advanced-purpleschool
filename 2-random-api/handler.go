package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandHandler struct{}

func NewRandHandler(router *http.ServeMux) {
	handler := &RandHandler{}
	router.HandleFunc("/rand", handler.Rand())
}

func (handler *RandHandler) Rand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rn := strconv.Itoa(rand.Intn(5) + 1)
		//fmt.Println("test", rn)
		w.Write([]byte(rn))
	}
}
