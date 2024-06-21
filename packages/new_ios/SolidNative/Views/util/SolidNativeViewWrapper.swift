//
//  SolidNativeViewWrapper.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import SwiftUI
import Snmobile

typealias SolidNativeProps = [String: SNSnmobileJSValue]
typealias SolidNativeChildren = SNSnmobileStringArray


extension SNSnmobileStringArray: RandomAccessCollection {
    public typealias Element = SolidNativeViewWrapper
    public typealias Index = Int
    
    public var startIndex: Index {
        return 0
    }
    
    public var endIndex: Index {
        return length()
    }
    
    public subscript(position: Index) -> Element {
        let nodeId = get(position)
        return SharedSolidNativeCore.viewWrapperRegistry[nodeId]!
    }
    
    public func index(after i: Index) -> Index {
        return i + 1
    }
    
    public func index(before i: Index) -> Index {
        return i - 1
    }
}

/// Manages Flex Layout. Nodes take in a wrapper. Wrapper takes in view struct def to instanciate
/// View takes in view types. (Managed in render for now)
public class SolidNativeViewWrapper: ObservableObject {
    @Published
    var revision: UInt = 0
    
    /// For SwiftUI iterators
    let id = UUID().uuidString
    
    // Props + Children only info needed. Pass that down to
    var props: SolidNativeProps = [:]
    var solidNativeViewType: any SolidNativeView.Type
    var children: SolidNativeChildren = SNSnmobileStringArray()
    var layoutMetrics = SNSnmobileLayoutMetrics()
    
    init(viewType: any SolidNativeView.Type) {
        self.solidNativeViewType = viewType
    }
    

    /// Notify SwiftUI of changes
    func updateRevisionCount() {
        revision += 1
    }
    
    
    
    func render() -> some View {
        _SolidNativeViewWrapper(wrapper: self, view: solidNativeViewType.init(props: props, children: children))
    }
}


private struct _SolidNativeViewWrapper: View {
    /// Causes the update in swiftUI
    
    @ObservedObject var wrapper: SolidNativeViewWrapper
    
    let view: any SolidNativeView
    
    func layout(_ view: some View) -> some View {
      let layoutMetrics = wrapper.layoutMetrics
      return view
        .frame(
          width: CGFloat(layoutMetrics.width),
          height: CGFloat(layoutMetrics.height),
          alignment: .topLeading
        )
    }

    func style<InputType: View>(_ view: InputType) -> some View {
      let props = wrapper.props
        var backgroundColor = Color.clear
        
        if let bg = props["backgroundColor"],
           bg.isString() {
            backgroundColor = Color(hex: bg.getString())
        }
    
        let foregroundColor = Color.white
        let opacity = 1.0
      // let backgroundColor = props.color("backgroundColor", .clear)
      // let foregroundColor = props.color("color", .white)
      // let opacity = props.double("opacity", 1.0)

      return view.background(backgroundColor)
         .foregroundColor(foregroundColor)
         .overlay(Border())
         .opacity(opacity)
    }
    
    var body: some View {
        AnyView(
            style(layout(view))
                .offset(
                    x: CGFloat(wrapper.layoutMetrics.x),
                    y: CGFloat(wrapper.layoutMetrics.y)
                )
        ).edgesIgnoringSafeArea(.all)
    }
    
    func Border() -> some View {
      // TODO: Pull from props
      let width = 0.0
      let color = Color.clear

      return AnyView(
        Rectangle()
          .fill(Color.clear)
          .overlay(
            Rectangle()
              .frame(
                width: nil,
                height: width,
                alignment: .top
              )
              .foregroundColor(color),
            alignment: .top
          )
          .overlay(
            Rectangle()
              .frame(
                width: width,
                height: nil,
                alignment: .trailing
              )
              .foregroundColor(color),
            alignment: .trailing
          )
          .overlay(
            Rectangle()
              .frame(
                width: nil,
                height: width,
                alignment: .bottom
              )
              .foregroundColor(color),
            alignment: .bottom
          )
          .overlay(
            Rectangle()
              .frame(
                width: width,
                height: nil,
                alignment: .leading
              )
              .foregroundColor(color),
            alignment: .leading
          )
      )
    }
}


/**
 Full border implementation:
 func Border() -> AnyView {
   let props = descriptor.props
   let width = props.cgFloat("borderWidth", 0.0)
   let color = props.color("borderColor", Color.clear)

   return AnyView(
     Rectangle()
       .fill(Color.clear)
       .overlay(
         Rectangle()
           .frame(
             width: nil,
             height: props.cgFloat("borderTopWidth", width),
             alignment: .top
           )
           .foregroundColor(props.color("borderTopColor", color)),
         alignment: .top
       )
       .overlay(
         Rectangle()
           .frame(
             width: props.cgFloat("borderRightWidth", width),
             height: nil,
             alignment: .trailing
           )
           .foregroundColor(props.color("borderRightColor", color)),
         alignment: .trailing
       )
       .overlay(
         Rectangle()
           .frame(
             width: nil,
             height: props.cgFloat("borderBottomWidth", width),
             alignment: .bottom
           )
           .foregroundColor(props.color("borderBottomColor", color)),
         alignment: .bottom
       )
       .overlay(
         Rectangle()
           .frame(
             width: props.cgFloat("borderLeftWidth", width),
             height: nil,
             alignment: .leading
           )
           .foregroundColor(props.color("borderLeftColor", color)),
         alignment: .leading
       )
   )
 }
 
 */
