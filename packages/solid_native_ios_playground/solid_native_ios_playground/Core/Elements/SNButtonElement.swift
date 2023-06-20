//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNButtonElement: SolidNativeTextElement {

    struct SNButton: View {
        @ObservedObject var props: SolidNativeProps
        
        var body: some View {
            let title = props.getProp(name: "title", default: "")
            let onPress = props.getProp(name: "onPress", default: {})
            Button(title) {
                onPress()
            }
        }
    }
    

    override func render() -> AnyView {
        AnyView(SNButton(props: self.props))
    }
    
}
