package main

import (
	"fmt"
	"net/http"
	"snippetbox/pkg/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	app.serveError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	app.serveError(w, err)
	// }

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serveError(w, err)
		return
	}
	fmt.Fprintf(w, "%v", s)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "0 snail\n Climb Mount Fuji,\nBut slowly, slowly!\n\n - Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serveError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
