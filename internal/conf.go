package internal

import (
    "context"
    "database/sql"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "Jesus-ggez/rrhh-api/internal/data"
    "Jesus-ggez/rrhh-api/internal/servo"

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


func RunServo() {
    go func() {
        log.Printf("Serve started at port `:%s`", PORT)
        if err := SERVO.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error: %v", err)
        }
    }()


    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit


    log.Println("Closed serve ...")
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    defer cancel()

    if err := SERVO.Shutdown(ctx); err != nil {
        log.Fatalf("Error to off servo: %v", err)
    }
    log.Println("Servo closed correctly")
}
