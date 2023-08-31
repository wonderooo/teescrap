package webapp

import "net/http"

func HandleStaticContent() http.Handler {
	return http.FileServer(http.Dir("./static"))
}