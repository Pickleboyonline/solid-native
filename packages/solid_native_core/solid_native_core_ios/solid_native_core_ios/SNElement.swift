//
//  SNBaseView.swift
//  solid_native_core_ios
//
//  Created by Imran Shitta-Bey on 6/9/23.
//  Copyright © 2023 orgName. All rights reserved.
//

import Foundation
import JavaScriptCore
import SwiftUI

typealias ViewTag = String

@objc protocol SNElementNodeJSExportProtocal: JSExport {
    var nextSibling: SNElementNodeJSExport? {get}
    var firstChild: SNElementNodeJSExport?  {get}
    var parentNode: SNElementNodeJSExport? {get}
    var viewTag: String {get}
    var isTextNode: Bool {get}
    
    var setProp: @convention(block) (String, Any) -> Void { get }
    
    // func setProp(name: String, value: Any) -> Void
    func removeChildNode(childNode: SNElementNodeJSExport) -> Void
    func addChildNodeBeforeAnchorNode(childNode: SNElementNodeJSExport, anchorNode: SNElementNodeJSExport?) -> Void
}

@objc public class SNElementNodeJSExport : NSObject, SNElementNodeJSExportProtocal {
    var node: SNElementNode
    
    init(node: SNElementNode) {
        self.node = node
    }
    
    dynamic var nextSibling: SNElementNodeJSExport? {
        get {
            self.node.nextSibling?.getJsObjectExport()
        }
    }
    
    dynamic var firstChild: SNElementNodeJSExport? {
        get {
            self.node.firstChild?.getJsObjectExport()
        }
    }
    
    dynamic var parentNode: SNElementNodeJSExport? {
        get {
            self.node.parentNode?.getJsObjectExport()
        }
    }
    
    dynamic var viewTag: String {
        get {
            self.node.value.viewTag
        }
    }
    
    dynamic var isTextNode: Bool {
        get {
            self.node.value.isTextNode()
        }
    }
    
    
    func removeChildNode(childNode: SNElementNodeJSExport) {
        node.removeChildNode(childNode: childNode.node)
    }

    
    lazy var setProp: @convention(block) (String, Any) -> Void = { name, value in
        self.node.value.setProp(name: name, value: value)
    }
    
    func addChildNodeBeforeAnchorNode(childNode: SNElementNodeJSExport, anchorNode: SNElementNodeJSExport?) {
        node.addChildNodeBeforeAnchorNode(childNode: childNode.node, anchorNode: anchorNode?.node)
    }
    
}




class SNElementNode:LinkedList<SNElement>.Node {
    
    // TODO: Build JSValue that represents its object here.
    // TODO: it has the ability to return other JS values for other element objects
    // TOOD: Need to work on reference equality
    
    var children = LinkedList<SNElement>()
    
    /**
     Note, this gets built once and sent over.
     */
    var jsObjectExport: SNElementNodeJSExport?
    
    var parentNode: SNElementNode?
    
    init(
        children: LinkedList<SNElement> = LinkedList<SNElement>(),
        parentNode: SNElementNode? = nil,
        value: SNElement
    ) {
        self.children = children
        self.parentNode = parentNode
        super.init(value: value)
    }
    
    func getJsObjectExport() -> SNElementNodeJSExport {
        if jsObjectExport != nil {
            return self.jsObjectExport!
        }
        return SNElementNodeJSExport(node: self)
    }
    
    var nextSibling: SNElementNode? {
        get {
            self.next as! SNElementNode?
        }
    }
    
    
    var firstChild: SNElementNode? {
        get {
            children.head as! SNElementNode?
        }
    }

    
    /**
     ISSUE: JS uses a JSValue to represent the object, which can cause some problems.
     We need a way to make sure JS can reference the same view natively.
     We can use  a map of viewTag -> SNElement
     */
    func addChildNodeBeforeAnchorNode(childNode: SNElementNode, anchorNode: SNElementNode?) {
        childNode.parentNode = self
        if anchorNode != nil {
            // We need the anchor nodes index.
            // ViewTag => Index
            // TODO: Can make this more effeicnt, currently opperates in O(n)
            children.insertBefore(newNode: childNode, beforeNode: anchorNode!)
        } else {
            children.insert(childNode, at: 0)
        }
    }
    
    func removeChildNode(childNode: SNElementNode) {
        childNode.parentNode = nil
        children.remove(node: childNode)
    }

    
    
}

struct Component: View {
    @ObservedObject var props: SNProps
    
    var body: some View {
        let name = props.getProp(name: "name", defaultValue: "Imran")
        Text("Hello, \(name)!")
    }
}

class SNElement {
    
    let jsContext: JSContext
    let jsObject: JSValue
    let props: SNProps
    
    var viewTag: String

    init(jsContext: JSContext, viewTag: ViewTag) {
        self.jsContext = jsContext
        self.jsObject = JSValue(newObjectIn: jsContext)
        self.props = SNProps()
        self.viewTag = viewTag
    }
    
    func setValue(name: String, value: Any) {
        jsObject.setValue(value, forProperty: name)
    }
    
    /**
     OVERIDE ME!!!!
     */
    func isTextNode() -> Bool {
        return true
    }
    

    func setProp(name: String, value: Any) {
        props.setProp(name: name, value: value)
    }
    
    @ViewBuilder
    func render() -> some View {
        Component(props: props)
    }
    
}
