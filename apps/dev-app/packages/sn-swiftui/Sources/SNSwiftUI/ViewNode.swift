//
//  ViewNode.swift
//  SNSwiftUI
//
//  Represents a node in the virtual tree that maps to SwiftUI views
//

import Foundation

/// Represents a node in the view tree
public class ViewNode: Identifiable, ObservableObject {
    public let id: String
    public let type: String

    @Published public var properties: [String: String] = [:]
    @Published public var children: [ViewNode] = []
    @Published public var textContent: String?

    public weak var parent: ViewNode?

    public init(id: String, type: String) {
        self.id = id
        self.type = type
    }

    /// Check if this is a text node
    public var isTextNode: Bool {
        return type == "text"
    }

    /// Add a child node at a specific position
    public func addChild(_ child: ViewNode, before anchor: ViewNode? = nil) {
        child.parent = self

        if let anchor = anchor, let index = children.firstIndex(where: { $0.id == anchor.id }) {
            children.insert(child, at: index)
        } else {
            children.append(child)
        }
    }

    /// Remove a child node
    public func removeChild(_ child: ViewNode) {
        children.removeAll { $0.id == child.id }
        child.parent = nil
    }

    /// Update a property
    public func updateProperty(key: String, value: String) {
        properties[key] = value
    }

    /// Update text content (for text nodes)
    public func updateText(_ text: String) {
        textContent = text
    }
}

extension ViewNode: CustomStringConvertible {
    public var description: String {
        if isTextNode {
            return "TextNode(\(id)): '\(textContent ?? "")'"
        } else {
            return "ViewNode(\(id)): \(type) with \(children.count) children"
        }
    }
}
