//
//  SNButtonView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import SwiftUI


struct SNButton: SolidNativeView {
    static var name: String = "sn_button"
    
    var props: SolidNativeProps
    
    var children: SolidNativeChildren
    
    func onPress() {
        print("Hello World!")
    }
    
    var body: some View {
        let title = props.getString("title")
        Button(title) {
            onPress()
        }
    }
}
