//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNVStackElement: AnySolidNativeElement {
    
    class override var name: String {
        "sn_v_stack"
    }
    
    var isOn: Bool = false
    
    struct SNVStack: View {
        @ObservedObject var props: SolidNativeProps
        
        // var isOn: Binding<Bool>
        
        var body: some View {
            let children = props.getChildren()
            VStack {
                ForEach(children, id: \.id) { child in
                    child.render()
                }
                // Toggle("", isOn: isOn)
            }
        }
    }
    
    override func render() -> AnyView {
//        let valueBinding = Binding<Bool>(
//          get: {
//              return self.props.getProp(name: "value", default: self.isOn)
//          },
//          set: {
//            if self.props.getProp(name: "value", default: self.isOn) != $0 {
//              self.isOn = $0
//              let callback = self.props.getProp(name: "onChange", default: {(_ newValue: Bool) -> Void in
//
//              })
//              callback($0)
//            }
//          }
//        )
        
        return AnyView(SNVStack(props: self.props))
    }
    
}
