//
//  SolidNativeCore.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore

@objc protocol SolidNativeCoreJSExport: JSExport {
    var createElement: @convention(block) (_ name: String) -> String { get }
    
    var isTextElement: @convention(block) (String) -> Bool { get }
    
    var createTextElement: @convention(block) () -> String { get }
    
    var removeElement: @convention(block) (String, String) -> Void { get }
    
    var insertElement: @convention(block) (String, String, String?) -> Void { get }
    
    var getParentElementId: @convention(block) (String) -> String? { get }
    
    var getFirstChildElementId: @convention(block) (String) -> String? { get }
    
    var getNextSiblingElementId: @convention(block) (String) -> String? { get }
    
    var setPropertyOnElement: @convention(block) (String, String, Any?) -> Void { get }
    
    var getRootElement: @convention(block) () -> String { get }
}

typealias ElementName = String

let ROOT_ELEMENT_CLASS = SNVStackElement.self

@objc public class SolidNativeCore: NSObject, SolidNativeCoreJSExport {
    
    lazy var createElement: @convention(block) (_ name: String) -> String = { [self] name in
        if let elementType = elementRegistry[name] {
            let newElement = elementType.init()
            createdElementsById[newElement.id.uuidString] = newElement
            return newElement.id.uuidString
        }
        assertionFailure("\(name) is not in element registry!")
        return ""
    }
    
    lazy var isTextElement: @convention(block) (_ id: String) -> Bool = { [self] id in
        createdElementsById[id]?.isTextElement ?? false
    }
    
    lazy var createTextElement: @convention(block) () -> String = {
        self.createElement(SNTextElement.name)
    }
    
    lazy var removeElement: @convention(block) (_ parentId: String, _ childId: String) -> Void = { [self] parentId, childId in
        if let parentElement = createdElementsById[parentId],
           let childElement = createdElementsById[childId] {
            parentElement.removeChild(element: childElement)
            createdElementsById[childId] = nil
        }
    }
    
    lazy var insertElement:
    @convention(block) (
        _ parentId: String,
        _ newElementId: String,
        _ anchorId: String?) -> Void
    = {
        [self] parentId, newElementId, anchorId in
        if let parentElement = createdElementsById[parentId],
           let newElement = createdElementsById[newElementId] {
            parentElement.insertBefore(element: newElement, anchor: (anchorId != nil) ? createdElementsById[anchorId!] : nil)
        }
    }
    
    lazy var getParentElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
        self.createdElementsById[elementId]?.parentElement?.id.uuidString
    }
    
    lazy var getFirstChildElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
        self.createdElementsById[elementId]?.firstChild?.id.uuidString
        // return nil
    }
    
    lazy var getNextSiblingElementId: @convention(block) (_ elementId: String) -> String? = { elementId in
        self.createdElementsById[elementId]?.next?.id.uuidString
    }
    
    lazy var setPropertyOnElement: @convention(block) (_ elementId: String, _ propertyName: String, _ value: Any?) -> Void =
    { elementId, propertyName, value in
        if let element = self.createdElementsById[elementId] {
            element.setProp(propertyName, value as Any)
        }
    }
    
    
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
    
    
    lazy var getRootElement: @convention(block) () -> String = {
        self.rootElement.id.uuidString
    }
}
