package core

import "github.com/buke/quickjs-go"

type GoModuleDefinition struct {
	Name string
	// Should be a Value, but can be anything
	Value quickjs.Value
}

// Modules are pure Go and have access to all functions
// But are instantiated on the Host platform
// Modules are any Native code that can be called from the JS side
type GoModule interface {
	Define(ctx *quickjs.Context) GoModuleDefinition
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
