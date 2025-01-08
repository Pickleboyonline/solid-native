//
//  SNButtonView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import SwiftUI

class SNButton: SolidNativeView {
    static var name: String { "sn_button" }
    
    func onPress() {
        print("Hello World!")
    }
    
    func render() -> some View {
        let title = props["text"]?.getString() ?? ""
        Button(title) {
            self.onPress()
        }
    }
}
