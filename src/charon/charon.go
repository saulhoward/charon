// A simple go server.

package main

import (
    //"fmt"
    "os"
    "path/filepath"
    "net/http"
    //"html/template"
)

const (
    lenPhotosPath int = len("/photos/")
    photosDir string = "/srv/saulandkannika/photos/"
)

type photo struct {
    Title string
    Body string
}

func main() {
    http.HandleFunc("/photos/", photoHandler)
    //http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}

func photoHandler(w http.ResponseWriter, r *http.Request) {

    // parse the path, eg, /photos/dir/file
    // The requested photo starts after the /photos/
    reqPhoto := r.URL.Path[lenPhotosPath:]
    osPath = photosDir + reqPhoto

    //http.Error(w, base, http.StatusInternalServerError)

    // is this asking for an actual file (is there an extension)
    if len(filepath.Ext(reqPhoto)) > 0 {
        f, err := os.Open(osPath)
        defer f.Close()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.ServeFile(w, r, osPath)
        return
    }

    // otherwise fetch the photo metadata
    // and the page metadata (user, key, etc)
    //p, _ := getPhoto(name)
    //base := filepath.Base(fullName)
    //dir := filepath.Dir(fullName)

    // and serve the template

    data := map[string]string{}
    data["imgSrc"] = "balls.jpg"
    tmpl, err := template.New("index.html").ParseFiles("/srv/charon/templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

/*
func getPhoto(n string) (p photo, err error) {
    p = photo{
        Title: "hey",
        Body: "ho",
    }
    return
}
*/


