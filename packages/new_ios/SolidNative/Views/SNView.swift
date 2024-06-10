//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI
import YogaSwiftUI

class SNView: SolidNativeView {
    
    class override var name: String {
        "sn_view"
    }

    struct _SNView: View {
        @ObservedObject var props: SolidNativeProps
        
        var body: some View {
            let flexStyles = props.getPropAsJSValue(name: "style")
            let direction = YogaTypeMarshaller.convertToYGFlexDirection(flexStyles?.objectForKeyedSubscript("flexDirection")) ?? .column
            let justifyContent = YogaTypeMarshaller.convertToYGJustify(_:)(flexStyles?.objectForKeyedSubscript("justifyContent")) ?? .flexStart
            let alignItems = YogaTypeMarshaller.convertToYGAlign(flexStyles?.objectForKeyedSubscript("alignItems")) ?? .stretch
            let alignContent = YogaTypeMarshaller.convertToYGAlign(flexStyles?.objectForKeyedSubscript("alignContent")) ?? .flexStart
            let wrap = YogaTypeMarshaller.convertToYGWrap(_:)(flexStyles?.objectForKeyedSubscript("wrap")) ?? .noWrap
            
            
            SolidNativeViewStyler.applyLayoutStyles(props: props) {
                Flex(
                    direction: direction,
                    justifyContent: justifyContent,
                    alignItems: alignItems,
                    alignContent: alignContent,
                    wrap: wrap
                ) {
                    let children = props.getChildren()
                    ForEach(children, id: \.id) { child in
                        child.render()
                    }
                }
            }.ignoresSafeArea()
        }
    }
    
    override func render() -> AnyView {
        return AnyView(_SNView(props: self.props))
    }
    
}
