package httpservo

import "net/http"

func ServeVideo(w http.ResponseWriter, r *http.Request, filePath string) {
	http.ServeFile(w, r, filePath)
}
