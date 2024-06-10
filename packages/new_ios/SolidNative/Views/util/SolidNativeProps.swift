//
//  SnViewProps.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore
import SwiftUI

class SolidNativeProps: ObservableObject {
  @Published var values: [String: JSValue?] = [:]
  // TODO: Type this!
  @Published var children: [SolidNativeView] = []

  @Published var parent: SolidNativeView?

  @Published var updateCount = 0

  func getProp<T>(name: String) -> T? {
    if let prop = values[name] as? T {
      return prop
    }
    return nil
  }

  func getString(name: String, `default`: String = "") -> String {
    if let prop = (values[name] ?? nil) {
      return prop.toString()
    }
    return `default`
  }

  func getBool(name: String, `default`: Bool = false) -> Bool {
    if let prop = (values[name] ?? nil) {
      //print("Found Bool!")
      return prop.toBool()
    }
    return `default`
  }

  func getNumber(name: String, `default`: NSNumber = 0) -> NSNumber {
    if let prop = (values[name] ?? nil) {
      return prop.toNumber()
    }
    return `default`
  }

  func getChildren() -> [SolidNativeView] {
    children
  }

  func getPropAsJSValue(name: String) -> JSValue? {
    values[name] ?? nil
  }

  func callCallbackWithArgs(name: String, args: [Any]) {
    if let callback = getPropAsJSValue(name: "onPress") {
      callback.call(withArguments: args)
    }
  }
}
