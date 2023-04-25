package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//Запись логов в файл
	//logFile, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer logFile.Close()
	//
	//infoLog := log.New(logFile, "INFO\t", log.Ldate|log.Ltime)
	//errorLog := log.New(logFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
