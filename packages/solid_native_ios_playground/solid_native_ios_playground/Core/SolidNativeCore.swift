//
//  SolidNativeCore.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore

protocol SolidNativeCoreJSExport: JSExport {
    
}

typealias ElementName = String

@objc public class SolidNativeCore: NSObject, SolidNativeCoreJSExport {
    
    private var elementRegistry: [ElementName: SolidNativeElement] = [:]
    
    // Only thing it needs is the View type, the View Name, whether or not it a text view
    func registerElement() {
        
    }
    
    func createElement(name: String) -> SolidNativeElement {
        SolidNativeElement()
    }
}
