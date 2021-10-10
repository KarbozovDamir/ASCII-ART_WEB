package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	server "web/service"
)

const port = ":1989"

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)

	http.HandleFunc("/", server.MainPageHandler)
	http.HandleFunc("/ascii-art", server.PostPageHandler)
	fmt.Println("Listening on the port " + port + "\nhttp://localhost" + port + "/")
	log.Println("Listening on the port " + port)
	log.Println(http.ListenAndServe(port, nil))
}
