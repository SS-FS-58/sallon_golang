package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("salon"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("salon", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(UserName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": UserName,
	}
	if encoded, err := cookieHandler.Encode("salon", value); err == nil {
		cookie := &http.Cookie{
			Name:  "salon",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getVatNumberSession(r *http.Request) (VatNumber string) {
	if cookie, err := r.Cookie("vatnumber"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("vatnumber", cookie.Value, &cookieValue); err == nil {
			VatNumber = cookieValue["name"]
		}
	}
	return VatNumber
}

func setVatNumberSession(VatNumber string, w http.ResponseWriter) {
	value := map[string]string{
		"name": VatNumber,
	}
	if encoded, err := cookieHandler.Encode("vatnumber", value); err == nil {
		cookie := &http.Cookie{
			Name:  "vatnumber",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "salon",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func clearVatnumberSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "vatnumber",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
