package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestLoggerImpl struct {
	ErrorOutput []string
	WarnOutput  []string
	InfoOutput  []string
	DebugOutput []string
	TraceOutput []string
}

func (l *TestLoggerImpl) Error(msg string) {
	l.ErrorOutput = append(l.ErrorOutput, msg)
}

func (l *TestLoggerImpl) Warn(msg string) {
	l.WarnOutput = append(l.WarnOutput, msg)
}

func (l *TestLoggerImpl) Info(msg string) {
	l.InfoOutput = append(l.InfoOutput, msg)
}

func (l *TestLoggerImpl) Debug(msg string) {
	l.DebugOutput = append(l.DebugOutput, msg)
}

func (l *TestLoggerImpl) Trace(msg string) {
	l.TraceOutput = append(l.TraceOutput, msg)
}

func TestError_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Error("test error")

	assert.Len(t, l.ErrorOutput, 1)
	assert.Equal(t, "test error", l.ErrorOutput[0])
}

func TestErrorf_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Errorf("test error %s", "format")

	assert.Len(t, l.ErrorOutput, 1)
	assert.Equal(t, "test error format", l.ErrorOutput[0])
}

func TestWarn_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Warn("test warning")

	assert.Len(t, l.WarnOutput, 1)
	assert.Equal(t, "test warning", l.WarnOutput[0])
}

func TestWarnf_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Warnf("test warning %s", "format")

	assert.Len(t, l.WarnOutput, 1)
	assert.Equal(t, "test warning format", l.WarnOutput[0])
}

func TestInfo_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Info("test info")

	assert.Len(t, l.InfoOutput, 1)
	assert.Equal(t, "test info", l.InfoOutput[0])
}

func TestInfof_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Infof("test info %s", "format")

	assert.Len(t, l.InfoOutput, 1)
	assert.Equal(t, "test info format", l.InfoOutput[0])
}

func TestDebug_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Debug("test debug")

	assert.Len(t, l.DebugOutput, 1)
	assert.Equal(t, "test debug", l.DebugOutput[0])
}

func TestDebugf_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Debugf("test debug %s", "format")

	assert.Len(t, l.DebugOutput, 1)
	assert.Equal(t, "test debug format", l.DebugOutput[0])
}

func TestTrace_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Trace("test trace")

	assert.Len(t, l.TraceOutput, 1)
	assert.Equal(t, "test trace", l.TraceOutput[0])
}

func TestTracef_LogEntryAdded(t *testing.T) {
	l := &TestLoggerImpl{}
	SetLogger(l)

	Tracef("test trace %s", "format")

	assert.Len(t, l.TraceOutput, 1)
	assert.Equal(t, "test trace format", l.TraceOutput[0])
}
