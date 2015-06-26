package main

import (
    "fmt"
    "log"
    "os"

    "gopkg.in/mgo.v2"
)

var DB *mgo.Database

func MustConnectMongo() {
    if err := ConnectMongo(); err != nil {
        panic(err)
    }
}

func InitDB() {
    defer func() {
        if e := recover(); e != nil {
            log.Println(e)
        }
    }()
    Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321"})
    Insert(&Person{Name: "Cla", Phone: "+66 33 1234 5678"})
}

var (
    username string
    password string
    host     string
    port     string
    instance string
)

func Config() {

    username = os.Getenv("MONGODB_USERNAME")
    password = os.Getenv("MONGODB_PASSWORD")
    host = os.Getenv("MONGODB_PORT_27017_TCP_ADDR")

    if len(host) == 0 {
        host = "localhost"
    }

    port = os.Getenv("MONGODB_PORT_27017_TCP_PORT")
    if len(port) == 0 {
        port = "27017"
    }

    instance = os.Getenv("MONGODB_INSTANCE_NAME")
}

func ConnectMongo() error {
    conn := ""
    if len(username) > 0 {
        conn += username

        if len(password) > 0 {
            conn += ":" + password
        }

        conn += "@"
    }

    conn += fmt.Sprintf("%s:%s/%s", host, port, instance)

    session, err := mgo.Dial(conn)
    if err != nil {
        return err
    }

    DB = session.DB(instance)
    peopleC = DB.C("people")

    return nil
}

func Drop() {
    // Drop Collection
    err := peopleC.DropCollection()
    if err != nil {
        log.Println(err)
    }
}
