package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// Page Service
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := filepath.Join("data", p.Title+".txt")
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join("data", title+".txt")
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Caching
var templates *template.Template

// Menambahkan fungsi konversi []byte ke string dalam template
func init() {
	// Menambahkan fungsi ke dalam template.FuncMap terlebih dahulu
	templates = template.New("").Funcs(template.FuncMap{
		"bytesToString": func(b []byte) string {
			return string(b) // Konversi []byte ke string
		},
	})

	// Parsing template setelah fungsi ditambahkan
	templates = template.Must(templates.ParseGlob("tmpl/*.html"))
}

// Handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	p.Body = convertLinks(p.Body)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Inter-page linking
var linkPattern = regexp.MustCompile(`\[(.+?)\]`)

func convertLinks(text []byte) []byte {
	return linkPattern.ReplaceAllFunc(text, func(match []byte) []byte {
		pageName := string(match[1 : len(match)-1])
		pageURL := "/view/" + regexp.MustCompile(`\s+`).ReplaceAllString(pageName, "%20")
		return []byte(`<a href="` + pageURL + `">` + template.HTMLEscapeString(pageName) + `</a>`)
	})
}

// Validation
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9_-]+)$")

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
