//
//  SolidNativeView.swift
//  SNSwiftUI
//
//  Simplified protocol for Solid Native view components.
//  No layout metrics needed - Yoga-SwiftUI handles all layout directly.
//

import Foundation
import SNCore
import SwiftUI

/// Protocol for Solid Native view components
protocol SolidNativeViewProtocol {
    /// Component name in lower snake case (e.g., "sn_view", "sn_text")
    static var name: String { get }

    /// Whether this component is a text element (for text node handling)
    static var isTextElement: Bool { get }

    init(wrapper: SolidNativeViewWrapper)
    var wrapper: SolidNativeViewWrapper { get }

    associatedtype V: View

    @ViewBuilder
    func render() -> V
}

extension SolidNativeViewProtocol {
    static var isTextElement: Bool { false }

    var props: SolidNativeProps {
        wrapper.props
    }

    var children: SolidNativeChildren {
        wrapper.children
    }
}

/// Base class for Solid Native views
class BaseView {
    var wrapper: SolidNativeViewWrapper

    required init(wrapper: SolidNativeViewWrapper) {
        self.wrapper = wrapper
    }
}

/// Combines BaseView class with SolidNativeViewProtocol
typealias SolidNativeView = BaseView & SolidNativeViewProtocol
