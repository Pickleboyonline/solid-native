//
//  SNRender.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore
/**
 Acts as renderer.
 */
class SNRender: SolidNativeModule {
    class override var name: String {
        "SNRender"
    }
    
    private let viewManager = ViewManager()
    
    private let rootView = SolidNativeCore.shared.rootElement
    
    required init() {
        let _ = viewManager.addViewToRegistry(view: rootView)
    }
    
    override func getJSValueRepresentation() -> JSValue {
        let builder = JSValueBuilder()
        
        builder.addSyncFunction("print") { (_ str: String) in
            print(str)
        }
        
        builder.addSyncFunction("getRootView") { 
            return self.rootView.id.uuidString
        }
        
        builder.addSyncFunction("getFirstChild") { (_ id: String) in
            let view = self.viewManager.getViewById(id)
            return view.id.uuidString
        }
        
        builder.addSyncFunction("getParent") { (_ id: String) in
            let view = self.viewManager.getViewById(id)
            return view.parentElement?.id.uuidString
        }
        
        builder.addSyncFunction("setProp") { (_ id: String, name: String, value: JSValue?) in
            let view = self.viewManager.getViewById(id)
            view.setProp(name, value)
        }
        
        builder.addSyncFunction("isTextElement") { (_ id: String) in
            let view = self.viewManager.getViewById(id)
            return view.isTextElement
        }
        
        builder.addSyncFunction("removeChild") { (_ id: String, childId: String) in
            let view = self.viewManager.getViewById(id)
            let viewChild = self.viewManager.getViewById(childId)
            view.removeChild(viewChild)
        }
        
        builder.addSyncFunction("insertBefore") { (_ id: String, elementId: String, anchorId: String) in
            let view = self.viewManager.getViewById(id)
            let element = self.viewManager.getViewById(elementId)
            let anchor = self.viewManager.getViewById(anchorId)
            return view.insertBefore(element, anchor)
        }
        
        builder.addSyncFunction("next") { (_ id: String) in
            let view = self.viewManager.getViewById(id)
            return view.next?.id.uuidString
        }
        
        builder.addSyncFunction("prev") { (_ id: String) in
            let view = self.viewManager.getViewById(id)
            return view.prev?.id.uuidString
        }
        
        
        return builder.value
    }
}


// Need some way to build function.
private class ViewManager {
    private var viewRegistry: [String : SolidNativeView.Type] = [:]
    private var createdViewRegistry: [String: SolidNativeView] = [:]
    
    fileprivate func addViewToRegistry(view: SolidNativeView) -> String {
        let id = view.id.uuidString
        createdViewRegistry[id] = view
        return id
    }

    // Only thing it needs is the View type, the View Name, whether or not it a text view
    func registerElement(_ viewType: SolidNativeView.Type) {
        viewRegistry[viewType.name] = viewType.self
    }
    /**
     @returns view id
     */
    func createViewByName(_ name: String) -> String {
        if let viewType = viewRegistry[name] {
            let newView = viewType.init()
            return addViewToRegistry(view: newView)
        }
        assertionFailure("\(name) is not in element registry!")
        return ""
    }
    
    func getViewById(_ id: String) -> SolidNativeView {
        if let view = createdViewRegistry[id] {
            return view
        }
        assertionFailure("view with id \(id) is not in element registry!")
        return SolidNativeView()
    }
    
}

