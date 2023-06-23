//
//  SolidNativeCore.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore

@objc protocol SolidNativeCoreJSExport: JSExport {
    func createElementByName(_ name: String) -> AnySolidNativeElement?
    func getRootElement() -> AnySolidNativeElement
    func createTextElement() -> AnySolidNativeElement?
}

typealias ElementName = String

let ROOT_ELEMENT_CLASS = SNVStackElement.self

// TODO: Make into singleton
@objc public class SolidNativeCore: NSObject, SolidNativeCoreJSExport {

    private var elementRegistry: [ElementName: AnySolidNativeElement.Type] = [ROOT_ELEMENT_CLASS.name : ROOT_ELEMENT_CLASS.self]
    
    // Only thing it needs is the View type, the View Name, whether or not it a text view
    private func registerElement(_ elementType: AnySolidNativeElement.Type) {
        elementRegistry[elementType.name] = elementType.self
    }
    
    func registerElements() {
        registerElement(SNVStackElement.self)
        registerElement(SNButtonElement.self)
        registerElement(SNTextElement.self)
    }
    
    private var _rootElement: AnySolidNativeElement?
    
    private func createRootElement(name: String) -> AnySolidNativeElement {
        let newElement = elementRegistry[name]!.init()
        return newElement
    }
    
    var rootElement: AnySolidNativeElement {
        if let element = _rootElement {
            return element
        }
        let newRootElement = createRootElement(name: ROOT_ELEMENT_CLASS.name)
        _rootElement = newRootElement
        return newRootElement
    }
    
    func createElementByName(_ name: String) -> AnySolidNativeElement? {
        if let elementType = elementRegistry[name] {
            let newElement = elementType.init()
            // createdElementsById[newElement.id.uuidString] = newElement
            return newElement
        }
        assertionFailure("\(name) is not in element registry!")
        return nil
    }
    
    func getRootElement() -> AnySolidNativeElement {
        self.rootElement
    }
    
    func createTextElement() -> AnySolidNativeElement? {
        createElementByName(SNTextElement.name)
    }
    
    func registerCoreInJSContext() {
        SharedJSConext.sharedContext.setObject(self, forKeyedSubscript: "_SolidNativeCore" as NSString)
    }
}
