package logger

import (
	"testing"
	"time"
)

func TestNewZapConsoleLogger(t *testing.T) {
	var scenarios = []struct {
		Name  string
		Level string
	}{
		{
			Name:  "TEST INFO LOGGING",
			Level: "info",
		},
		{
			Name:  "TEST WARN LOGGING",
			Level: "warn",
		},
		{
			Name:  "TEST ERROR LOGGING",
			Level: "error",
		},
		{
			Name:  "TEST DEFAULT LOGGING",
			Level: "other",
		},
	}
	for _, scenario := range scenarios {
		logger := NewZapConsoleLogger(scenario.Level)

		switch scenario.Level {
		case "info":
			logger.Info("INFO")
			t.Logf("PASSED: %s", scenario.Name)
		case "warn":
			logger.Warn("WARN")
			t.Logf("PASSED: %s", scenario.Name)
		case "error":
			logger.Error("ERROR")
			t.Logf("PASSED: %s", scenario.Name)
		default:
			logger.Info("INFO")
			t.Logf("PASSED: %s", scenario.Name)
		}
	}
}

func TestNewZapFileLogger(t *testing.T) {
	logsPath := t.TempDir() + "temp.logs"
	var scenarios = []struct {
		Name  string
		Level string
		Path  string
	}{
		{
			Name:  "TEST DEBUG LOGGING",
			Level: "debug",
			Path:  logsPath,
		},
		{
			Name:  "TEST INFO LOGGING",
			Level: "info",
			Path:  logsPath,
		},
		{
			Name:  "TEST WARN LOGGING",
			Level: "warn",
			Path:  logsPath,
		},
		{
			Name:  "TEST ERROR LOGGING",
			Level: "error",
			Path:  logsPath,
		},
		{
			Name:  "TEST DEFAULT LOGGING",
			Level: "other",
			Path:  logsPath,
		},
	}
	for _, scenario := range scenarios {
		logger := NewZapFileLogger(scenario.Path, scenario.Level, 1, 1, 1)

		switch scenario.Level {
		case "debug":
			logger.Debug("DEBUG")
			t.Logf("PASSED: %s", scenario.Name)
		case "info":
			logger.Info("INFO")
			t.Logf("PASSED: %s", scenario.Name)
		case "warn":
			logger.Warn("WARN")
			t.Logf("PASSED: %s", scenario.Name)
		case "error":
			logger.Error("ERROR")
			t.Logf("PASSED: %s", scenario.Name)
		default:
			logger.Info("INFO")
			t.Logf("PASSED: %s", scenario.Name)
		}
	}
	time.Sleep(5 * time.Second)
}
