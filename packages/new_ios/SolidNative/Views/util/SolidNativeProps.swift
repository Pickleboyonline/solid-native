//
//  SnViewProps.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore
import SwiftUI



/// Contains props and helpful conversions
/// TODO: UPDATE ME!
struct SolidNativeProps {
    private var props: [String: JSValue] = [:]
    
    mutating func setProp(_ name: String, _ value: JSValue) {
        self.props[name] = value
    }
    
    mutating func removeProp(_ name: String) {
        self.props.removeValue(forKey: name)
    }
    
    func getString(_ name: String, _ `default`: String = "") -> String {
        self.props[name]?.toString() ?? `default`
    }
    
    func getCGFloat(_ name: String, _ `default`: CGFloat = 0) -> CGFloat {
        if let value = props[name],
           value.isNumber,
           let num = value.toNumber() {
            return CGFloat(truncating: num)
        }
        return CGFloat()
    }
    
    func getJSValue(_ name: String) -> JSValue? {
        self.props[name]
    }

}
