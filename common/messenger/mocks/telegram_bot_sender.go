// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	mock "github.com/stretchr/testify/mock"
)

// TelegramBotSender is an autogenerated mock type for the TelegramBotSender type
type TelegramBotSender struct {
	mock.Mock
}

// Send provides a mock function with given fields: c
func (_m *TelegramBotSender) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	ret := _m.Called(c)

	var r0 tgbotapi.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Chattable) (tgbotapi.Message, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.Chattable) tgbotapi.Message); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(tgbotapi.Message)
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.Chattable) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTelegramBotSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewTelegramBotSender creates a new instance of TelegramBotSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTelegramBotSender(t mockConstructorTestingTNewTelegramBotSender) *TelegramBotSender {
	mock := &TelegramBotSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}