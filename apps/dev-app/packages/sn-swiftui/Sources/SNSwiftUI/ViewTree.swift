//
//  ViewTree.swift
//  SNSwiftUI
//
//  Manages the view tree and provides thread-safe access
//

import Foundation
import Combine

/// Manages the view tree structure
public class ViewTree: ObservableObject {
    /// Root node of the tree
    @Published public private(set) var root: ViewNode?

    /// All nodes indexed by ID for quick lookup
    private var nodes: [String: ViewNode] = [:]

    /// Queue for thread-safe access
    private let queue = DispatchQueue(label: "com.solidnative.viewtree", attributes: .concurrent)

    public init() {}

    // MARK: - Node Management

    /// Add or update a node in the tree
    public func addNode(_ node: ViewNode) {
        queue.async(flags: .barrier) { [weak self] in
            self?.nodes[node.id] = node
        }
    }

    /// Get a node by ID
    public func getNode(_ id: String) -> ViewNode? {
        return queue.sync {
            return nodes[id]
        }
    }

    /// Remove a node from the tree
    public func removeNode(_ id: String) {
        queue.async(flags: .barrier) { [weak self] in
            guard let node = self?.nodes[id] else { return }

            // Remove from parent
            if let parent = node.parent {
                DispatchQueue.main.async {
                    parent.removeChild(node)
                }
            }

            // Remove from index
            self?.nodes.removeValue(forKey: id)
        }
    }

    /// Set the root node
    public func setRoot(_ node: ViewNode) {
        DispatchQueue.main.async { [weak self] in
            self?.root = node
        }
        addNode(node)
    }

    /// Clear all nodes
    public func clear() {
        queue.async(flags: .barrier) { [weak self] in
            self?.nodes.removeAll()
        }
        DispatchQueue.main.async { [weak self] in
            self?.root = nil
        }
    }

    /// Get all node IDs
    public func allNodeIds() -> [String] {
        return queue.sync {
            return Array(nodes.keys)
        }
    }

    /// Print tree structure for debugging
    public func printTree() {
        guard let root = root else {
            print("Tree is empty")
            return
        }
        print("Tree structure:")
        printNode(root, indent: 0)
    }

    private func printNode(_ node: ViewNode, indent: Int) {
        let indentation = String(repeating: "  ", count: indent)
        print("\(indentation)\(node)")
        for child in node.children {
            printNode(child, indent: indent + 1)
        }
    }
}
