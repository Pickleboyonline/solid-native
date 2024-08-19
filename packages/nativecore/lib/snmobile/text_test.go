package snmobile

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_generateTextDescriptor(t *testing.T) {
	type args struct {
		node *NodeContainer
	}
	tests := []struct {
		name  string
		args  args
		want  []TextDescriptor
		want1 *NodeContainer
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				node: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateTextDescriptor(tt.args.node)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTextDescriptor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("generateTextDescriptor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// How to test me:
// go test -timeout 30s -v -run ^TestParentTextDescriptor$ nativecore/lib/snmobile
func TestParentTextDescriptor(t *testing.T) {
	// This setup translates to the following JSX:
	//
	// ```
	// <Text style={{color: "red", fontSize: 12}}>
	//	Hello World!
	// </Text>
	// ```
	//
	// Note that the <Text> is the parent and has no text value, only styles. whereas the
	// "Hello World!" is the child with no styles but has text.
	parentNode := newNodeContainer(true)
	childNode := newNodeContainer(true)

	parentNode.children = []*NodeContainer{&childNode}

	styles := map[string]JSValue{
		"color":    NewJSValue("red"),
		"fontSize": NewJSValue(12),
	}
	text := "Hello World!"

	parentNode.styleMap = styles
	childNode.text = text

	textDescriptor, returnedParent := generateTextDescriptor(&parentNode)

	if returnedParent != &parentNode {
		t.Fatalf("returned parent did not equal parent")
	}

	want := []TextDescriptor{
		{
			Text:   text,
			styles: styles,
		},
	}

	if !reflect.DeepEqual(textDescriptor, want) {
		t.Errorf("generateTextDescriptor() got1 = %v, want %v", textDescriptor, want)
	}

	fmt.Printf("Text Descriptor: %#v \n", textDescriptor)

}

// How to test me:
// go test -timeout 30s -v -run ^TestOverridingStylesTextDescriptor$ nativecore/lib/snmobile
func TestOverridingStylesTextDescriptor(t *testing.T) {
	// This setup translates to the following JSX:
	//
	// ```
	// <Text style={{color: "red", fontSize: 12}}> => parent
	//  <Text style={{color: "blue" }}>Hello{ /* => node3 */} </Text> => node1
	//  {" World!"} => node2
	// </Text>
	// ```
	//
	// Note that the <Text> is the parent and has no text value, only styles. whereas the
	// "Hello World!" is the child with no styles but has text.

	// Make nodes
	parentNode := newNodeContainer(true)
	parentNode.id = "parentNode"

	node1 := newNodeContainer(true)
	node1.id = "node1"

	node2 := newNodeContainer(true)
	node2.id = "node2"

	node3 := newNodeContainer(true)
	node3.id = "node3"

	// Add props

	parentNode.styleMap = map[string]JSValue{
		"color":    NewJSValue("red"),
		"fontSize": NewJSValue(12),
	}

	node1.styleMap = map[string]JSValue{
		"color": NewJSValue("blue"),
	}

	node3.text = "Hello"
	node2.text = " World!"

	// Children

	parentNode.children = []*NodeContainer{&node1, &node2}
	node1.parent = &parentNode
	node2.parent = &parentNode

	node1.children = []*NodeContainer{&node3}
	node3.parent = &node1

	textDescriptor, returnedParent := generateTextDescriptor(&node2)

	if returnedParent != &parentNode {
		t.Fatalf("returned parent did not equal parent")
	}

	want := []TextDescriptor{
		{
			Text: "Hello",
			styles: map[string]JSValue{
				"color":    NewJSValue("blue"),
				"fontSize": NewJSValue(12),
			},
		},
		{
			Text: " World!",
			styles: map[string]JSValue{
				"color":    NewJSValue("red"),
				"fontSize": NewJSValue(12),
			},
		},
	}

	if !reflect.DeepEqual(textDescriptor, want) {
		t.Errorf("generateTextDescriptor(): \ngot1 = %v \nwant = %v", textDescriptor, want)
	}

	fmt.Printf("Text Desriptor: %#v \n", textDescriptor)

}
