//
//  SNSwiftUITests.swift
//  SNSwiftUI
//
//  Tests for the SNSwiftUI package
//

import XCTest
@testable import SNSwiftUI

final class SNSwiftUITests: XCTestCase {

    func testViewNodeCreation() {
        let node = ViewNode(id: "test-1", type: "div")
        XCTAssertEqual(node.id, "test-1")
        XCTAssertEqual(node.type, "div")
        XCTAssertFalse(node.isTextNode)
        XCTAssertTrue(node.children.isEmpty)
    }

    func testTextNodeCreation() {
        let node = ViewNode(id: "test-text", type: "text")
        node.updateText("Hello, World!")

        XCTAssertTrue(node.isTextNode)
        XCTAssertEqual(node.textContent, "Hello, World!")
    }

    func testViewTreeManagement() {
        let tree = ViewTree()
        let node1 = ViewNode(id: "node-1", type: "div")
        let node2 = ViewNode(id: "node-2", type: "span")

        tree.addNode(node1)
        tree.addNode(node2)

        XCTAssertNotNil(tree.getNode("node-1"))
        XCTAssertNotNil(tree.getNode("node-2"))

        tree.removeNode("node-1")
        XCTAssertNil(tree.getNode("node-1"))
    }

    func testNodeHierarchy() {
        let parent = ViewNode(id: "parent", type: "div")
        let child1 = ViewNode(id: "child1", type: "span")
        let child2 = ViewNode(id: "child2", type: "text")

        parent.addChild(child1)
        parent.addChild(child2)

        XCTAssertEqual(parent.children.count, 2)
        XCTAssertEqual(child1.parent?.id, "parent")
        XCTAssertEqual(child2.parent?.id, "parent")

        parent.removeChild(child1)
        XCTAssertEqual(parent.children.count, 1)
        XCTAssertNil(child1.parent)
    }

    func testHostDelegateCallbacks() {
        let tree = ViewTree()
        let delegate = SNHostDelegateImpl(viewTree: tree)

        // Test node creation
        delegate.onNodeCreated(nodeId: "test-node", nodeType: "div")
        XCTAssertNotNil(tree.getNode("test-node"))

        // Test node removal
        delegate.onNodeRemoved(nodeId: "test-node")
        XCTAssertNil(tree.getNode("test-node"))
    }

    func testChildrenChange() {
        let tree = ViewTree()
        let delegate = SNHostDelegateImpl(viewTree: tree)

        // Create parent and children
        delegate.onNodeCreated(nodeId: "parent", nodeType: "div")
        delegate.onNodeCreated(nodeId: "child1", nodeType: "span")
        delegate.onNodeCreated(nodeId: "child2", nodeType: "span")

        // Update children
        delegate.onChildrenChange(nodeId: "parent", nodeIds: ["child1", "child2"])

        let parent = tree.getNode("parent")
        XCTAssertEqual(parent?.children.count, 2)
    }

    func testManagerTreeBuilding() {
        let manager = SNSwiftUIManager()

        // Create a simple tree structure
        let root = manager.createRoot(type: "vstack")
        let button = manager.createElement(type: "button")
        let text = manager.createTextNode(text: "Click me")

        manager.insertNode(parent: root, child: button)
        manager.insertNode(parent: button, child: text)

        XCTAssertEqual(root.children.count, 1)
        XCTAssertEqual(button.children.count, 1)
        XCTAssertEqual(text.textContent, "Click me")

        // Test properties
        manager.setProperty(node: button, key: "backgroundColor", value: "blue")
        XCTAssertEqual(button.properties["backgroundColor"], "blue")
    }

    func testViewNodeProperties() {
        let node = ViewNode(id: "styled-node", type: "div")
        node.updateProperty(key: "padding", value: "16")
        node.updateProperty(key: "backgroundColor", value: "red")
        node.updateProperty(key: "cornerRadius", value: "8")

        XCTAssertEqual(node.properties["padding"], "16")
        XCTAssertEqual(node.properties["backgroundColor"], "red")
        XCTAssertEqual(node.properties["cornerRadius"], "8")
    }

    func testIsTextElementDetection() {
        let tree = ViewTree()
        let delegate = SNHostDelegateImpl(viewTree: tree)

        XCTAssertTrue(delegate.isTextElementByNodeType(nodeType: "text"))
        XCTAssertTrue(delegate.isTextElementByNodeType(nodeType: "TEXT"))
        XCTAssertFalse(delegate.isTextElementByNodeType(nodeType: "div"))

        delegate.onNodeCreated(nodeId: "text-node", nodeType: "text")
        XCTAssertTrue(delegate.isTextElementByNodeId(nodeId: "text-node"))

        delegate.onNodeCreated(nodeId: "div-node", nodeType: "div")
        XCTAssertFalse(delegate.isTextElementByNodeId(nodeId: "div-node"))
    }
}
