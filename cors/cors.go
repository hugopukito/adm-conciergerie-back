package cors

import (
	"net/http"
)

func EnableCors(w *http.ResponseWriter, r *http.Request) string {
	header := (*w).Header()

	allowList := map[string]bool{
		"http://localhost:5500":             true,
		"https://adm-conciergerie.com/back": true,
		"http://151.80.155.148/back":        true,
		"http://adm-conciergerie.com/back": true,
	}

	if origin := r.Header.Get("Origin"); allowList[origin] {
		header.Add("Access-Control-Allow-Origin", origin)
	}

	header.Add("Access-Control-Allow-Headers", "Authorization, Content-Type")
	header.Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		(*w).Header().Add("Access-Control-Max-Age", "3600")
		(*w).WriteHeader(http.StatusOK)
		return "options"
	}
	return ""
}
