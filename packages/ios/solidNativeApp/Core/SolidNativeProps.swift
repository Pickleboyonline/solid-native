//
//  SolidNativeProps.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore
import SwiftUI

// Not meant to be exteneded
class SolidNativeProps: ObservableObject {
    @Published var values: [String:JSValue?] = [:];
    @Published var children: [AnySolidNativeElement] = [];
    
    // TODO: Figure out how to deal with events and refs
    
//    func getProp<T>(name: String, `default`: T) -> T {
//        if let prop = values[name] as? T {
//            return prop
//        }
//        return `default`
//    }
//
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
    
    /**
     Usefull for callbacks
     */
    func getPropAsJSValue(name: String) -> JSValue? {
        values[name] ?? nil
    }
    
    func getChildren() -> [AnySolidNativeElement] {
        children
    }
    
    func getChildrenAsView() -> some View {
        ForEach(getChildren(), id: \.id) { child in
            child.render()
        }
    }
}
