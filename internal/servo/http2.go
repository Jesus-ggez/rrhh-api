package servo

import "net/http"

func NewHTTP2() *http.HTTP2Config {
    return &http.HTTP2Config{
        MaxConcurrentStreams: 250,
        StrictMaxConcurrentRequests: true,

        MaxDecoderHeaderTableSize: 4096,
        MaxEncoderHeaderTableSize: 4096,

        MaxReadFrameSize: 1 << 20, // 1mb
        /*
        MaxReceiveBufferPerConnection: int,
        MaxReceiveBufferPerStream: int,
        SendPingTimeout: time.Duration,
        PingTimeout: time.Duration,
        WriteByteTimeout: time.Duration,
        PermitProhibitedCipherSuites: bool,
        CountError func(errType: string),
        */
    }
}
