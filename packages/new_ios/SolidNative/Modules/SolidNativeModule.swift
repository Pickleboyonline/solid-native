//
//  SNModule.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore

class SolidNativeModule {
    let jsContext = SolidNativeCore.shared.jsContext
    let id = UUID()
    class var name: String {
        "SolidNativeModule"
    }
    
    required init() {
    }
    
    func getJSValueRepresentation() -> JSValue {
        return JSValue(undefinedIn: jsContext)
    }
}
