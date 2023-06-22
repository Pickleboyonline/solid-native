//
//  solid_native_ios_playgroundApp.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import SwiftUI
import JavaScriptCore

let sharedSolidNativeCore = SolidNativeCore()

@main
struct solid_native_ios_playgroundApp: App {

    init() {
        // Initialize core object into global
        // Download JS
        // print("JSValue: " + SharedJSConext.sharedContext.evaluateScript("globalThis").toString()!)
        sharedSolidNativeCore.registerElements()
        print("root element Id: " + sharedSolidNativeCore.rootElement.id.uuidString)
        // setupApp()
        setupAppSync()
    }
    
    var body: some Scene {
        WindowGroup {
            Group {
                sharedSolidNativeCore.rootElement.render()
            }
        }
    }
}


func setupApp() {

    Task {
        if let url = URL(string: "http://localhost:8080") {
            do {
                let (data, res) = try await URLSession.shared.data(from: url)
    
                guard let res = res as? HTTPURLResponse,
                      res.statusCode == 200 /* OK */ else {
                    throw URLError(.badServerResponse)
                }
    
                let bundle = String(decoding: data, as: UTF8.self)
                
                SharedJSConext.sharedContext.exceptionHandler = { (_, value) in
                    print("Error: " + value!.toString()!)
                }
    
                let jsPrint: @convention(block) (_ contents: String) -> Void = { str in
                    print(str)
                }
                
                SharedJSConext.sharedContext.isInspectable = true
                TimerJS.registerInto(jsContext: SharedJSConext.sharedContext)
                SharedJSConext.sharedContext.setObject(jsPrint, forKeyedSubscript: "_print" as NSString)
                SharedJSConext.sharedContext.setObject(sharedSolidNativeCore, forKeyedSubscript: "_SolidNativeCore" as NSString)
                SharedJSConext.sharedContext.evaluateScript(bundle)
                
    
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

func setupAppSync() {
    
    if let url = URL(string: "http://localhost:8080") {
        do {
            let bundle = try String(contentsOf: url)
            
            SharedJSConext.sharedContext.exceptionHandler = { (_, value) in
                print("Error: " + value!.toString()!)
            }

            let jsPrint: @convention(block) (_ contents: String) -> Void = { str in
                print(str)
            }
            SharedJSConext.sharedContext.isInspectable = true
            TimerJS.registerInto(jsContext: SharedJSConext.sharedContext)
            SharedJSConext.sharedContext.setObject(jsPrint, forKeyedSubscript: "_print" as NSString)
            SharedJSConext.sharedContext.setObject(sharedSolidNativeCore, forKeyedSubscript: "_SolidNativeCore" as NSString)
            SharedJSConext.sharedContext.evaluateScript(bundle)
            

        } catch {
            // contents could not be loaded
            print("Url was bad!")
        }
    } else {
        // the URL was bad!
        print("ERROR: Url was bad!")
    }
    
}
