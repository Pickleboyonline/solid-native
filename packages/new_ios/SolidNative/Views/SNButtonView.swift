//
//  SNButtonView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import SwiftUI

class SNButtonView: SolidNativeView {

  class override var name: String {
    "sn_button"
  }

  struct SNButton: View {
    @ObservedObject var props: SolidNativeProps

    func onPress() {
      if let callback = props.getPropAsJSValue(name: "onPress") {
        callback.call(withArguments: nil)
      }
    }

    var body: some View {
      let title = props.getString(name: "title")
      Button(title) {
        onPress()
      }
    }
  }

  override func render() -> AnyView {
    AnyView(SNButton(props: self.props))
  }

}
