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
        let root = sharedSolidNativeCore.createRootElement(name: "sn_v_stack_view")
        let button =  sharedSolidNativeCore.createRootElement(name: "sn_button_view")
        let text = sharedSolidNativeCore.createRootElement(name: "sn_text_view")
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
        
        var toggle = true
        
        root.setProp("value", toggle)
        
        root.setProp("onChange") {
            (_ newValue: Bool) -> Void in
            toggle = newValue
            root.setProp("value", newValue)
        }
        
        Timer.scheduledTimer(withTimeInterval: 1.0, repeats: true) { _ in
            toggle = !toggle
            root.setProp("value", toggle)
        }
    }
    
    var body: some Scene {
        WindowGroup {
            Group {
                rootElement.render()
            }
        }
    }
}

