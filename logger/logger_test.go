package logger

import (
	"fmt"
	"os"
	"testing"
)

func TestLogDebugf(t *testing.T) {
	LogFile = "./configtest.json"
	ReloadConfiguration()

	Debugf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if os.IsNotExist(err) {
		t.Errorf("Config test file should exist.")
	}

	remove(LogFile)
}

func TestLogDebugfWhenLogFileIsEmpty(t *testing.T) {
	LogFile = ""
	DebugEnabled = false
	ReloadConfiguration()

	Debugf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if !os.IsNotExist(err) {
		t.Errorf("Config test file should NOT exist.")
	}
}

func TestLogInfof(t *testing.T) {
	LogFile = "./configtest.json"
	ReloadConfiguration()

	Infof("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if os.IsNotExist(err) {
		t.Errorf("Config test file should exist.")
	}

	remove(LogFile)
}

func TestLogInfofWhenLogFileIsEmpty(t *testing.T) {
	LogFile = ""
	DebugEnabled = false
	ReloadConfiguration()

	Infof("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if !os.IsNotExist(err) {
		t.Errorf("Config test file should NOT exist.")
	}
}

func TestLogError(t *testing.T) {
	LogFile = "./configtest.json"
	ReloadConfiguration()

	Error("Try file")

	_, err := os.Stat(LogFile)
	if os.IsNotExist(err) {
		t.Errorf("Config test file should exist.")
	}

	remove(LogFile)
}

func TestLogErrorWhenLogFileIsEmpty(t *testing.T) {
	LogFile = ""
	DebugEnabled = false
	ReloadConfiguration()

	Error("Try file")

	_, err := os.Stat(LogFile)
	if !os.IsNotExist(err) {
		t.Errorf("Config test file should NOT exist.")
	}
}

func TestLogErrorf(t *testing.T) {
	LogFile = "./configtest.json"
	ReloadConfiguration()

	Errorf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if os.IsNotExist(err) {
		t.Errorf("Config test file should exist.")
	}

	remove(LogFile)
}

func TestLogErrorfWhenLogFileIsEmpty(t *testing.T) {
	LogFile = ""
	DebugEnabled = false
	ReloadConfiguration()

	Errorf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if !os.IsNotExist(err) {
		t.Errorf("Config test file should NOT exist.")
	}
}

func TestLogFatalf(t *testing.T) {
	LogFile = "./configtest.json"
	ReloadConfiguration()

	Fatalf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if os.IsNotExist(err) {
		t.Errorf("Config test file should exist.")
	}

	remove(LogFile)
}

func TestLogFatalfWhenLogFileIsEmpty(t *testing.T) {
	LogFile = ""
	DebugEnabled = false
	ReloadConfiguration()

	Fatalf("Try file %s",LogFile)

	_, err := os.Stat(LogFile)
	if !os.IsNotExist(err) {
		t.Errorf("Config test file should NOT exist.")
	}
}

func remove(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Could not remove config file: " + file)
	}
}
