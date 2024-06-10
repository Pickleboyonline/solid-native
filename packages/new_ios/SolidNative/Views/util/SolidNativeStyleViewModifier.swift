//
//  SolidNativeStyleViewModifier.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/10/24.
//

import Foundation
import SwiftUI

// TODO: Handle things like background color and boarder
struct SolidNativeStyleViewModifier: ViewModifier {
    var props: SolidNativeProps
    
    func body(content: Content) -> some View {
        if let styles = props.getPropAsJSValue(name: "style"),
           styles.isObject,
           let color = styles.objectForKeyedSubscript("backgroundColor"),
           color.isString,
           let color = color.toString() {
            content.background(Color(hex: color))
        } else {
            content
        }
    }
}

extension View {
    func applySolidNativeStyles(props: SolidNativeProps) -> some View {
        modifier(SolidNativeStyleViewModifier(props: props))
    }
}
