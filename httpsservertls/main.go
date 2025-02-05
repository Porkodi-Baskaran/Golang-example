package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	CACertFilePath = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/ca-cert.pem"
	CertFilePath   = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/server-cert.pem"
	KeyFilePath    = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/server-key.pem"
)

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}
func main() {
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}
	// Configure the server to trust TLS client cert issued by your CA.
	certPool := x509.NewCertPool()
	if caCertPEM, err := ioutil.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		Certificates: []tls.Certificate{serverTLSCert},
	}
	server := http.Server{
		Addr:      "host1.bastionxp.com:8080",
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}
	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
