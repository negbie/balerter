package runner

import (
	"github.com/balerter/balerter/internal/modules"
	"github.com/balerter/balerter/internal/script/script"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
	"testing"
	"time"
)

type alertManagerMock struct {
	mock.Mock
}

func (m *alertManagerMock) Loader(s *script.Script) lua.LGFunction {
	args := m.Called(s)
	return args.Get(0).(lua.LGFunction)
}

type dsManagerMock struct {
	mock.Mock
}

func (m *dsManagerMock) Get() []modules.Module {
	args := m.Called()
	return args.Get(0).([]modules.Module)
}

type moduleMock struct {
	mock.Mock
}

func (m *moduleMock) Stop() error {
	args := m.Called()
	return args.Error(0)
}

func (m *moduleMock) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *moduleMock) GetLoader() lua.LGFunction {
	args := m.Called()
	return args.Get(0).(lua.LGFunction)
}

func TestRunner_createLuaState(t *testing.T) {

	m1 := &moduleMock{}
	m1.On("Name").Return("module1")
	m1.On("GetLoader").Return(func() lua.LGFunction {
		return func(state *lua.LState) int {
			return 0
		}
	}())

	dsManager := &dsManagerMock{}
	dsManager.On("Get").Return([]modules.Module{m1})

	alertManager := &alertManagerMock{}
	alertManager.On("Loader", mock.Anything).Return(func() lua.LGFunction {
		return func(state *lua.LState) int {
			return 0
		}
	}())

	rnr := &Runner{
		logger:       zap.NewNop(),
		dsManager:    dsManager,
		alertManager: alertManager,
	}

	L := rnr.createLuaState(&Job{name: "job1"})

	m1.AssertCalled(t, "Name")
	m1.AssertCalled(t, "GetLoader")

	require.NotNil(t, L)
	m1.AssertExpectations(t)
	dsManager.AssertExpectations(t)
}

func TestJob_Stop(t *testing.T) {
	j := &Job{
		stop: make(chan struct{}),
	}

	j.Stop()

	select {
	case <-j.stop:
	case <-time.After(time.Millisecond * 100):
		t.Fatal("wrong wait close a channel")
	}
}
