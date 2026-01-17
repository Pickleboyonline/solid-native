//
//  HostReceiver.swift
//  SNSwiftUI
//
//  Component registry and delegate implementation.
//  Receives callbacks from Rust core and manages SwiftUI view wrappers.
//

import SNCore
import SwiftUI

/// Main receiver for callbacks from the Rust core.
/// Manages view registration and forwards updates to SwiftUI.
public class HostReceiver: HostDelegate, @unchecked Sendable {
    /// Registry of view types by name
    let viewTypeRegistry: [String: any SolidNativeView.Type] = [
        SNView.name: SNView.self,
        SNText.name: SNText.self,
        SNImage.name: SNImage.self,
    ]

    /// Registry of created view wrappers by node ID
    var viewWrapperRegistry: [String: SolidNativeViewWrapper] = [:]

    /// The root node ID
    var rootNodeId: String?

    /// The SolidNative core instance
    public var core: SolidNativeCore?

    public init() {}

    // MARK: - HostDelegate Implementation

    public func onNodeCreated(nodeId: String, nodeType: String) {
        guard let viewType = viewTypeRegistry[nodeType] else {
            print("Warning: Unknown view type '\(nodeType)'")
            return
        }
        viewWrapperRegistry[nodeId] = SolidNativeViewWrapper(
            id: nodeId,
            viewType: viewType,
            hostReceiver: self
        )

        // Set the first created node as root if not set
        if rootNodeId == nil {
            rootNodeId = nodeId
        }
    }

    public func onNodeRemoved(nodeId: String) {
        viewWrapperRegistry.removeValue(forKey: nodeId)
    }

    public func onPropUpdated(nodeId: String, key: String, value: JsValue) {
        viewWrapperRegistry[nodeId]?.props[key] = value
    }

    public func onChildrenChange(nodeId: String, nodeIds: [String]) {
        viewWrapperRegistry[nodeId]?.children = nodeIds
    }

    public func onTextDescriptorsChange(nodeId: String, descriptors: [TextDescriptor]) {
        viewWrapperRegistry[nodeId]?.textDescriptors = descriptors
    }

    public func onUpdateRevisionCount(nodeId: String) {
        viewWrapperRegistry[nodeId]?.updateRevisionCount()
    }

    public func isTextElementByNodeId(nodeId: String) -> Bool {
        return viewWrapperRegistry[nodeId]?.solidNativeViewType.isTextElement ?? false
    }

    public func isTextElementByNodeType(nodeType: String) -> Bool {
        return viewTypeRegistry[nodeType]?.isTextElement ?? false
    }

    // MARK: - Public API

    /// Gets the root view wrapper
    public func getRootNode() -> SolidNativeViewWrapper? {
        guard let rootId = rootNodeId else { return nil }
        return viewWrapperRegistry[rootId]
    }

    /// Initializes the Rust core with this receiver as the delegate
    public func initializeCore() throws {
        core = try SolidNativeCore(delegate: self)
    }

    /// Creates a root node
    public func createRoot(tag: String) -> String {
        guard let core = core else {
            fatalError("Core not initialized. Call initializeCore() first.")
        }
        let nodeId = core.createRoot(tag: tag)
        rootNodeId = nodeId
        return nodeId
    }

    /// Evaluates JavaScript code
    public func evalJs(_ code: String) throws -> String {
        guard let core = core else {
            throw NSError(domain: "SNSwiftUI", code: 1, userInfo: [
                NSLocalizedDescriptionKey: "Core not initialized"
            ])
        }
        return try core.evalJs(code: code)
    }

    /// Evaluates JavaScript module code
    public func evalModule(_ code: String, moduleName: String) throws -> String {
        guard let core = core else {
            throw NSError(domain: "SNSwiftUI", code: 1, userInfo: [
                NSLocalizedDescriptionKey: "Core not initialized"
            ])
        }
        return try core.evalModule(code: code, moduleName: moduleName)
    }
}
