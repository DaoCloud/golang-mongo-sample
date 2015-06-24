package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    MustConnectMongo()
    http.HandleFunc("/", hello)
    http.HandleFunc("/env", env)

    fmt.Println("Start listening...")
    go Log()
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        panic(err)
    }
}

func Log() {
    for {
        time.Sleep(30 * time.Second)
        fmt.Println("Hello World !")
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    defer func() {
        if e := recover(); e != nil {
            fmt.Fprintln(res, e)
        }
    }()
    fmt.Fprintln(res, "Hello World, ", GetResult())
}

func env(res http.ResponseWriter, req *http.Request) {
    env := os.Environ()

    fmt.Fprintln(res, "List of Environtment variables : \n")

    for index, value := range env {
        name := strings.Split(value, "=") // split by = sign

        fmt.Fprintf(res, "[%d] %s : %v\n", index, name[0], name[1])
    }
}
