package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {
	r.
		_, err := w.Write([]byte("Привет из Snippetbox"))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
