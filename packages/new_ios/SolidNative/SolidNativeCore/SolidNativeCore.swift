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
    
    private init() {
        // jsContext.isInspectable = true
        TimerJS.registerInto(jsContext: jsContext)
    }
    
    
    func getRootElement() {
        
    }
    
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
