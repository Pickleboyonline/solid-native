//
//  SolidNativeCore.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore

@objc protocol SolidNativeCoreJSExport: JSExport {
//    var createElementByName: @convention(block) (_ name: String) -> String { get }
//
//    var isTextElement: @convention(block) (String) -> Bool { get }
//
//    var createTextElement: @convention(block) () -> String { get }
//
//    var removeElement: @convention(block) (String, String) -> Void { get }
//
//    var insertElement: @convention(block) (String, String, String?) -> Void { get }
//
//    var getParentElementId: @convention(block) (String) -> String? { get }
//
//    var getFirstChildElementId: @convention(block) (String) -> String? { get }
//
//    var getNextSiblingElementId: @convention(block) (String) -> String? { get }
//
//    var setPropertyOnElement: @convention(block) (String, String, Any?) -> Void { get }
//
//    var getRootElement: @convention(block) () -> String { get }
    func createElementByName(_ name: String) -> AnySolidNativeElement?
    func getRootElement() -> AnySolidNativeElement
    func createTextElement() -> AnySolidNativeElement?
}

typealias ElementName = String

let ROOT_ELEMENT_CLASS = SNVStackElement.self

@objc public class SolidNativeCore: NSObject, SolidNativeCoreJSExport {

    private var elementRegistry: [ElementName: AnySolidNativeElement.Type] = [ROOT_ELEMENT_CLASS.name : ROOT_ELEMENT_CLASS.self]
    
    private var createdElementsById: [String: AnySolidNativeElement] = [:]
    
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
        createdElementsById[newElement.id.uuidString] = newElement
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
    
    

    
//    func registerCoreInJSContext() {
//
//
//
//        let isTextElement: @convention(block) (_ id: String) -> Bool = { [self] id in
//            createdElementsById[id]?.isTextElement ?? false
//        }
//
//        let createTextElement: @convention(block) () -> String = {
//            createElementByName(SNTextElement.name)
//        }
//
//        let removeElement: @convention(block) (_ parentId: String, _ childId: String) -> Void = { [self] parentId, childId in
//            if let parentElement = createdElementsById[parentId],
//               let childElement = createdElementsById[childId] {
//                parentElement.removeChild(element: childElement)
//                createdElementsById[childId] = nil
//            }
//        }
//
//        let insertElement:
//        @convention(block) (
//            _ parentId: String,
//            _ newElementId: String,
//            _ anchorId: String?) -> Void
//        = {
//            [self] parentId, newElementId, anchorId in
//            if let parentElement = createdElementsById[parentId],
//               let newElement = createdElementsById[newElementId] {
//                parentElement.insertBefore(element: newElement, anchor: (anchorId != nil) ? createdElementsById[anchorId!] : nil)
//            }
//        }
//
//        let getParentElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
//            self.createdElementsById[elementId]?.parentElement?.id.uuidString
//        }
//
//        let getFirstChildElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
//            self.createdElementsById[elementId]?.firstChild?.id.uuidString
//            // return nil
//        }
//
//        let getNextSiblingElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
//            self.createdElementsById[elementId]?.next?.id.uuidString
//        }
//
//        let setPropertyOnElement: @convention(block) (_ elementId: String, _ propertyName: String, _ value: Any?) -> Void =
//        { elementId, propertyName, value in
//            if let element = self.createdElementsById[elementId] {
//                element.setProp(propertyName, value as Any)
//            }
//        }
//
//        let getRootElement: @convention(block) () -> String = {
//            self.rootElement.id.uuidString
//        }
//
//        let obj = JSValue(newObjectIn: SharedJSConext.sharedContext)!
//
//
//
//        let globalObject: [String:Any] = [
//            "createElementByName": createElementByName,
//            "createTextElement": createTextElement,
//            "getRootElement": getRootElement
//        ]
//
//        for (property, value) in globalObject {
//            obj.setValue(value, forProperty: property)
//        }
//
//        SharedJSConext.sharedContext.setObject(obj, forKeyedSubscript: "_SolidNativeCore" as NSString)
//
//    }
}
