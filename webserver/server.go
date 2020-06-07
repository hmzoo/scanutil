package webserver

import (

	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./webserver/home.html")
}


func Serve() {
  hub := newHub()
  	go hub.run()
  	http.HandleFunc("/", serveHome)
  	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
  		serveWs(hub, w, r)
  	})
	  log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
