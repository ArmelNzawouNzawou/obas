package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/institutions"
)

func Institutions(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", SchoolHandler(app))
	r.Get("/", UniversitysHandler(app))
	return r
}

func SchoolHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allSchools, err := io.GetSchools()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			schools []io.Schools
			name    string
		}
		data := PageData{allSchools, ""}

		files := []string{
			app.Path + "/html/institutions/institutions.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func UniversitysHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allUniversitys, err := io.GetUniversitys()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			univ []io.Universitys
			name string
		}
		data := PageData{allUniversitys, ""}

		files := []string{
			app.Path + "/html/institutions/institutions.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}
