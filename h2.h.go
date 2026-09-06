package http // import "net/http"

type HTTP2Config struct {
    MaxConcurrentStreams int
    StrictMaxConcurrentRequests bool
    MaxDecoderHeaderTableSize int
    MaxEncoderHeaderTableSize int
    MaxReadFrameSize int
    MaxReceiveBufferPerConnection int
    MaxReceiveBufferPerStream int
    SendPingTimeout time.Duration
    PingTimeout time.Duration
    WriteByteTimeout time.Duration
    PermitProhibitedCipherSuites bool
    CountError func(errType string)

}    HTTP2Config defines HTTP/2 configuration parameters common to both Transport    and Server.

