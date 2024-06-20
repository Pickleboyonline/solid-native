//
//  SNModule.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore

protocol SolidNativeModule {
    static var name: String {get}
    init()
    func getJSValueRepresentation() -> JSValue
}

