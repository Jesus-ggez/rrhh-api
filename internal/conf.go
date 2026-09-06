package internal

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    "Jesus-ggez/rrhh-api/internal/servo"
    "Jesus-ggez/rrhh-api/internal/data"

    "github.com/joho/godotenv"
)

var (
    DATABASE_AUTH_TOKEN string
    DATABASE_URL        string
    PORT                string
    POOL                *sql.DB
    SERVO               *http.Server
)

func InitAppConfig() {
    if err := godotenv.Load(); err != nil {
        log.Print("Error loading dotenv file: " + err.Error())
    }

    DATABASE_AUTH_TOKEN = os.Getenv("DATABASE_AUTH_TOKEN")
    if DATABASE_AUTH_TOKEN == "" { log.Panic("environ DATABASE_AUTH_TOKEN not loaded") }

    DATABASE_URL = os.Getenv("DATABASE_URL")
    if DATABASE_URL == "" { log.Panic("environ DATABASE_URL not loaded") }

    PORT = os.Getenv("PORT")
    if PORT == "" { PORT = "3000" }

    SERVO = servo.NewServer(PORT)

    var err error
    POOL, err = data.NewPool(DATABASE_AUTH_TOKEN, DATABASE_URL)
    if err != nil { log.Panic("Error building main connection pool") }
}
