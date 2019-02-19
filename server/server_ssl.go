package server

import (
	"crypto/tls"
	"fmt"
	"github.com/Oppodelldog/go-simple-webapp-template/config"
	"net/http"
	"path"

	"github.com/sirupsen/logrus"
)

func StartServer(wait chan bool) *http.Server {

	addr := fmt.Sprintf("%s:%s", config.ServerAddressBindTLS, config.ServerPortTLS)
	logrus.Infof("Staring https server on %v", addr)

	srv := &http.Server{
		Addr:         addr,
		Handler:      NewRouter(),
		TLSConfig:    getTlsConfig(),
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	certFile := path.Join(config.CertDir, config.CertFile)
	keyFile := path.Join(config.CertDir, config.KeyFile)

	go func() {
		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil {
			logrus.Errorf("error running the server: %v", err)
			close(wait)
		}
	}()

	return srv
}

func getTlsConfig() *tls.Config {
	return &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
}
