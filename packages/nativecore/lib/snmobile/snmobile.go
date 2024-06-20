/*
Package serves as a way to interface with iOS and Android
TODO: Mobile will create an object here.
*/
package snmobile

import (
	"fmt"
	"nativecore/lib/yoga"

	"gopkg.in/olebedev/go-duktape.v3"
)

// Houses important info.
type SolidNativeMobile struct {
	// Get chidren
	nodeChildren  map[int][]int
	yogaNodes     map[int]*yoga.YGNode
	nodeStyleKeys map[int]Set
	nodeParent    map[int]int
	hostReceiver  HostReceiver
	dukContext    *duktape.Context
	// Set to -1 initally, need to set before calulcating layouts
	rootNodeId       *int
	deviceScreenSize *Size
}

func NewSolidNativeMobile(hostReceiver HostReceiver) *SolidNativeMobile {
	ctx := duktape.New()

	// ctx.PushGoFunction()
	return &SolidNativeMobile{
		yogaNodes:    make(map[int]*yoga.YGNode),
		hostReceiver: hostReceiver,
		dukContext:   ctx,
		rootNodeId:   nil,
		nodeParent:   make(map[int]int),
		nodeChildren: make(map[int][]int),
		// We use this to keep track of when keys are removed
		// Since Yoga works via mutation
		nodeStyleKeys:    make(map[int]Set),
		deviceScreenSize: hostReceiver.GetDeviceScreenSize(),
	}
}

// TODO: Give iterator type to retreive all
// native modules with reciever function.
// May need to use flex for type conversion here.
func (s *SolidNativeMobile) RegistureModules() {
	s.dukContext.PushTimers()
	s.dukContext.PushGlobalGoFunction("log", func(c *duktape.Context) int {
		fmt.Println(c.SafeToString(-1))
		return 0
	})
	s.registureRenderer()
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) RunJsFromServer(url string) error {
	s.RegistureModules()
	return s.downloadAndRunJs(url)
}

// Registure core into system, download, and runs js.
func (s *SolidNativeMobile) EvalJs(jsToEval string) error {
	s.RegistureModules()
	return s.dukContext.PevalString(jsToEval)
}

// Yoga and JSContext need to be cleaned up before this object is deallocated
func (s *SolidNativeMobile) FreeMemory() {
	s.dukContext.DestroyHeap()
}

// TODO: determine how this is implimneted and if i can just update the device dimensions
func (s *SolidNativeMobile) OnOrientationChange() {

}

// Create root node and return its ID.
// Not to be called on JS
// This removes the callback like effect and allows the host to create its root node immediatly
// to present it to the screen.
//
// Use the nodetype to tell whether we need to measure it or not
func (s *SolidNativeMobile) CreateRootNode(nodeType string) int {
	nodeId := s.createNodeAndDoNotNotifyHost(nodeType)

	toInt := func(i int) *int {
		return &i
	}

	s.rootNodeId = toInt(nodeId)

	yogaNode := s.yogaNodes[nodeId]

	// Ensure proper height/width
	yogaNode.SetWidth(s.deviceScreenSize.Width)
	yogaNode.SetHeight(s.deviceScreenSize.Height)

	return nodeId
}
