// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: FilterRunner)

package pipeline

import (
	sync "sync"
	gomock "code.google.com/p/gomock/gomock"
	time "time"
)

// Mock of FilterRunner interface
type MockFilterRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockFilterRunnerRecorder
}

// Recorder for MockFilterRunner (not exported)
type _MockFilterRunnerRecorder struct {
	mock *MockFilterRunner
}

func NewMockFilterRunner(ctrl *gomock.Controller) *MockFilterRunner {
	mock := &MockFilterRunner{ctrl: ctrl}
	mock.recorder = &_MockFilterRunnerRecorder{mock}
	return mock
}

func (_m *MockFilterRunner) EXPECT() *_MockFilterRunnerRecorder {
	return _m.recorder
}

func (_m *MockFilterRunner) Deliver(_param0 *PipelinePack) {
	_m.ctrl.Call(_m, "Deliver", _param0)
}

func (_mr *_MockFilterRunnerRecorder) Deliver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Deliver", arg0)
}

func (_m *MockFilterRunner) Filter() Filter {
	ret := _m.ctrl.Call(_m, "Filter")
	ret0, _ := ret[0].(Filter)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Filter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Filter")
}

func (_m *MockFilterRunner) InChan() chan *PipelineCapture {
	ret := _m.ctrl.Call(_m, "InChan")
	ret0, _ := ret[0].(chan *PipelineCapture)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) InChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InChan")
}

func (_m *MockFilterRunner) Inject(_param0 *PipelinePack) bool {
	ret := _m.ctrl.Call(_m, "Inject", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Inject(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Inject", arg0)
}

func (_m *MockFilterRunner) LogError(_param0 error) {
	_m.ctrl.Call(_m, "LogError", _param0)
}

func (_mr *_MockFilterRunnerRecorder) LogError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogError", arg0)
}

func (_m *MockFilterRunner) LogMessage(_param0 string) {
	_m.ctrl.Call(_m, "LogMessage", _param0)
}

func (_mr *_MockFilterRunnerRecorder) LogMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogMessage", arg0)
}

func (_m *MockFilterRunner) MatchRunner() *MatchRunner {
	ret := _m.ctrl.Call(_m, "MatchRunner")
	ret0, _ := ret[0].(*MatchRunner)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) MatchRunner() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MatchRunner")
}

func (_m *MockFilterRunner) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockFilterRunner) Plugin() Plugin {
	ret := _m.ctrl.Call(_m, "Plugin")
	ret0, _ := ret[0].(Plugin)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Plugin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Plugin")
}

func (_m *MockFilterRunner) PluginGlobals() *PluginGlobals {
	ret := _m.ctrl.Call(_m, "PluginGlobals")
	ret0, _ := ret[0].(*PluginGlobals)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) PluginGlobals() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PluginGlobals")
}

func (_m *MockFilterRunner) RetainPack(_param0 *PipelineCapture) {
	_m.ctrl.Call(_m, "RetainPack", _param0)
}

func (_mr *_MockFilterRunnerRecorder) RetainPack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetainPack", arg0)
}

func (_m *MockFilterRunner) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockFilterRunnerRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}

func (_m *MockFilterRunner) Start(_param0 PluginHelper, _param1 *sync.WaitGroup) error {
	ret := _m.ctrl.Call(_m, "Start", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}

func (_m *MockFilterRunner) Ticker() <-chan time.Time {
	ret := _m.ctrl.Call(_m, "Ticker")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

func (_mr *_MockFilterRunnerRecorder) Ticker() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ticker")
}
