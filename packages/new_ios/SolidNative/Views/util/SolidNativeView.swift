//
//  SNView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/18/24.
//

import Foundation
import JavaScriptCore
import SwiftUI
import Yoga
import YogaSwiftUI

class SolidNativeView {
    dynamic var next: SolidNativeView?
    dynamic var prev: SolidNativeView?
    let id = UUID()
    
    // TODO: Make Static helper method to apply Layout Styles. Takes view and props and does what it needs
    static func applyLayoutStyles(view: any View, props: SolidNativeProps) -> any View {
        var styledView = view

        if let justifyContent = YogaTypeMarshaller.convertToYGJustify(props.getPropAsJSValue(name: "justifyContent")) {
            styledView = styledView.justifyContent(justifyContent)
        }
        
        if let alignItems = YogaTypeMarshaller.convertToYGAlign(props.getPropAsJSValue(name: "alignItems")) {
            styledView = styledView.alignItems(alignItems)
        }
        
        if let alignSelf = YogaTypeMarshaller.convertToYGAlign(props.getPropAsJSValue(name: "alignSelf")) {
            styledView = styledView.alignSelf(alignSelf)
        }
        
        if let flexDirection = YogaTypeMarshaller.convertToYGFlexDirection(props.getPropAsJSValue(name: "flexDirection")) {
            styledView = styledView.flexDirection(flexDirection)
        }
        
        if let flexWrap = YogaTypeMarshaller.convertToYGWrap(props.getPropAsJSValue(name: "flexWrap")) {
            styledView = styledView.flexWrap(flexWrap)
        }
        
        if let flexBasis = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "flexBasis")) {
            styledView = styledView.flexBasis(flexBasis)
        }
        
        if let flexGrow = YogaTypeMarshaller.convertToCGFloat(props.getPropAsJSValue(name: "flexGrow")) {
            styledView = styledView.flexGrow(flexGrow)
        }
        
        if let flexShrink = YogaTypeMarshaller.convertToCGFloat(props.getPropAsJSValue(name: "flexShrink")) {
            styledView = styledView.flexShrink(flexShrink)
        }
        
        if let flex = YogaTypeMarshaller.convertToCGFloat(props.getPropAsJSValue(name: "flex")) {
            
            if flex > 0 {
                styledView = styledView.flexGrow(flex)
                styledView = styledView.flexShrink(1)
                styledView = styledView.flexBasis(YGValue(value: 0, unit: .point))
            } else if flex == 0 {
                styledView = styledView.flexGrow(flex)
                styledView = styledView.flexShrink(1)
                styledView = styledView.flexBasis(YGValue(value: 0, unit: .auto))
            } else if flex == -1 {
                styledView = styledView.flexGrow(flex)
                styledView = styledView.flexShrink(1)
                styledView = styledView.flexBasis(YGValue(value: 0, unit: .auto))
            }
            
        }
        
        if let width = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "width")) {
            styledView = styledView.width(width)
        }
        
        if let maxWidth = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "maxWidth")) {
            styledView = styledView.maxWidth(maxWidth)
        }
        
        if let minWidth = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "minWidth")) {
            styledView = styledView.minWidth(minWidth)
        }
        
        if let height = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "height")) {
            styledView = styledView.height(height)
        }
        
        if let maxHeight = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "maxHeight")) {
            styledView = styledView.maxHeight(maxHeight)
        }
        
        if let minHeight = YogaTypeMarshaller.convertToYogaDimension(props.getPropAsJSValue(name: "minHeight")) {
            styledView = styledView.minHeight(minHeight)
        }
        
        if let position = YogaTypeMarshaller.convertToYGPositionType(props.getPropAsJSValue(name: "position")) {
            styledView = styledView.position(position)
        }
        
        if let marginTop = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "marginTop")) {
            styledView = styledView.marginTop(marginTop)
        }
        
        if let marginRight = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "marginRight")) {
            styledView = styledView.marginRight(marginRight)
        }
        
        if let marginBottom = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "marginBottom")) {
            styledView = styledView.marginBottom(marginBottom)
        }
        
        if let marginLeft = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "marginLeft")) {
            styledView = styledView.marginLeft(marginLeft)
        }
        
        if let paddingTop = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "paddingTop")) {
            styledView = styledView.paddingTop(paddingTop)
        }
        
        if let paddingRight = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "paddingRight")) {
            styledView = styledView.paddingRight(paddingRight)
        }
        
        if let paddingBottom = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "paddingBottom")) {
            styledView = styledView.paddingBottom(paddingBottom)
        }
        
        if let paddingLeft = YogaTypeMarshaller.convertToYGValue(props.getPropAsJSValue(name: "paddingLeft")) {
            styledView = styledView.paddingLeft(paddingLeft)
        }
        
        if let borderWidth = YogaTypeMarshaller.convertToCGFloat(props.getPropAsJSValue(name: "borderWidth")) {
            styledView = styledView.borderWidth(borderWidth)
        }
        
        return styledView
    }
    
    class var name: String {
        "sn_view"
    }
    
    required init() {
    }
    
    
    let props = SolidNativeProps()
    
    @objc public func setProp(_ name: String, _ value: JSValue?) {
        assert(name != "children", "Err: User `removeChild` or `insertBefore` to update children!")
        props.values[name] = value
        updateCount()
    }
    
    // Forces parent component to render if this changes
    // Used for nest text components
    func updateCount() {
        if !isTextElement {
            return
        }
        var currentNode = self
        var prevNode = self
        // Go up chain,
        while currentNode.isTextElement && currentNode.parentElement != nil {
            if let parentElement = currentNode.parentElement {
                prevNode = currentNode
                currentNode = parentElement
            }
        }
        prevNode.props.updateCount = (prevNode.props.updateCount + 1) % 10
    }
    
    
    var children: [SolidNativeView] = []
    
    // Can be getter
    dynamic var firstChild: SolidNativeView?
    
    dynamic var parentElement: SolidNativeView?
    
    // Iterate over first child prop
    // O(n)
    private func updateChildrenInProps() {
        var newChildren: [SolidNativeView] = []
        
        var nextChild = firstChild
        while let child = nextChild {
            newChildren.append(child)
            nextChild = child.next
        }
        
        children = newChildren
        props.children = newChildren
        updateCount()
    }
    
    
    // O(1)
    func removeChild(_ element: SolidNativeView) {
        // Link the nodes prev and next of it
        if let childNextSibling = element.next,
           let childPrevSibling = element.prev {
            childPrevSibling.next = childNextSibling
            childNextSibling.prev = childPrevSibling
            // Is first element
        } else if let childNextSibling = element.next {
            childNextSibling.prev = nil
            firstChild = childNextSibling
        } else if let childPrevSibling = element.prev {
            childPrevSibling.next = nil
        } else {
            firstChild = nil
        }
        element.parentElement = nil
        props.parent = nil
        element.next = nil
        element.prev = nil
        // Update the nodes first child
        // Remove as parent
        updateChildrenInProps()
    }
    
    // O(1)
    func insertBefore(_ element: SolidNativeView, _ anchor: SolidNativeView?) {
        // If no anchor set first child to view (make head)
        //
        if let anchor = anchor {
            
            if anchor === firstChild {
                anchor.prev = element
                element.next = anchor
                firstChild = element
            } else if let anchorPrevSibling = anchor.prev {
                anchorPrevSibling.next = element
                element.prev = anchorPrevSibling
                anchor.prev = element
                element.next = anchor
            }
            
        } else if let firstChild = firstChild {
            // Make it at the end of the list
            var nextChild: SolidNativeView? = firstChild
            while let child = nextChild {
                nextChild = child.next
                
                if nextChild == nil {
                    let lastChild = child
                    lastChild.next = element;
                    element.prev = lastChild
                }
            }
            
        } else {
            firstChild = element
        }
        
        element.parentElement = self
        props.parent = nil
        
        updateChildrenInProps()
    }
    
    // TODO: You need to override this!
    // This is how SolidJS will deliver a text prop.
    dynamic var isTextElement: Bool {
        false
    }
    
    
    func render() -> AnyView {
        AnyView(EmptyView())
    }
}


private class YogaTypeMarshaller {
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
