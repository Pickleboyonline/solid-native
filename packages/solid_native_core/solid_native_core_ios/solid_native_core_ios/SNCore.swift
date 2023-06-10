//
//  SNCore.swift
//  solid_native_core_ios
//
//  Created by Imran Shitta-Bey on 6/9/23.
//  Copyright © 2023 orgName. All rights reserved.
//

import Foundation
import JavaScriptCore

@objc protocol SNCoreJSExportsProtocal: JSExport {
    func createElement(name: String) -> SNElementNodeJSExport
}

@objc public class SNCoreJSExport : NSObject, SNCoreJSExportsProtocal {
    var snCore: SNCore
    
    init(snCore: SNCore) {
        self.snCore = snCore
    }
    
    func createElement(name: String) -> SNElementNodeJSExport {
        return snCore.createElementNode(name: name).getJsObjectExport()
    }
}

class SNCore {
    let jsContext = JSContext()!
    
    /**
     Name:SNModule
     */
    var knownElementRegistry: [String: SNElement.Type] = [:]
    
    /**
     Meant for
     */
    func registerElement(name: String, elementType: SNElement.Type) {
        knownElementRegistry[name] = elementType.self
    }
    
    var createdElementsRegistry: [ViewTag: SNElementNode] = [:]
    
    var currentViewTagNum = 0
    
    func createElementNode(name: String) -> SNElementNode {
        let node = SNElementNode(value: SNElement(jsContext: jsContext, viewTag: String(currentViewTagNum)))
        currentViewTagNum += 1
        return node
    }
    
    func getNodeWith(viewTag: String?) -> SNElementNode? {
        if viewTag != nil {
            return createdElementsRegistry[viewTag!]
        }
        return nil
    }
    
    // Export a function that does operations, but returns viewtags instead of
    // nodes
    func registureGlobalObject() {
        let exportedModule = SNCoreJSExport(snCore: self)
        jsContext.setObject(exportedModule, forKeyedSubscript: "SNCore" as NSString)
    }
    
    
    func start() {
        // Create global object that has createElement.
        registureGlobalObject()
        // jsContext.evaluateScript("SNCore.createElement('Text')")
    }
}
