package internal

import (
    "database/sql"
    "net/http"
)


var (
    DATABASE_AUTH_TOKEN string
    DATABASE_URL        string
    PORT                string
    POOL                *sql.DB
    SERVO               *http.Server
)

func InitAppConfig() {
}

func newServer() {
    sv := &http.Server{
        Addr: ":" + PORT,
    }

    sv.Addr
}
