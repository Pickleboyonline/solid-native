package core

import (
	"fmt"
	"io"
	"net/http"

	"github.com/buke/quickjs-go"
	polyfill "github.com/buke/quickjs-go-polyfill"
)

type Core struct {
	rt  *quickjs.Runtime
	ctx *quickjs.Context
}

const GlobalCoreName = "_SolidNativeCore"

func NewCore() *Core {
	rt := quickjs.NewRuntime(quickjs.WithMaxStackSize(1024 * 1024))
	ctx := rt.NewContext()
	// TODO: Register(_SolidNativeCore into global this so it can grab modules)
	polyfill.InjectAll(ctx)
	/*
		Needed function def:
		- Just a way to get the module honestly, so something like SolidNativeCore.modules dictionary
	*/

	// Define SolidNativeCore
	core := ctx.Object()

	modules := ctx.Object()

	core.Set("modules", modules)

	ctx.Globals().Set(GlobalCoreName, core)

	return &Core{
		rt:  &rt,
		ctx: ctx,
	}
}

// Basically, add it to the the decision. It takes in the context and can do whatever it pleases
// Just make sure to return an object (not sure if i want to )
// Should be called on host platform (for example, should pass in Renderer)
// For now, we'll start with pure Go module, but then i will look at a HelloWorld module where the full implementation is
// in the host. The reasoning is that in order for the Host platform to have access, I need to wrap the context in something exposal
// for the host.
func (c *Core) RegisterGoModule(module Module) {
	definition := module.Define(c.ctx)
	c.ctx.Globals().Get(GlobalCoreName).Get("modules").Set(definition.Name, definition.Value)
}

func (c *Core) RegisterHostModule(module HostModule) {
	obj := c.ctx.Object()
	wrapper := &HostContextWrapper{
		ctx: c.ctx,
		obj: &obj,
	}
	definition := module.Define(wrapper)
	c.ctx.Globals().Get(GlobalCoreName).Get("modules").Set(definition.name, *wrapper.obj)
}

func (c *Core) Free() {
	c.ctx.Close()
	c.rt.Close()
}

// Should be called on a background thread when starting from the host platform
func (c *Core) StartFromServer(url string) error {
	// Make a GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Print the retrieved text
	jsToEval := string(body)

	jsValue, err := c.ctx.Eval(jsToEval)
	defer jsValue.Free()

	return err
}

func (c *Core) StartFromJS(js string) error {
	jsValue, err := c.ctx.Eval(js)
	defer jsValue.Free()

	return err
}

// TODO: Debugger connects to server for JS runtime
// From the perspective of the mobile device, the JS will run on the dev server
// running and debugging in the deno runtime.
// Basically all modules commication will be done with sockets for bidirection communication
// Will need to abstract way the server communicated with modules for this.
func (c *Core) StartWithDebugger(debugServerUrl string) error {
	return nil
}
