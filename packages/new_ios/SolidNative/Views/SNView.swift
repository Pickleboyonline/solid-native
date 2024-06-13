//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI
import YogaSwiftUI

struct SNView: SolidNativeView {
    static var name: String { "sn_view" }
    
    var props: SolidNativeProps
    
    var children: SolidNativeChildren
    
    var body: some View {
        ForEach(children, id: \.id) { child in
            child.render()
        }
    }
}
