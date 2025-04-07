package errkit

import (
	"os"

	"github.com/ed-henrique/voz/internal/logger"
)

func FinalErr(err error) {
	logger.Error("error was final", err.Error())
	os.Exit(1)
}
