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
    init() {
        let s = SNSnmobileSize(10, height: 10)
    }
    
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
