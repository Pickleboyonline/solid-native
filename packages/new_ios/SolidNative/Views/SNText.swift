//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI

class SNTextView: SolidNativeView {

  class override var name: String {
    "sn_text"
  }

  override var isTextElement: Bool {
    true
  }

  struct SNTextView: View {

    @ObservedObject var props: SolidNativeProps

    var body: some View {
      dfs(start: props)
    }
  }

  override func render() -> AnyView {
    AnyView(SNTextView(props: self.props))
  }

}

// Call for node without text parent and with children
func dfs(start: SolidNativeProps) -> Text {
  print("RUN!")
  // If theres children, essentially we want the text of those children.

  // If no children, we return

  let childrenCount = start.getChildren().count

  if childrenCount == 0 {
    return processTextView(start)
  }

  var txt = Text("")

  for i in 0...(childrenCount - 1) {
    let child = start.children[i]
    // Get other children and do the same
    let newText = dfs(start: child.props)
    if i == 0 {
      txt = newText
    } else {
      txt = txt + newText
    }
  }

  return styleTextViewFromSNView(props: start, text: txt)
}

// For views that do have children, wrap them with this to properly deal with it.
func styleTextViewFromSNView(props: SolidNativeProps, text: Text) -> Text {
  // TODO: Style it
  if props.getBool(name: "bold") {
    return text.bold()
  }
  return text
}

// Convert view with no children to text
func processTextView(_ props: SolidNativeProps) -> Text {
  return Text(props.getString(name: "text"))
}
