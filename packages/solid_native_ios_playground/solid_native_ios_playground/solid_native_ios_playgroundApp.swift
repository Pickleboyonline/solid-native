//
//  solid_native_ios_playgroundApp.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import SwiftUI
import JavaScriptCore

@main
struct solid_native_ios_playgroundApp: App {
    
    var core: SolidNativeCore
    var rootView: AnySolidNativeElement
    let jsContext = JSContext()!
    
    init() {
        core = SolidNativeCore()
        core.registerElements()
        let textElement = core.createElement(name: "SNTextElement")!
        textElement.setText(text: "Hello World!")
        rootView = textElement
        // rootView = core.createElement(name: "SNVStackElement")!
    }
    
    var body: some Scene {
        WindowGroup {
            rootView.render()
            // Should be static view, that takes in the element ID
        }
    }
}


func renderView(elemnentId: Int) {
    // Get element
    // Get viewType
    // Get view
    // Get element props <== Should be mutable, observable object that shadow the view props
    // return view
}

// How to handle Refs:
// Refs are JSValue Objects that are passed by the onRef callback.
// If that callback exists,
// It may be easier to represent the views with a class and ViewBuilder, but eh.
