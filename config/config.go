package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// ServerAddressBindTLS defines the network interface to bind the https listener to
var ServerAddressBindTLS string

// ServerPortTLS for secure API like hooks or REST access
var ServerPortTLS string

//CertDir defines where to put certificate files
var CertDir string

//CertFile filepath to certificate
var CertFile string

//KeyFile filepath to certificate key
var KeyFile string

//ReadTimeout used as http server config
var ReadTimeout time.Duration

//WriteTimeout used as http server config
var WriteTimeout time.Duration

//IdleTimeout used as http server config
var IdleTimeout time.Duration

//GracefulShutdownPeriod seconds to wait for api server to shutdown while shutting down the service
var GracefulShutdownPeriod time.Duration

//UseStaticFiles may be set to false during development to not use static assets
var UseStaticFiles bool

//AbsoluteAssetsPath must be set when UseStaticFiles is set to false
var AbsoluteAssetsPath string

const EnvVarPrefix = "MYAPP"

func init() {
	ServerPortTLS = getEnv("SERVER_PORT_TLS", "10443")
	ServerAddressBindTLS = getEnv("ADDRESS_BIND_TLS", "127.0.0.1")

	CertDir = getEnv("CERT_DIR", ".cert")
	CertFile = getEnv("CERT_FILE", "sample.crt")
	KeyFile = getEnv("KEY_FILE", "sample.key")

	ReadTimeout = getDurationFromEnv("READ_TIMEOUT", time.Second*5)
	WriteTimeout = getDurationFromEnv("WRITE_TIMEOUT", time.Second*5)
	IdleTimeout = getDurationFromEnv("READ_TIMEOUT", time.Second*120)
	GracefulShutdownPeriod = getDurationFromEnv("SHUTDOWN_TIMEOUT", time.Second*6)
	IdleTimeout = getDurationFromEnv("READ_TIMEOUT", time.Second*120)

	logLevel := getEnv("LOG_LEVEL", "debug")
	loglevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(loglevel)
	}

	UseStaticFiles = getEnvBool("USE_STATIC_FILES", true)
	AbsoluteAssetsPath = getEnv("ABSOLUTE_ASSETS_PATH", "")
}

func getEnvBool(key string, def bool) bool {
	key = fmt.Sprintf("%s_%s", EnvVarPrefix, key)

	if _, ok := os.LookupEnv(key); !ok {
		return def
	}

	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return def
	}
	return val

}

func getEnv(key, def string) string {
	key = fmt.Sprintf("%s_%s", EnvVarPrefix, key)
	if s, ok := os.LookupEnv(key); ok {
		return s
	} else {
		return def
	}
}

func getDurationFromEnv(key string, def time.Duration) time.Duration {
	key = fmt.Sprintf("%s_%s", EnvVarPrefix, key)
	if value, ok := os.LookupEnv(key); ok {
		iVal, err := strconv.Atoi(value)
		if err != nil {
			logrus.Warnf("could not parse duration from %s: %v", key, err)
			return def
		} else {
			return time.Duration(iVal) * time.Second
		}
	}

	return def
}
