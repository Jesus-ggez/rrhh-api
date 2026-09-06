package servo

import (
    "net/http"
    "time"
)

func NewHTTP2() *http.HTTP2Config {
    return &http.HTTP2Config{
        MaxConcurrentStreams:        250,
        StrictMaxConcurrentRequests: true,

        MaxDecoderHeaderTableSize: 4096,
        MaxEncoderHeaderTableSize: 4096,

        // 1mb == 1 << 20
        MaxReceiveBufferPerConnection: 1 << 20,
        MaxReceiveBufferPerStream:     1 << 20,
        MaxReadFrameSize:              1 << 20,

        SendPingTimeout: 10 * time.Second,
        PingTimeout:     15 * time.Second,

        WriteByteTimeout: 30 * time.Second,

        PermitProhibitedCipherSuites: false,

        // nil
        // CountError: func(errType string) {
        /*
           Use configure
           UnImplemented for this case
        */
        // },
    }
}
