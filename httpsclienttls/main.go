package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	CACertFilePath     = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/ca-cert.pem"
	ClientCertFilePath = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/client-cert.pem"
	ClientKeyFilePath  = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/client-key.pem"
)

func httpsClient(url string) []byte {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(ClientCertFilePath, ClientKeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		return nil
	}
	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	if caCertPEM, err := ioutil.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	msg, _ := io.ReadAll(resp.Body)
	return msg
}

func main() {
	// hello, world.
	msg := httpsClient("https://host1.bastionxp.com:8080")
	fmt.Println("Msg: ", string(msg))
}
