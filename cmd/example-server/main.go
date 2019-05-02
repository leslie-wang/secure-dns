package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func main() {
	// load client cert
	cert, err := tls.LoadX509KeyPair("../../certs/https/fullchain1.pem", "../../certs/https/privkey1.pem")
	if err != nil {
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile("../../certs/tls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// https client tls config
	// InsecureSkipVerify true means not validate server certificate (so no need to set RootCAs)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()

	l, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:      ":443",
		Handler:   helloHandler{},
		TLSConfig: tlsConfig,
	}

	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
