//
//  SNRender.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore

class SNRender: SolidNativeModule {
    class var name: String {
        "SNRender"
    }
    
    // TODO: Move to core
    // TODO: This is really available via Core
    // TODO: Eventually, this needs to be abstracted further for
    // TODO: Rust implimentation.
    var viewTypes: [String: any SolidNativeViewP.Type]
    var rootNode: RenderViewNode
    
    var nodeRegistry: [String: RenderViewNode]
    
    
    required init() {
        viewTypes = [BaseSolidNativeView.name: BaseSolidNativeView.self]
        rootNode = RenderViewNode(view: BaseSolidNativeView())
        nodeRegistry = [rootNode.id: rootNode]
    }    
}


extension SNRender {
    /// Returns NodeID
    func createViewByName(_ name: String) -> String {
        if let viewType = viewTypes[name] {
            let node = RenderViewNode(view: viewType.init())
            nodeRegistry[node.id] = node
            return node.id
        }
        assertionFailure("\(name) is not in element registry!")
        return ""
    }
    
    func getJSValueRepresentation() -> JSValue {
        let builder = JSValueBuilder()
        
        builder.addSyncFunction("print") { (_ str: String) in
            print(str)
        }
        
        builder.addSyncFunction("getRootView") {
            return self.rootNode.id
        }
        
        builder.addSyncFunction("getFirstChild") { (_ id: String) in
            return self.nodeRegistry[id]?.id
        }
        
        builder.addSyncFunction("getParent") { (_ id: String) in
            self.nodeRegistry[id]?.parent?.id
        }
        
        builder.addSyncFunction("setProp") { (_ id: JSValue, name: JSValue, value: JSValue) in
            guard let node = self.nodeRegistry[id.toString()!] else {
                return
            }
           
            if value.isNull || value.isUndefined {
                node.removeProp(name.toString()!)
            } else {
                node.setProp(name.toString()!, value.toObject()!)
            }
            
        }
        
        builder.addSyncFunction("isTextElement") { (_ id: String) in
            return self.nodeRegistry[id]?.view.isTextElement
        }
        
        builder.addSyncFunction("removeChild") { (_ id: String, childId: String) in
            if let node = self.nodeRegistry[id],
               let childNode = self.nodeRegistry[id] {
                node.removeChild(childNode)
                self.nodeRegistry.removeValue(forKey: childId)
            }
        }
        
        builder.addSyncFunction("insertBefore") { (_ id: JSValue, elementId: JSValue, anchorId: JSValue) in
            
            guard let node = self.nodeRegistry[id.toString()!],
                  let element = self.nodeRegistry[elementId.toString()!] else {
                return
            }
            
            if anchorId.isString,
               let anchor = self.nodeRegistry[anchorId.toString()!] {
                node.insertBefore(element, anchor)
            } else {
                node.insertBefore(element, nil)
            }
        }
        
        builder.addSyncFunction("next") { (_ id: String) in
            return self.nodeRegistry[id]?.next?.id
        }
        
        builder.addSyncFunction("prev") { (_ id: String) in
            return self.nodeRegistry[id]?.prev?.id
        }
        
        builder.addSyncFunction("createNodeByName") { (_ name: String) in
            let viewId = self.createViewByName(name);
            return viewId
        }
        
        
        return builder.value
    }
}


