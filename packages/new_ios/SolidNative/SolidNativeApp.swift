//
//  SolidNativeApp.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/12/24.
//

import SwiftUI
import Snmobile

@main
struct SolidNativeApp: App {

  init() {
      let s = SNSnmobileSize(20, height: 30)!
      
      print(s.width)
  }

  var body: some Scene {
    WindowGroup {

    }
  }
}
