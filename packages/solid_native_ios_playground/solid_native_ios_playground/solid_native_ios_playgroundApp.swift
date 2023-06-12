//
//  solid_native_ios_playgroundApp.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import SwiftUI

@main
struct solid_native_ios_playgroundApp: App {
    var rootView = SolidNativeView()
    
    var body: some Scene {
        WindowGroup {
            rootView.render()
        }
    }
}
