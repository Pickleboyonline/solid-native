//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import JavaScriptCore
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
      dfs(start: props).applySolidNativeLayout(props: props).applySolidNativeStyles(props: props)
    }
  }

  override func render() -> AnyView {
    print("Render text with value:" + props.getString(name: "text"))
    return AnyView(SNTextView(props: self.props))
  }

}

// Call for node without text parent and with children
func dfs(start: SolidNativeProps) -> Text {
  // print("RUN!")
  // If theres children, essentially we want the text of those children.

  // If no children, we return

  let childrenCount = start.getChildren().count

  if childrenCount == 0 {
    return processTextView(start)
  }

  var txt = dfs(start: start.children[0].props)

  for i in 0...(childrenCount - 1) {
    if i == 0 {
      continue
    }

    let child = start.children[i]

    // Get other children and do the same
    let newText = dfs(start: child.props)

    txt = txt + newText

  }

  return styleTextViewFromSNView(props: start, text: txt)
}

// Convert view with no children to text
func processTextView(_ props: SolidNativeProps) -> Text {
  return Text(props.getString(name: "text"))
}

func styleTextViewFromSNView(props: SolidNativeProps, text: Text) -> Text {
  var styledText = text

  let style: JSValue? = props.getProp(name: "style")

  if let style = style, let styleDict = style.toDictionary() {

    if let color = styleDict["color"] as? String {
      styledText = styledText.foregroundColor(Color(hex: color))
    }

    let fontWeight = styleDict["fontWeight"] as? String ?? "regular"
    let swiftUIFontWieght = toFontWeight(fontWeight)

    let fontSize = styleDict["fontSize"] as? Float ?? 17.0

    styledText = styledText.font(
      .system(
        size: CGFloat(fontSize),
        weight: swiftUIFontWieght))

    if let textDecorationLine = styleDict["textDecorationLine"] as? String {
      styledText = applyTextDecoration(text: styledText, decoration: textDecorationLine)
    }

    if let fontStyle = styleDict["fontStyle"] as? String, fontStyle == "italic" {
      styledText = styledText.italic()
    }

    if let numberOfLines = styleDict["numberOfLines"] as? Int, numberOfLines >= 0 {
      styledText = styledText.lineLimit(numberOfLines) as! Text
    }

    if let ellipsizeMode = styleDict["ellipsizeMode"] as? String {
      styledText = styledText.truncationMode(ellipsizeModeToTruncationMode(ellipsizeMode)) as! Text
    }

  }

  return styledText
}

// Helper functions and extensions to support the additional properties

func applyTextDecoration(text: Text, decoration: String) -> Text {
  var newText = text
  switch decoration {
  case "underline":
    newText = newText.underline()
  case "line-through":
    newText = newText.strikethrough()
  default:
    break
  }
  return newText
}

// Additional extensions for fontWeight, fontStyle, etc., can be modeled similar to the UIColor extension provided previously.
func toFontWeight(_ fontWeight: String?) -> Font.Weight {
  switch fontWeight {
  case "100":
    return .ultraLight
  case "200":
    return .thin
  case "300", "light":
    return .light
  case "400", "normal":
    return .regular
  case "500", "medium":
    return .medium
  case "600":
    return .semibold
  case "700", "bold":
    return .bold
  case "800":
    return .heavy
  case "900":
    return .black
  default:
    return .regular
  }
}

func ellipsizeModeToTruncationMode(_ ellipsizeMode: String?) -> Text.TruncationMode {
  // TODO: Add support for "clip".
  switch ellipsizeMode {
  case "head":
    return .head
  case "middle":
    return .middle
  default:
    return .tail
  }
}

