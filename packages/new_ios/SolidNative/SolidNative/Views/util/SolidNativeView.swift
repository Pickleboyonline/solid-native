//
//  SolidNativeView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import SwiftUI

/// Temp protocol for eventual implimention
/// Prob would have to be a class, because we need to get the definitions of them all
/// Need to support "refs" basically js representation of class
/// Callbacks are just refs
protocol SolidNativeView: View where Body: View  {
    /// Needs to be in lower snake case
    static var name: String {get}
    static var isTextElement: Bool {get}
    var props: SolidNativeProps {get}
    var children: SolidNativeChildren {get}
    init(props: SolidNativeProps, children: SolidNativeChildren)
}

extension SolidNativeView {
    static var isTextElement: Bool {false}
}



