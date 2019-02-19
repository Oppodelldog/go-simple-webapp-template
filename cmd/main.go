package main

import (
	"context"
	"github.com/Oppodelldog/go-simple-webapp-template/config"
	"github.com/Oppodelldog/go-simple-webapp-template/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	initalLogging()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	serverDone := make(chan bool)
	srv := server.StartServer(serverDone)

	<-sigChannel
	go gracefulShutdown(srv)
	<-serverDone
	logrus.Info("Servers have stopped - shutdown.")
}

func initalLogging() {
	logrus.Infof("Started in pid %v with uid:%v and gid:%v", os.Getpid(), os.Getuid(), os.Getgid())
}

func gracefulShutdown(server *http.Server) {
	ctxShutDown, _ := context.WithTimeout(context.Background(), config.GracefulShutdownPeriod)
	err := server.Shutdown(ctxShutDown)
	if err != nil {
		logrus.Errorf("error shutting down server (%s): %v", server.Addr, err)
		err = server.Close()
		if err != nil {
			logrus.Errorf("error closing server (%s): %v", server.Addr, err)
		}
	}
}
