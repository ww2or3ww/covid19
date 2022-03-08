package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"

	"app/utils/logger"
)

func init() {
	if os.Getenv("LOG_LEVEL") == "" {
		godotenv.Load(".env")
	}
	logLv := logger.Error
	envLogLv := os.Getenv("LOG_LEVEL")
	if envLogLv != "" {
		n, _ := strconv.Atoi(envLogLv)
		logLv = logger.LogLv(n)
	}
	logger.LogInitialize(logLv, 25)
}

func TestHandlerSuccess(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		path string
		ret  bool
	}{
		{
			name: "normal",
			path: "",
			ret:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := handler(tt.path)
			if err != nil {
				t.Errorf(tt.name)
			} else if ret != tt.ret {
				t.Errorf(tt.name)
			}
		})
	}
}
