//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNView: SolidNativeView {
    static var name: String {
        "sn_view"
    }
    
    func render() -> some View {
        ZStack(alignment: .topLeading) {
            ForEach(children, id: \.description) { nodeId in
                self.wrapper.hostReceiver.viewWrapperRegistry[nodeId]!.render()
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
