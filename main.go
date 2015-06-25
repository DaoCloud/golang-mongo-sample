package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    MustConnectMongo()
    http.HandleFunc("/", index)
    http.HandleFunc("/new", insert)
    http.HandleFunc("/drop", drop)
    http.HandleFunc("/env", env)

    log.Println("Start listening...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func index(res http.ResponseWriter, req *http.Request) {
    defer func() {
        if e := recover(); e != nil {
            log.Println(res, e)
            res.WriteHeader(http.StatusInternalServerError)
        }
    }()

    log.Println("Index home")

    t, err := template.New("foo").Parse(string(tpl))
    if err != nil {
        log.Println(err)
        res.WriteHeader(500)
        return
    }

    data := make(map[string]interface{})
    data["List"] = List()
    t.Execute(res, data)
}

func insert(res http.ResponseWriter, req *http.Request) {
    person := &Person{}
    person.Name = req.FormValue("name")
    person.Phone = req.FormValue("phone")

    log.Println("Insert new person %v", *person)
    Insert(person)
    http.Redirect(res, req, "/", 302)
    // log.Println(name, phone)
}

func drop(res http.ResponseWriter, req *http.Request) {

    log.Println("drop database")
    Drop()

    http.Redirect(res, req, "/", 302)
}

func env(res http.ResponseWriter, req *http.Request) {
    env := os.Environ()

    fmt.Fprintln(res, "List of Environtment variables : \n")

    for index, value := range env {
        name := strings.Split(value, "=") // split by = sign

        fmt.Fprintf(res, "[%d] %s : %v\n", index, name[0], name[1])
    }
}
