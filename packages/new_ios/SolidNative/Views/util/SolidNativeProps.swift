//
//  SnViewProps.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import SwiftUI
import JavaScriptCore

class SolidNativeProps: ObservableObject {
    @Published var values: [String:JSValue?] = [:];
    // TODO: Type this!
    @Published var children: [SolidNativeView] = [];
    
    func getProp<T>(name: String, `default`: T) -> T {
        if let prop = values[name] as? T {
            return prop
        }
        return `default`
    }
    
    func getString(name: String, `default`: String = "") -> String {
        if let prop = (values[name] ?? nil) {
            return prop.toString()
        }
        return `default`
    }
    
    func getNumber(name: String, `default`: NSNumber = 0) -> NSNumber {
        if let prop = (values[name] ?? nil) {
            return prop.toNumber()
        }
        return `default`
    }
    
    
    func getChildren() -> [SolidNativeView] {
        children
    }
    
    func getPropAsJSValue(name: String) -> JSValue? {
        values[name] ?? nil
    }
    
    func callCallbackWithArgs(name: String, args: [Any]) {
        if let callback = getPropAsJSValue(name: "onPress") {
            callback.call(withArguments: args)
        }
    }
}
