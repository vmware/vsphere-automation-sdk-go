package utils

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

var httpClient http.Client

func GetHttpClientInstance() *http.Client {
	skipServerVerification := Arguments.Get("skip-server-verfication").(bool)
	truststorepath := Arguments.Get("certfile").(string)
	truststorekey := Arguments.Get("certkeyfile").(string)
	cert, err := tls.LoadX509KeyPair(truststorepath, truststorekey)
	if err != nil && !skipServerVerification {
		log.Fatal(err)
		os.Exit(1)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: skipServerVerification,
		Certificates:       []tls.Certificate{cert}}
	httpClient := http.Client{}
	return &httpClient
}
