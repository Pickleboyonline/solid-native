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
        "v_stack"
    }

    struct SNVStack: View {
        @ObservedObject var props: SolidNativeProps
        
        var body: some View {
            let children = props.getChildren()
            VStack {
                ForEach(children, id: \.id) { child in
                    child.render()
                }
            }
        }
    }
    
    override func render() -> AnyView {
        AnyView(SNVStack(props: self.props))
    }
    
}
