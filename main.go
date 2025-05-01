package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
    "strings"

    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
)

var tmpl *template.Template
var renderer *html.Renderer

type PostLink struct {
    Title string
    Path  string
}

func main() {
    // Load templates
    tmpl = template.Must(template.ParseFiles("templates/base.html"))

    // Set up markdown renderer
    renderer = html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})

    // Routes
    http.HandleFunc("/", router)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Serving at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func router(w http.ResponseWriter, r *http.Request) {
    path := strings.Trim(r.URL.Path, "/")

    if path == "" {
        handleIndex(w)
    } else {
        handlePost(w, path)
    }
}

func handleIndex(w http.ResponseWriter) {
    posts := listPosts()

    err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
        "Title":   "Morss Blog",
        "Posts":   posts,
        "Content": template.HTML("<p>Select a post from the list.</p>"),
    })
    if err != nil {
        log.Printf("index render error: %v", err)
    }
}

func handlePost(w http.ResponseWriter, name string) {
    file := fmt.Sprintf("content/%s.md", name)

    data, err := ioutil.ReadFile(file)
    if err != nil {
        http.NotFound(w, nil)
        return
    }

    htmlContent := markdown.ToHTML(data, nil, renderer)
    posts := listPosts()

    err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
        "Title":   name,
        "Posts":   posts,
        "Content": template.HTML(htmlContent),
    })
    if err != nil {
        log.Printf("post render error: %v", err)
    }
}

func listPosts() []PostLink {
    files, _ := filepath.Glob("content/*.md")
    var posts []PostLink

    for _, file := range files {
        name := strings.TrimSuffix(filepath.Base(file), ".md")
        posts = append(posts, PostLink{Title: name, Path: "/" + name})
    }
    return posts
}

