package core

import "github.com/buke/quickjs-go"

type ModuleDefinition struct {
	Name string
	// Should be a Value, but can be anything
	Value quickjs.Value
}

// Modules are pure Go and have access to all functions
// But are instantiated on the Host platform
// Modules are any Native code that can be called from the JS side
// But for now, assume that JS only really interacts with Go,
// and for Go to interact with the the host platform, I will use Delegate pattern.
// This is because it's easier and relies on gomobile's auto type gen.
// Later, i can make the interface sexier like what Expo Modules API has done.
type Module interface {
	Define(ctx *quickjs.Context) *ModuleDefinition
}

type HostModuleDefinition struct {
	name string
}

func NewHostModuleDefinition(name string) *HostModuleDefinition {
	return &HostModuleDefinition{name}
}

// This is so host platforms can make their own modules
type HostModule interface {
	Define(ctx *HostContextWrapper) *HostModuleDefinition
}

// Helper so that host platforms can create their own modules
type HostContextWrapper struct {
	ctx *quickjs.Context
	obj *quickjs.Value
}
