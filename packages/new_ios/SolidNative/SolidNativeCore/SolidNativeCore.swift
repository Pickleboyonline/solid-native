//
//  SolidNativeCore.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore


/**
 
 */
class SolidNativeCore {
    static let shared = SolidNativeCore()
    
    let jsContext = JSContext()!
    private let moduleManager = ModuleManager()
    
    private init() {
        // jsContext.isInspectable = true
        TimerJS.registerInto(jsContext: jsContext)
        
        // Configure base module
        moduleManager.registerModule(SNRender.self)
        
        injectCoreIntoContext()
        
        // Needs to inject function to grab other modules from registry and return their JSValues
    }
    
    private func injectCoreIntoContext() {
        let getNativeModule: @convention(block) (_ name: String) -> JSValue = { str in
            return self.moduleManager.createModuleJsValueByName(str)!
        }
        jsContext.setObject(getNativeModule, forKeyedSubscript:
                                "_getNativeModule" as NSString)
    }
    
    // SNRender will pull it from the singleton
    let rootElement = SNView()
    
    
    func downloadAndRunJsBundleSync() {
        if let url = URL(string: "http://localhost:8080"),
           let sourceUrl = URL(string: "http://localhost:8080/source"){
            do {
                let bundle = try String(contentsOf: url)
                
                jsContext.exceptionHandler = { (_, value) in
                    print("JS Error: " + value!.toString()!)
                }
                
                let jsPrint: @convention(block) (_ contents: String) -> Void = { str in
                    print(str)
                }
                jsContext.isInspectable = true
                TimerJS.registerInto(jsContext: jsContext)
                jsContext.setObject(jsPrint, forKeyedSubscript: "_print" as NSString)
                // SharedJSConext.sharedContext.setObject(sharedSolidNativeCore, forKeyedSubscript: "_SolidNativeCore" as NSString)
                jsContext.evaluateScript(bundle, withSourceURL: sourceUrl)
            } catch {
                // contents could not be loaded
                print("Url was bad!")
            }
        } else {
            // the URL was bad!
            print("ERROR: Url was bad!")
        }
    }
    
}





private class ModuleManager {
    // Registry is needed to look
    private var moduleRegistry: [String : SolidNativeModule.Type] = [SNRender.name : SNRender.self];
    private var moduleJSValueRegistry: [String : JSValue] = [:];
    
    func registerModule(_ moduleType: SolidNativeModule.Type) {
        moduleRegistry[moduleType.name] = moduleType.self
    }
    
    func createModuleJsValueByName(_ name: String) -> JSValue? {
        // Should return a JS Value
        if let moduleType = moduleRegistry[name] {
            
            if let jsValue = moduleJSValueRegistry[name] {
                return jsValue
            } else {
                let newModule = moduleType.init()
                let jsValue = newModule.getJSValueRepresentation()
                
                moduleJSValueRegistry[name] = jsValue
                return jsValue
            }
        
        }
        assertionFailure("\(name) is not in module registry!")
        return nil
    }
    
}
