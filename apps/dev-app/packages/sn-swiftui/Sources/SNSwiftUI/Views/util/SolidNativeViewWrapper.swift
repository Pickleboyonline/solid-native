//
//  SolidNativeViewWrapper.swift
//  SNSwiftUI
//
//  Manages view state and provides a stable reference for SwiftUI.
//  Simplified version - no layout metrics, Yoga-SwiftUI handles layout directly.
//

import Foundation
import SNCore
import SwiftUI

typealias SolidNativeProps = [String: JsValue]
typealias SolidNativeChildren = [String]

/// Wrapper class that provides a stable reference for SwiftUI
/// and manages view updates from the Rust core.
public class SolidNativeViewWrapper: ObservableObject {
    @MainActor
    @Published
    var revision: UInt = 0

    /// Unique identifier for SwiftUI iterators
    let id: String

    /// Props received from JavaScript
    var props: SolidNativeProps = [:]

    /// View type for this wrapper
    var solidNativeViewType: any SolidNativeView.Type

    /// Child node IDs
    var children: SolidNativeChildren = []

    /// Text descriptors for text elements
    var textDescriptors: [TextDescriptor]?

    /// The instantiated view
    var solidNativeView: (any SolidNativeView)?

    /// Reference to the host receiver for accessing other wrappers
    weak var hostReceiver: HostReceiver?

    init(id: String, viewType: any SolidNativeView.Type, hostReceiver: HostReceiver) {
        self.id = id
        self.solidNativeViewType = viewType
        self.hostReceiver = hostReceiver
        self.solidNativeView = solidNativeViewType.init(wrapper: self)
    }

    /// Notify SwiftUI of changes
    func updateRevisionCount() {
        Task { @MainActor in
            revision += 1
        }
    }

    @ViewBuilder func render() -> some View {
        _SolidNativeViewWrapperView(wrapper: self, view: solidNativeView!)
    }
}

/// Internal SwiftUI view that observes the wrapper for updates
private struct _SolidNativeViewWrapperView: View {
    @ObservedObject var wrapper: SolidNativeViewWrapper
    let view: any SolidNativeView

    func style(_ view: some View) -> some View {
        let props = wrapper.props
        var backgroundColor = Color.clear
        var foregroundColor: Color? = nil
        var opacity = 1.0

        if let style = props["style"] {
            if case .object(let styleProps) = style {
                if let bg = styleProps["backgroundColor"], case .string(let bgHex) = bg {
                    backgroundColor = Color(hex: bgHex)
                }
                if let fg = styleProps["color"], case .string(let fgHex) = fg {
                    foregroundColor = Color(hex: fgHex)
                }
                if let o = styleProps["opacity"], case .number(let opacityVal) = o {
                    opacity = opacityVal
                }
            }
        }

        var result = view
            .background(backgroundColor)
            .opacity(opacity)

        if let fgColor = foregroundColor {
            return AnyView(result.foregroundColor(fgColor))
        }
        return AnyView(result)
    }

    var body: some View {
        AnyView(style(view.render()))
    }
}
