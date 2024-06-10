//
//  YogaTypeMarshaller.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/10/24.
//

import Foundation
import JavaScriptCore
import Yoga
import YogaSwiftUI

public class YogaTypeMarshaller {
  // Helper methods for type conversions
  static func convertToYGValue(_ jsValue: JSValue?) -> YGValue? {
    if let jsValue = jsValue {

      if jsValue.isString {
        let stringValue = jsValue.toString()!
        if let floatValue = Float(stringValue.replacingOccurrences(of: "%", with: "")) {
          return YGValue(value: floatValue / 100, unit: .percent)
        }
      } else if jsValue.isNumber {
        return YGValue(value: jsValue.toNumber().floatValue, unit: .point)
      } else if jsValue.isUndefined {
        return YGValue(value: 0, unit: .undefined)
      }
    }
    return nil
  }

  static func convertToYogaDimension(_ jsValue: JSValue?) -> YogaDimension? {
    if let jsValue = jsValue {
      if jsValue.isString {
        let stringValue = jsValue.toString()!
        if stringValue == "auto" {
          return .auto
        } else if let floatValue = Float(stringValue.replacingOccurrences(of: "%", with: "")) {
          return .percent(floatValue / 100)
        }
      } else if jsValue.isNumber {
        return .point(jsValue.toNumber().floatValue)
      }
    }
    return nil
  }

  static func convertToCGFloat(_ jsValue: JSValue?) -> CGFloat? {
    if let jsValue = jsValue, jsValue.isNumber {
      return CGFloat(jsValue.toNumber().floatValue)
    }
    return nil
  }

  static func convertToYGWrap(_ jsValue: JSValue?) -> YGWrap? {
    if let jsValue = jsValue, jsValue.isString {
      switch jsValue.toString() {
      case "nowrap":
        return .noWrap
      case "wrap":
        return .wrap
      case "wrap-reverse":
        return .wrapReverse
      default:
        return .noWrap
      }
    }
    return nil
  }

  static func convertToYGAlign(_ jsValue: JSValue?) -> YGAlign? {
    if let jsValue = jsValue, jsValue.isString {
      switch jsValue.toString() {
      case "flex-start":
        return .flexStart
      case "center":
        return .center
      case "flex-end":
        return .flexEnd
      case "stretch":
        return .stretch
      case "baseline":
        return .baseline
      default:
        return .auto
      }
    }
    return nil
  }

  static func convertToYGJustify(_ jsValue: JSValue?) -> YGJustify? {
    if let jsValue = jsValue, jsValue.isString {
      switch jsValue.toString() {
      case "flex-start":
        return .flexStart
      case "center":
        return .center
      case "flex-end":
        return .flexEnd
      case "space-between":
        return .spaceBetween
      case "space-around":
        return .spaceAround
      case "space-evenly":
        return .spaceEvenly
      default:
        return .flexStart
      }
    }
    return nil
  }

  static func convertToYGPositionType(_ jsValue: JSValue?) -> YGPositionType? {
    if let jsValue = jsValue, jsValue.isString {
      switch jsValue.toString() {
      case "relative":
        return .relative
      case "absolute":
        return .absolute
      default:
        return .relative
      }
    }
    return nil
  }

  static func convertToYGFlexDirection(_ jsValue: JSValue?) -> YGFlexDirection? {
    if let jsValue = jsValue, jsValue.isString {
      switch jsValue.toString() {
      case "row":
        return .row
      case "row-reverse":
        return .rowReverse
      case "column":
        return .column
      case "column-reverse":
        return .columnReverse
      default:
        return .column
      }
    }
    return nil
  }

}
