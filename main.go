package main

import "Jesus-ggez/rrhh-api/internal"

func main() {
    app := internal.App

    app.InitAppConfig()
    defer app.Pool.Close()

    app.Run()
}
