// This file was generated by counterfeiter
// counterfeiter -o ios/iexec/exec_fake.go --fake-name Fake ios/iexec/exec.go Exec

package iexec

import (
	"bytes"
	"io/ioutil"
	"sync"

	"github.com/BooleanCat/igo/ios"
)

//Fake ...
type Fake struct {
	CommandStub        func(string, ...string) Cmd
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	commandReturns struct {
		result1 Cmd
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//NewFake is the preffered way to initialise a Fake
func NewFake() *Fake {
	return new(Fake)
}

//Command ...
func (fake *Fake) Command(arg1 string, arg2 ...string) Cmd {
	fake.commandMutex.Lock()
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("Command", []interface{}{arg1, arg2})
	fake.commandMutex.Unlock()
	if fake.CommandStub != nil {
		return fake.CommandStub(arg1, arg2...)
	}
	return fake.commandReturns.result1
}

//CommandCallCount ...
func (fake *Fake) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

//CommandArgsForCall ...
func (fake *Fake) CommandArgsForCall(i int) (string, []string) {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return fake.commandArgsForCall[i].arg1, fake.commandArgsForCall[i].arg2
}

//CommandReturns ...
func (fake *Fake) CommandReturns(result1 Cmd) {
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 Cmd
	}{result1}
}

//Invocations ...
func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return fake.invocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

/*
NestedCommandFake is a struct containing a fake Exec with nested initialised fake
members.

The following Fakes are available:
- Cmd: a FakeCmd returned by Exec.Command()
- Process: a FakeProcess returned by Cmd.GetProcess()
*/
type NestedCommandFake struct {
	Exec    *Fake
	Cmd     *CmdFake
	Process *ios.ProcessFake
}

//NewNestedCommandFake returns a fake Exec with nested new fakes within
func NewNestedCommandFake() *NestedCommandFake {
	execFake := NewFake()
	processFake := new(ios.ProcessFake)
	cmdFake := newNestedCmdFake(processFake)
	execFake.CommandReturns(cmdFake)
	return &NestedCommandFake{
		Exec:    execFake,
		Cmd:     cmdFake,
		Process: processFake,
	}
}

func newNestedCmdFake(process ios.Process) *CmdFake {
	fake := NewCmdFake()
	fake.GetProcessReturns(process)
	fake.StdoutPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	fake.StderrPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	return fake
}

var _ Exec = new(Fake)
