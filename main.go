//go:build js && wasm

package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	js.Global().Set("goExec", goExec)

	// block indefinitely, as expected by Go WASM
	select {}
}

// Execute arbitary Go code.
//
// REALLY NOT SAFE!!
//
// Seriously, do **_NOT_** use this.
func exec(code string) (string, string) {
	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)

	i := interp.New(interp.Options{
		Stdout: stdoutBuf,
		Stderr: stderrBuf,
	})

	i.Use(stdlib.Symbols)

	_, err := i.Eval(code)
	if err != nil {
		stderrBuf.WriteString(fmt.Sprintf("err: %v", err.Error()))
	}

	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()
	return stdout, stderr
}

var goExec = js.FuncOf(func(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		panic("want one argument")
	}

	stdout, stderr := exec(args[0].String())
	return js.ValueOf(map[string]interface{}{
		"stdout": stdout,
		"stderr": stderr,
	})
})
