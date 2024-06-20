//
//  SolidNativeViewWrapper.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import SwiftUI


/// Manages Flex Layout. Nodes take in a wrapper. Wrapper takes in view struct def to instanciate
/// View takes in view types. (Managed in render for now)
class SolidNativeViewWrapper: ObservableObject {
    @Published
    var revision: UInt = 0
    
    /// For SwiftUI iterators
    let id = UUID().uuidString
    
    // Props + Children only info needed. Pass that down to
    var props = SolidNativeProps()
    var solidNativeViewType: any SolidNativeView.Type
    var children: SolidNativeChildren = []
    var layoutMetrics = LayoutMetrics()
    
    init(viewType: any SolidNativeView.Type) {
        self.solidNativeViewType = viewType
    }
    
    /// Notify SwiftUI of changes
    func notifyChanges() {
        revision += 1
    }
    
    struct _SolidNativeViewWrapper: View {
        /// Causes the update in swiftUI
        @ObservedObject
        var wrapper: SolidNativeViewWrapper
        let view: any SolidNativeView
        
        func layout(_ view: some View) -> some View {
          let layoutMetrics = wrapper.layoutMetrics
          return view
            .frame(
              width: layoutMetrics.width,
              height: layoutMetrics.height,
              alignment: .topLeading
            )
        }

        func style<InputType: View>(_ view: InputType) -> some View {
          let props = wrapper.props
          // let backgroundColor = props.color("backgroundColor", .clear)
          // let foregroundColor = props.color("color", .white)
          // let opacity = props.double("opacity", 1.0)

          return view
            // .background(backgroundColor)
            // .foregroundColor(foregroundColor)
            // .overlay(Border())
            // .opacity(opacity)
        }
        
        var body: some View {
            EmptyView()
        }
    }
    
    func render() -> some View {
        _SolidNativeViewWrapper(wrapper: self, view: solidNativeViewType.init(props: props, children: children))
    }
}

