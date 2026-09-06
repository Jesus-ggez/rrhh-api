package internal

import (
    "context"
    "crypto/tls"
    "database/sql"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "time"
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

