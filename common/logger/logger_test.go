package logger

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/logger/mocks"
)

func newMockLogger() (*Logger, *mocks.BaseLogger) {
	c := NewConfig()
	m := &mocks.BaseLogger{}

	return &Logger{logger: m, config: c}, m
}

func TestLogger_Debug(t *testing.T) {
	l, m := newMockLogger()
	m.On("Debug", "debug").Return(nil).Once()
	l.Debug("debug")
	m.AssertExpectations(t)
}

func TestLogger_Debugf(t *testing.T) {
	l, m := newMockLogger()
	m.On("Debugf", "%s", "debug").Return(nil).Once()
	l.Debugf("%s", "debug")
	m.AssertExpectations(t)
}

func TestLogger_Info(t *testing.T) {
	l, m := newMockLogger()
	m.On("Info", "info").Return(nil).Once()
	l.Info("info")
	m.AssertExpectations(t)
}

func TestLogger_Infof(t *testing.T) {
	l, m := newMockLogger()
	m.On("Infof", "%s", "info").Return(nil).Once()
	l.Infof("%s", "info")
	m.AssertExpectations(t)
}

func TestLogger_Error(t *testing.T) {
	l, m := newMockLogger()
	m.On("Error", "error").Return(nil).Once()
	l.Error("error")
	m.AssertExpectations(t)
}

func TestLogger_Errorf(t *testing.T) {
	l, m := newMockLogger()
	m.On("Errorf", "%s", "error").Return(nil).Once()
	l.Errorf("%s", "error")
	m.AssertExpectations(t)
}

func TestLogger_Fatal(t *testing.T) {
	l, m := newMockLogger()
	m.On("Fatal", "fatal").Return(nil).Once()
	l.Fatal("fatal")
	m.AssertExpectations(t)
}

func TestLogger_Fatalf(t *testing.T) {
	l, m := newMockLogger()
	m.On("Fatalf", "%s", "fatal").Return(nil).Once()
	l.Fatalf("%s", "fatal")
	m.AssertExpectations(t)
}

// Should create a new logger.
func TestNewLogger(t *testing.T) {
	l := NewLogger()
	assert.NotNil(t, l)
	assert.NotNil(t, l.config)
	assert.NotNil(t, l.logger)
}

// Should panic.
func TestNewLoggerPanic(t *testing.T) {
	o := viper.Get(filePathKey)
	defer viper.Set(filePathKey, o)

	viper.Set(filePathKey, "")
	assert.Panics(t, func() {
		NewLogger()
	})

	viper.Set(filePathKey, "/invalid/path")
	assert.Panics(t, func() {
		NewLogger()
	})
}
