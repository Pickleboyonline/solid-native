//
//  SolidNativeViewWrapper.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import Snmobile
import SwiftUI

typealias SolidNativeProps = [String: SNSnmobileJSValue]
typealias SolidNativeChildren = SNSnmobileStringArray

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
    var solidNativeView: (any SolidNativeView)?
    
    init(viewType: any SolidNativeView.Type) {
        self.solidNativeViewType = viewType
    }
    
    /// Notify SwiftUI of changes
    func updateRevisionCount() {
        revision += 1
    }
    
    func getSolidNativeView() -> any SolidNativeView {
        if let solidNativeView {
            return solidNativeView
        } else {
            let v = solidNativeViewType.init(wrapper: self)
            solidNativeView = v
            return v
        }
    }
    
    @ViewBuilder func render() -> some View {
        _SolidNativeViewWrapper(wrapper: self, view: getSolidNativeView())
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

    func style(_ view: some View) -> some View {
        let props = wrapper.props
        var backgroundColor = Color.clear
        
        var foregroundColor = Color.white
        var opacity = 1.0
        
        if let style = props["style"],
           style.isObject()
        {
            if let bg = style.getForKey("backgroundColor"),
               bg.isString()
            {
                backgroundColor = Color(hex: bg.getString())
            }
            
            if let fg = style.getForKey("color"),
               fg.isString()
            {
                foregroundColor = Color(hex: fg.getString())
            }
            
            if let o = style.getForKey("opacity"),
               o.isNumber()
            {
                opacity = o.getNumber()
            }
        }
    
        return view.background(backgroundColor)
            .foregroundColor(foregroundColor)
            .overlay(Border())
            .opacity(opacity)
    }
    
    var body: some View {
        // TODO: Is there someway to remove the `AnyView` here?
        AnyView(style(layout(view.render()))
            .offset(
                x: CGFloat(wrapper.layoutMetrics.x),
                y: CGFloat(wrapper.layoutMetrics.y)
            ).ignoresSafeArea(.all))
    }
    
    func Border() -> some View {
        // TODO: Pull from props
        let width = 0.0
        let color = Color.clear

        return
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
