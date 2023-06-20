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
    var rootElement: AnySolidNativeElement
    let jsContext = JSContext()!
    
    
    
    init() {
        sharedSolidNativeCore.registerElements()
        let root = sharedSolidNativeCore.createRootElement(name: "v_stack")
        let button =  sharedSolidNativeCore.createRootElement(name: "button")
        let text = sharedSolidNativeCore.createRootElement(name: "sn_text")
        rootElement = root
        
        var inc = 0
        
        root.insertBefore(element: button, anchor: nil)
        
        root.insertBefore(element: text, anchor: button)
        
        
        text.setProp("text", "Count: 0")
        
        
        button.setProp("title", "Increment!")
        
        button.setProp("onPress", {
            inc += 1
            text.setProp("text", "Count \(inc)")
        })
        
        
        // rootView = core.createElement(name: "SNVStackElement")!
    }
    
    var body: some Scene {
        WindowGroup {
            Group {
                rootElement.render()
            }
            // Should be static view, that takes in the element ID
        }
    }
}

