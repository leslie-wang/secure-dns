package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// load client cert
	cert, err := tls.LoadX509KeyPair("../../certs/tls/client.crt", "../../certs/tls/client.key")
	if err != nil {
		fmt.Println("test")
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile("../../certs/tls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	letsCert, err := ioutil.ReadFile("../../certs/https/fullchain1.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	caCertPool.AppendCertsFromPEM(letsCert)

	// https client tls config
	// InsecureSkipVerify true means not validate server certificate (so no need to set RootCAs)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		//InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	// https client request
	url := "https://www.lomorage.com"
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{Transport: transport}

	// read response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	log.Println(string(contents))
}
