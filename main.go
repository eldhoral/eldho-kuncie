package main

import (
	"os"

	"github.com/eldhoral/eldho-kuncie/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Errorln("error on command execution", err.Error())
		os.Exit(1)
	}
}
