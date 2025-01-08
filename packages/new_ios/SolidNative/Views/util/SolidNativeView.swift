//
//  SolidNativeView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import SNLib
import SwiftUI

/// Temp protocol for eventual implimention
/// Prob would have to be a class, because we need to get the definitions of them all
/// Need to support "refs" basically js representation of class
/// Callbacks are just refs
protocol SolidNativeViewProtocol {
    /// Needs to be in lower snake case
    static var name: String { get }
    static var isTextElement: Bool { get }
    /// Used for things like text and textinput
    static var doesRequireMeasuring: Bool { get }

    func measureNode(_ nodeId: String) -> SNRendererSize

    init(wrapper: SolidNativeViewWrapper)
    var wrapper: SolidNativeViewWrapper { get }

    associatedtype V: View
    
    @ViewBuilder
    func render() -> V
}

extension SolidNativeViewProtocol {
    static var isTextElement: Bool { false }
    static var doesRequireMeasuring: Bool { false }
    
    func measureNode(_ nodeId: String) -> SNRendererSize {
        .init(0, height: 0)!
    }

    var props: SolidNativeProps {
        wrapper.props
    }

    var children: SolidNativeChildren {
        wrapper.children
    }
}

class BaseView {
    var wrapper: SolidNativeViewWrapper
    
    required init(wrapper: SolidNativeViewWrapper) {
        self.wrapper = wrapper
    }
}

typealias SolidNativeView = BaseView & SolidNativeViewProtocol
