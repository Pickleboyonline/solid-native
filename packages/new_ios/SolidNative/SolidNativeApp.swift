//
//  SolidNativeApp.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/20/24.
//

import SwiftUI
import Snmobile

@main
struct SolidNativeApp: App {
    
    let core: SolidNativeCore
    
    init() {
        core = SolidNativeCore()
        
        do {
            try core.start(jsUrl: "http://localhost:8080")
        } catch {
            print("Unexpected error: \(error).")
        }
        
    }
    
    var body: some Scene {
        WindowGroup {
            core.getRootNode().render()
        }
    }
}
