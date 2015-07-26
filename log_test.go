package log

import (
	"bytes"
	"strings"
	"testing"
)

const errLoggingLevels = "Logging levels are not being respected."

func TestLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, ERROR)

	if log.level != ERROR {
		t.Error("Was not set correctly")
	}

	log.Info("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "Visible Level"
	log.Error(msg)
	line = buf.String()
	if !strings.Contains(line, msg) {
		t.Error("Expected:", msg, "Got:", line)
	}
}
func TestFatal(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, FATAL)

	log.Error("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World is Ending"
	defer func() {
		if r := recover(); r != nil {
			line = buf.String()
			if !strings.Contains(line, "FATAL: ["+msg+"]") {
				t.Error("Expected:", msg, "Got:", line)
			}

		} else {
			t.Error("Expected a panic")
		}
	}()

	log.Fatal(msg)
}

func TestStdLogFatal(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(FATAL)
	Error("Hidden Level")

	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World is Ending"
	defer func() {
		if r := recover(); r != nil {
			line = buf.String()
			if !strings.Contains(line, "FATAL: ["+msg+"]") {
				t.Error("Expected:", msg, "Got:", line)
			}

		} else {
			t.Error("Expected a panic")
		}
	}()
	Fatal(msg)
}

func TestLogError(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, ERROR)

	log.Warn("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	log.Error(msg)
	line = buf.String()
	if !strings.Contains(line, "ERROR: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestStdLogError(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(ERROR)

	Warn("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	Error(msg)
	line = buf.String()
	if !strings.Contains(line, "ERROR: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestLogWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, WARN)

	log.Info("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	log.Warn(msg)
	line = buf.String()
	if !strings.Contains(line, "WARN: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestStdLogWarn(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(WARN)

	Info("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	Warn(msg)
	line = buf.String()
	if !strings.Contains(line, "WARN: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestLogInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, INFO)

	log.Debug("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	log.Info(msg)
	line = buf.String()
	if !strings.Contains(line, "INFO: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestStdLogInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(INFO)

	Debug("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	Info(msg)
	line = buf.String()
	if !strings.Contains(line, "INFO: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestLogDebug(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, DEBUG)

	log.Trace("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	log.Debug(msg)
	line = buf.String()
	if !strings.Contains(line, "DEBUG: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestStdLogDebug(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(DEBUG)

	Trace("Hidden Level")
	line := buf.String()
	if line != "" {
		t.Fatal(errLoggingLevels)
	}

	msg := "World could be ending"
	Debug(msg)
	line = buf.String()
	if !strings.Contains(line, "DEBUG: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestLogTrace(t *testing.T) {
	buf := new(bytes.Buffer)
	log := NewLogger(buf, TRACE)

	msg := "World could be ending"
	log.Trace(msg)
	line := buf.String()
	if !strings.Contains(line, "TRACE: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}

	log.Debug(msg)
	line = buf.String()
	if !strings.Contains(line, "DEBUG: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}

func TestStdLogTrace(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	Level(TRACE)

	msg := "World could be ending"
	Trace(msg)
	line := buf.String()
	if !strings.Contains(line, "TRACE: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}

	Debug(msg)
	line = buf.String()
	if !strings.Contains(line, "DEBUG: ["+msg+"]") {
		t.Error("Expected:", msg, "Got:", line)
	}
}
