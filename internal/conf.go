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

type Application struct {
    Servo               *http.Server
    Pool                *sql.DB
    Config              *AppConfig
}

type AppConfig struct {
    DATABASE_AUTH_TOKEN string
    DATABASE_URL        string
    PORT                string
    CERT                string
    KEY                 string
}
var App = Application {}

func (a *Application) InitAppConfig() {
    if err := godotenv.Load(); err != nil {
        log.Print("Error loading dotenv file: " + err.Error())
    }

    conf := &AppConfig{}

    conf.DATABASE_AUTH_TOKEN = os.Getenv("DATABASE_AUTH_TOKEN")
    if conf.DATABASE_AUTH_TOKEN == "" {
        log.Panic("environ DATABASE_AUTH_TOKEN not loaded")
    }

    conf.DATABASE_URL = os.Getenv("DATABASE_URL")
    if conf.DATABASE_URL == "" {
        log.Panic("environ DATABASE_URL not loaded")
    }

    conf.PORT = os.Getenv("PORT")
    if conf.PORT == "" {
        conf.PORT = "3000"
    }

    conf.CERT = os.Getenv("CERT")
    conf.KEY = os.Getenv("KEY")
    if conf.CERT == "" || conf.KEY == "" {
        log.Panic("environ CERTIFICATIONS not loaded\n",
            "cert: " + conf.CERT + "\n",
            "key: " + conf.KEY + "\n",
        )
    }

    a.Config = conf

    a.configureServo()
    a.configurePool()
}
func (a *Application) configureServo() {
    a.Servo = servo.NewServer(a.Config.PORT)
}

func (a *Application) configurePool() {
    var err error
    a.Pool, err = data.NewPool(a.Config.DATABASE_AUTH_TOKEN, a.Config.DATABASE_URL)
    if err != nil {
        log.Panic("Error building main connection pool")
    }
}

func (a *Application) Run() {
    go func() {
        log.Printf("Serve started at port `:%s`", a.Config.PORT)
        if err := a.Servo.ListenAndServeTLS(a.Config.CERT, a.Config.KEY); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Closed serve ...")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := a.Servo.Shutdown(ctx); err != nil {
        log.Fatalf("Error to off servo: %v", err)
    }
    log.Println("Servo closed correctly")
}
