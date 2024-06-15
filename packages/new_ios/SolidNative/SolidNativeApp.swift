//
//  SolidNativeApp.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/12/24.
//

import SwiftUI
import YogaSwiftUI

@main
struct SolidNativeApp: App {

  init() {

      
  }

  var body: some Scene {
    WindowGroup {
      SolidNativeCore.shared.rootElement.render()
    }
  }
}
