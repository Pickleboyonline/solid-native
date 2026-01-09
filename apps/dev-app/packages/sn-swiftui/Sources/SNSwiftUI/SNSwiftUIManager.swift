//
//  SNSwiftUIManager.swift
//  SNSwiftUI
//
//  High-level manager that coordinates the view tree, delegate, and renderer
//

import Foundation
import SwiftUI
// Note: Import SNCore once uniffi exports are available
// import SNCore

/// Main manager that coordinates the SwiftUI rendering system
public class SNSwiftUIManager: ObservableObject {
    public let viewTree: ViewTree
    public let hostDelegate: SNHostDelegateImpl

    // Reference to SNCore instance (once uniffi exports are available)
    // private var core: SolidNativeCore?

    public init() {
        self.viewTree = ViewTree()
        self.hostDelegate = SNHostDelegateImpl(viewTree: viewTree)
    }

    // MARK: - Initialization with SNCore
    // Once SNCore is properly exported via uniffi, this will create the integration

    /*
    public func initializeWithCore() throws {
        // Create the core with our delegate
        self.core = try SolidNativeCore(delegate: self.hostDelegate)
    }

    public func evaluateJavaScript(_ code: String) throws -> String {
        guard let core = core else {
            throw SNSwiftUIError.coreNotInitialized
        }
        return try core.evalJs(code: code)
    }

    public func evaluateModule(_ code: String, name: String) throws -> String {
        guard let core = core else {
            throw SNSwiftUIError.coreNotInitialized
        }
        return try core.evalModule(code: code, moduleName: name)
    }
    */

    // MARK: - Manual Tree Building (for testing without core)

    /// Creates a root view node
    public func createRoot(type: String = "view") -> ViewNode {
        let node = ViewNode(id: "root", type: type)
        viewTree.setRoot(node)
        hostDelegate.onNodeCreated(nodeId: node.id, nodeType: node.type)
        return node
    }

    /// Creates a new node
    public func createElement(type: String) -> ViewNode {
        let node = ViewNode(id: UUID().uuidString, type: type)
        viewTree.addNode(node)
        hostDelegate.onNodeCreated(nodeId: node.id, nodeType: node.type)
        return node
    }

    /// Creates a text node
    public func createTextNode(text: String) -> ViewNode {
        let node = ViewNode(id: UUID().uuidString, type: "text")
        node.textContent = text
        viewTree.addNode(node)
        hostDelegate.onNodeCreated(nodeId: node.id, nodeType: node.type)
        return node
    }

    /// Inserts a child into a parent node
    public func insertNode(parent: ViewNode, child: ViewNode, before anchor: ViewNode? = nil) {
        parent.addChild(child, before: anchor)
        hostDelegate.onChildrenChange(nodeId: parent.id, nodeIds: parent.children.map { $0.id })
        hostDelegate.onUpdateRevisionCount(nodeId: parent.id)
    }

    /// Removes a child from a parent node
    public func removeNode(parent: ViewNode, child: ViewNode) {
        parent.removeChild(child)
        hostDelegate.onNodeRemoved(nodeId: child.id)
        hostDelegate.onChildrenChange(nodeId: parent.id, nodeIds: parent.children.map { $0.id })
        hostDelegate.onUpdateRevisionCount(nodeId: parent.id)
    }

    /// Sets a property on a node
    public func setProperty(node: ViewNode, key: String, value: String) {
        node.updateProperty(key: key, value: value)
        hostDelegate.onUpdateRevisionCount(nodeId: node.id)
    }

    /// Clears the entire tree
    public func clear() {
        viewTree.clear()
    }

    /// Prints the current tree structure for debugging
    public func debugPrintTree() {
        viewTree.printTree()
    }
}

// MARK: - Errors

public enum SNSwiftUIError: Error, LocalizedError {
    case coreNotInitialized

    public var errorDescription: String? {
        switch self {
        case .coreNotInitialized:
            return "SNCore has not been initialized. Call initializeWithCore() first."
        }
    }
}
