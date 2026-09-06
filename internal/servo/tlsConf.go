package servo

import "crypto/tls"

func NewTLSConf() *tls.Config {
    return &tls.Config{
        /*
        me da weba conigurar esto en ese momento

        TODO: trace:x.?rdr.on


           Rand: io.Reader,
           Time func(): time.Time,
           Certificates: []Certificate,
           NameToCertificate: map[string]*Certificate,
           GetCertificate func(*ClientHelloInfo) (*Certificate,: error),
           GetClientCertificate func(*CertificateRequestInfo) (*Certificate,: error),
           GetConfigForClient func(*ClientHelloInfo) (*Config,: error),
           VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate): error,
           VerifyConnection func(ConnectionState): error,
           RootCAs: *x509.CertPool,
           NextProtos: []string,
           ServerName: string,
           ClientAuth: ClientAuthType,
           ClientCAs: *x509.CertPool,
           InsecureSkipVerify: bool,
           CipherSuites: []uint16,
           PreferServerCipherSuites: bool,
           SessionTicketsDisabled: bool,
           SessionTicketKey: [32]byte,
           ClientSessionCache: ClientSessionCache,
           UnwrapSession func(identity []byte, cs ConnectionState) (*SessionState,: error),
           WrapSession func(ConnectionState, *SessionState) ([]byte,: error),
           MinVersion: uint16,
           MaxVersion: uint16,
           CurvePreferences: []CurveID,
           DynamicRecordSizingDisabled: bool,
           Renegotiation: RenegotiationSupport,
           KeyLogWriter: io.Writer,
           EncryptedClientHelloConfigList: []byte,
           EncryptedClientHelloRejectionVerify func(ConnectionState): error,
           GetEncryptedClientHelloKeys func(*ClientHelloInfo) ([]EncryptedClientHelloKey,: error),
           EncryptedClientHelloKeys: []EncryptedClientHelloKey,
        */
    }
}
