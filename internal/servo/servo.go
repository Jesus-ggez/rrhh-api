package servo

import (
    "context"
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "time"
)

func NewServer(port string) *http.Server {
    mux := http.NewServeMux()

    sv := &http.Server{
        Addr:                         ":" + port,
        Handler:                      mux,
        DisableGeneralOptionsHandler: false,

        ReadHeaderTimeout: 10 * time.Second,
        WriteTimeout:      20 * time.Second,
        ReadTimeout:       15 * time.Second,
        IdleTimeout:       10 * time.Second,

        MaxHeaderBytes:      1 << 16, // 64kb
        MaxHeaderValueCount: 20,

        // servo logger
        ErrorLog: log.New(os.Stderr, "HTTP-ERR: ", log.LstdFlags|log.Lshortfile),

        // custom protocolos handle
        TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){
            "http/1.1": nil, // use default handler
        },

        // connection state
        ConnState: func(nc net.Conn, cs http.ConnState) {
            switch cs {
            case http.StateNew:
                log.Printf("New connection: %s", nc.RemoteAddr())
            case http.StateActive:
                log.Printf("Active connection: %s", nc.RemoteAddr())
            case http.StateClosed:
                log.Printf("Closed connection: %s", nc.RemoteAddr())
            case http.StateHijacked:
                log.Printf("Hj connection: %s", nc.RemoteAddr())
            case http.StateIdle:
                log.Printf("IDLE connection: %s", nc.RemoteAddr())
            }
        },

        BaseContext: func(nl net.Listener) context.Context {
            ctx := context.Background()
            ctx = context.WithValue(ctx, "servoStart", time.Now())
            ctx = context.WithValue(ctx, "listenerAddr", nl.Addr().String())
            return ctx
        },

        ConnContext: func(ctx context.Context, c net.Conn) context.Context {
            ctx = context.WithValue(ctx, "remoteAddr", c.RemoteAddr().String())
            ctx = context.WithValue(ctx, "connID", newConnID())
            return ctx
        },
        DisableClientPriority: false,

        // http/2 config
        HTTP2: NewHTTP2(),

        // tls config
        TLSConfig: NewTLSConf(),
    }
    return sv
}

var newConnID = func() string { return fmt.Sprintf("%d", time.Now().UnixNano()) }
