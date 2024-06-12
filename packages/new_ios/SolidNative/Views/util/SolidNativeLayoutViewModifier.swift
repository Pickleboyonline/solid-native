//
//  SolidNativeViewStyler.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/10/24.
//

import Foundation
import JavaScriptCore
import SwiftUI
import Yoga

struct SolidNativeLayoutViewModifier: ViewModifier {
  var props: SolidNativeProps

  func body(content: Content) -> some View {
    applyLayoutStyles(props: props) {
      content
    }
  }
}

extension View {
  func applySolidNativeLayout(props: SolidNativeProps) -> some View {
    modifier(SolidNativeLayoutViewModifier(props: props))
  }
}

private func applyLayoutStyles(
  props: SolidNativeProps, @ViewBuilder contentBuilder: () -> some View
) -> some View {
  var styledView = AnyView(contentBuilder())
  let style: JSValue? = props.getProp(name: "style")

  guard let style = style
  else {
    return styledView
  }

  if !style.isObject {
    return styledView
  }

  if let justifyContent = YogaTypeMarshaller.convertToYGJustify(
    style.objectForKeyedSubscript("justifyContent"))
  {

    styledView = AnyView(styledView.justifyContent(justifyContent))
  }

  if let alignItems = YogaTypeMarshaller.convertToYGAlign(
    style.objectForKeyedSubscript("alignItems")
  ) {
    styledView = AnyView(styledView.alignItems(alignItems))
  }

  if let alignSelf = YogaTypeMarshaller.convertToYGAlign(
    style.objectForKeyedSubscript("alignSelf")
  ) {
    styledView = AnyView(styledView.alignSelf(alignSelf))
  }

  if let flexDirection = YogaTypeMarshaller.convertToYGFlexDirection(
    style.objectForKeyedSubscript("flexDirection")
  ) {
    styledView = AnyView(styledView.flexDirection(flexDirection))
  }

  if let flexWrap = YogaTypeMarshaller.convertToYGWrap(
    style.objectForKeyedSubscript("flexWrap")
  ) {
    styledView = AnyView(styledView.flexWrap(flexWrap))
  }

  if let flexBasis = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("flexBasis")
  ) {
    styledView = AnyView(styledView.flexBasis(flexBasis))
  }

  if let flexGrow = YogaTypeMarshaller.convertToCGFloat(
    style.objectForKeyedSubscript("flexGrow")
  ) {
    styledView = AnyView(styledView.flexGrow(flexGrow))
  }

  if let flexShrink = YogaTypeMarshaller.convertToCGFloat(
    style.objectForKeyedSubscript("flexShrink")
  ) {
    styledView = AnyView(styledView.flexShrink(flexShrink))
  }

  if let flex = style.objectForKeyedSubscript("flex"),
    flex.isNumber,
    let flex = flex.toNumber()
  {
    if Float(truncating: flex) > 0 {
      styledView = AnyView(styledView.flexGrow(CGFloat(truncating: flex)))
      styledView = AnyView(styledView.flexShrink(1))
      styledView = AnyView(styledView.flexBasis(YGValue(value: 0, unit: .point)))
    } else if flex == 0 {
      styledView = AnyView(styledView.flexGrow(CGFloat(truncating: flex)))
      styledView = AnyView(styledView.flexShrink(1))
      styledView = AnyView(styledView.flexBasis(YGValue(value: 0, unit: .auto)))
    } else if flex == -1 {
      styledView = AnyView(styledView.flexGrow(CGFloat(truncating: flex)))
      styledView = AnyView(styledView.flexShrink(1))
      styledView = AnyView(styledView.flexBasis(YGValue(value: 0, unit: .auto)))
    }
  }

  if let width = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("width")
  ) {
    styledView = AnyView(styledView.width(width))
  }

  if let maxWidth = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("maxWidth")
  ) {
    styledView = AnyView(styledView.maxWidth(maxWidth))
  }

  if let minWidth = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("minWidth")
  ) {
    styledView = AnyView(styledView.minWidth(minWidth))
  }

  if let height = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("height")
  ) {
    styledView = AnyView(styledView.height(height))
  }

  if let maxHeight = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("maxHeight")
  ) {
    styledView = AnyView(styledView.maxHeight(maxHeight))
  }

  if let minHeight = YogaTypeMarshaller.convertToYogaDimension(
    style.objectForKeyedSubscript("minHeight")
  ) {
    styledView = AnyView(styledView.minHeight(minHeight))
  }

  if let position = YogaTypeMarshaller.convertToYGPositionType(
    style.objectForKeyedSubscript("position")
  ) {
    styledView = AnyView(styledView.position(position))
  }

  if let marginTop = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("marginTop")
  ) {
    styledView = AnyView(styledView.marginTop(marginTop))
  }

  if let marginRight = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("marginRight")
  ) {
    styledView = AnyView(styledView.marginRight(marginRight))
  }

  if let marginBottom = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("marginBottom")
  ) {
    styledView = AnyView(styledView.marginBottom(marginBottom))
  }

  if let marginLeft = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("marginLeft")
  ) {
    styledView = AnyView(styledView.marginLeft(marginLeft))
  }

  if let paddingTop = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("paddingTop")
  ) {
      styledView = AnyView(styledView.padding(.top, CGFloat(paddingTop.value) ))
  }

  if let paddingRight = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("paddingRight")
  ) {
      
      styledView = AnyView(styledView.padding(.trailing, CGFloat(paddingRight.value) ))
  }

  if let paddingBottom = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("paddingBottom")
  ) {
      styledView = AnyView(styledView.padding(.bottom, CGFloat(paddingBottom.value) ))
  }

  if let paddingLeft = YogaTypeMarshaller.convertToYGValue(
    style.objectForKeyedSubscript("paddingLeft")
  ) {
      styledView = AnyView(styledView.padding(.trailing, CGFloat(paddingLeft.value) ))
  }

  if let borderWidth = YogaTypeMarshaller.convertToCGFloat(
    style.objectForKeyedSubscript("borderWidth")
  ) {
    styledView = AnyView(styledView.borderWidth(borderWidth))
  }
  return styledView
}
