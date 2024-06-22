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
            try core.start(jsUrl: "http://10.0.0.175:8080")
            
//            for var (key, view) in core.viewWrapperRegistry {
//                print("For Key", key)
//                print(view.layoutMetrics.width, view.layoutMetrics.height)
//            }
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
