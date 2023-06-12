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
    
    func registerElement() {
        
    }
    
    func createElement(name: String) -> SolidNativeElement {
        SolidNativeElement()
    }
}
