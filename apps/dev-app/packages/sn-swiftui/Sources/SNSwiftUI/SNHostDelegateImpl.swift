//
//  SNHostDelegateImpl.swift
//  SNSwiftUI
//
//  Implements the HostDelegate protocol from SNCore
//  This bridges the Rust renderer to SwiftUI
//

import Foundation
// Note: Import SNCore once uniffi exports are available
// import SNCore

/// Implementation of HostDelegate that manages a SwiftUI view tree
public class SNHostDelegateImpl {
    private let viewTree: ViewTree

    public init(viewTree: ViewTree) {
        self.viewTree = viewTree
    }

    // MARK: - HostDelegate Methods
    // These methods will conform to the SNCore.HostDelegate protocol once it's exported via uniffi

    /// Called when a node is created by JS or the renderer
    public func onNodeCreated(nodeId: String, nodeType: String) {
        let node = ViewNode(id: nodeId, type: nodeType)
        viewTree.addNode(node)

        print("[SNHostDelegate] Node created: \(nodeId) of type: \(nodeType)")
    }

    /// Called when a node is removed from the tree
    public func onNodeRemoved(nodeId: String) {
        viewTree.removeNode(nodeId)
        print("[SNHostDelegate] Node removed: \(nodeId)")
    }

    /// Called when children of a node change
    public func onChildrenChange(nodeId: String, nodeIds: [String]) {
        guard let parent = viewTree.getNode(nodeId) else {
            print("[SNHostDelegate] Warning: Parent node \(nodeId) not found for children change")
            return
        }

        // Update children on main thread since it affects UI
        DispatchQueue.main.async { [weak self] in
            guard let self = self else { return }

            // Get the new children nodes
            let newChildren = nodeIds.compactMap { self.viewTree.getNode($0) }

            // Update parent's children
            parent.children = newChildren

            // Update parent references
            for child in newChildren {
                child.parent = parent
            }

            print("[SNHostDelegate] Children changed for \(nodeId): \(nodeIds.count) children")
        }
    }

    /// Called when it's time to update the UI
    /// This signals that a batch of changes is complete and the UI should re-render
    public func onUpdateRevisionCount(nodeId: String) {
        // Trigger a SwiftUI update by modifying the node
        if let node = viewTree.getNode(nodeId) {
            DispatchQueue.main.async {
                node.objectWillChange.send()
            }
        }
        print("[SNHostDelegate] Update revision for node: \(nodeId)")
    }

    /// Check if a specific node ID represents a text element
    public func isTextElementByNodeId(nodeId: String) -> Bool {
        guard let node = viewTree.getNode(nodeId) else {
            return false
        }
        return node.isTextNode
    }

    /// Check if a node type represents a text element
    public func isTextElementByNodeType(nodeType: String) -> Bool {
        return nodeType.lowercased() == "text"
    }
}

// MARK: - Future uniffi Protocol Conformance
// Once SNCore exports the HostDelegate protocol via uniffi, add:
// extension SNHostDelegateImpl: HostDelegate { }
