package main

import "Jesus-ggez/rrhh-api/internal"

func main() {
    internal.InitAppConfig()

    defer internal.POOL.Close()

    internal.RunServo()
}
