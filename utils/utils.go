package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
)

var tmpl *template.Template

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//SessionStore ..
var SessionStore = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

//LoadTemplates ..
func LoadTemplates(pattern string) {
	//tmpl = template.Must(template.ParseGlob(pattern))
}

//AuthRequired ..
func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := SessionStore.Get(r, "session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

//ClearCookie ..
func ClearCookie(response http.ResponseWriter) {
	Usercookie := &http.Cookie{
		Name:   "User",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, Usercookie)
	Authcookie := &http.Cookie{
		Name:   "Auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, Authcookie)
}

//ExecuteTemplate ..
func ExecuteTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, tmplName, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

//GetCookieUser ..
func GetCookieUser(r *http.Request) (User CmnModel.LoginModel, Auth map[string]bool) {

	if cookie, err := r.Cookie("User"); err == nil {
		cookieValue := CmnModel.LoginModel{}
		if err = cookieHandler.Decode("User", cookie.Value, &cookieValue); err == nil {
			User = cookieValue

		}
	}
	if cookie, err := r.Cookie("Auth"); err == nil {
		cookieValue := make(map[string]map[string]bool)
		if err = cookieHandler.Decode("Auth", cookie.Value, &cookieValue); err == nil {
			Auth = cookieValue["Auth"]
		}
	}
	return User, Auth
}


//SetCookieHandler ..
func SetCookieHandler(User CmnModel.LoginModel, Auth map[string]bool, r *http.Request, w http.ResponseWriter) {

	Authvalue := map[string]map[string]bool{
		"Auth": Auth,
	}
	encoded, err := cookieHandler.Encode("User", User)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "User",
			Value: encoded,
			//	Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	if encoded, err := cookieHandler.Encode("Auth", Authvalue); err == nil {
		cookie := &http.Cookie{
			Name:  "Auth",
			Value: encoded,
			//Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

//RespondwithJSON ..
func RespondwithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	data, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
