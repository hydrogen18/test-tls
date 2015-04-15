package common

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
)

func MustLoadCertificates() (tls.Certificate, *x509.CertPool) {
	l := len(os.Args)
	privateKeyFile := os.Args[l-3]
	certificateFile := os.Args[l-2]
	caFile := os.Args[l-1]

	mycert, err := tls.LoadX509KeyPair(certificateFile, privateKeyFile)
	if err != nil {
		panic(err)
	}

	pem, err := ioutil.ReadFile(caFile)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pem) {
		panic("Failed appending certs")
	}

	return mycert, certPool

}

func MustGetTlsConfiguration() *tls.Config {
	config := &tls.Config{}
	mycert, certPool := MustLoadCertificates()
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0] = mycert

	config.RootCAs = certPool
	config.ClientCAs = certPool

	config.ClientAuth = tls.RequireAndVerifyClientCert

	//Optional stuff

	//Use only modern ciphers
	config.CipherSuites = []uint16{tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256}

	//Use only TLS v1.2
	config.MinVersion = tls.VersionTLS12

	//Don't allow session resumption
	config.SessionTicketsDisabled = true
	return config
}
