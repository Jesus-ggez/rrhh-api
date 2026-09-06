package http // import "net/http"

type Server struct {
	Addr string

	Handler Handler // handler to invoke, http.DefaultServeMux if nil

	DisableGeneralOptionsHandler bool

	TLSConfig *tls.Config

	ReadTimeout time.Duration

	ReadHeaderTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration

	MaxHeaderBytes int

	MaxHeaderValueCount int

	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	ConnState func(net.Conn, ConnState)

	ErrorLog *log.Logger

	BaseContext func(net.Listener) context.Context

	ConnContext func(ctx context.Context, c net.Conn) context.Context

	HTTP2 *HTTP2Config

	Protocols *Protocols

	DisableClientPriority bool

}
    A Server defines parameters for running an HTTP server. The zero value for
    Server is a valid configuration.

func (s *Server) Close() error
func (s *Server) ListenAndServe() error
func (s *Server) ListenAndServeTLS(certFile, keyFile string) error
func (s *Server) RegisterOnShutdown(f func())
func (s *Server) Serve(l net.Listener) error
func (s *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
func (s *Server) SetKeepAlivesEnabled(v bool)
func (s *Server) Shutdown(ctx context.Context) error
