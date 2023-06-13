//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNViewElement: SolidNativeTextElement {
    struct SNView: View {
        @ObservedObject var props: SolidNativeProps
        
        var body: some View {
            Text(props.text)
        }
    }
    
    func render() -> some View {
        SNView(props: self.props)
    }
    
}
