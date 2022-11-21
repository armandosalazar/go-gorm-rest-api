package routes

import "net/http"

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
