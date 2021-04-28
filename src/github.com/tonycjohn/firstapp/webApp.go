package main

import (
	"net/http"
)

func webApp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
		//var p []byte
		host := r.Host
		method := r.Method
		//r.Context().Value()
		w.Write([]byte(host))
		w.Write([]byte(method))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err.Error())
	}

}
