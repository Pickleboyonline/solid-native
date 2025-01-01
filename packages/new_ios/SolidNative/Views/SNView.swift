//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNView: SolidNativeView {
    required init(wrapper: SolidNativeViewWrapper) {
        self.wrapper = wrapper
    }

    var wrapper: SolidNativeViewWrapper
    static var name: String {
        "sn_view"
    }

    @ViewBuilder
    func render() -> some View {
        ZStack(alignment: .topLeading) {
            ForEach(children, id: \.id) { child in
                child.render()
            }
        }
        // TODO: Place gestures:
        /*
         .onTapGesture {
           self.eventEmitter.dispatchEvent("tap")
         }
         .onLongPressGesture {
           self.eventEmitter.dispatchEvent("longPress")
         }
         */
    }
}
