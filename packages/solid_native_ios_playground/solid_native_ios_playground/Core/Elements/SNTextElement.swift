//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNTextElement: AnySolidNativeElement {
    struct SNTextView: View {
        
        @ObservedObject var props: SolidNativeProps
        
        var body: some View {
            Text(props.text)
        }
    }
    
    override func render() -> AnyView
    {
        AnyView(SNTextView(props: self.props))
    }
    
}
