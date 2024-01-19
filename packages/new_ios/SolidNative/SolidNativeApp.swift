//
//  SolidNativeApp.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/12/24.
//

import SwiftUI

@main
struct SolidNativeApp: App {
    
    init() {
        // TODO: Make better
        SolidNativeCore.shared.downloadAndRunJsBundleSync()
    }
    
    
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
