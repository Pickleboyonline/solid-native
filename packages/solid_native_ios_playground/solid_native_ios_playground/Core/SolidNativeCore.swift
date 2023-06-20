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
    
    private var elementRegistry: [ElementName: AnySolidNativeElement.Type] = [:]
    
    // Only thing it needs is the View type, the View Name, whether or not it a text view
    private func registerElement(_ name: String, elementType: AnySolidNativeElement.Type) {
        elementRegistry[name] = elementType.self
    }
    
    func registerElements() {
        registerElement("SNVStackElement", elementType: SNVStackElement.self)
        registerElement("SNButtonElement", elementType: SNButtonElement.self)
        registerElement("SNTextElement", elementType: SNTextElement.self)
    }
    
    func createElement(name: String) -> AnySolidNativeElement? {
        if let elementType = elementRegistry[name] {
            return elementType.init()
        }
        assertionFailure("\(name) is not in element registry!")
        return nil
    }
}
