//
//  SNViewWrapper.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation
import JavaScriptCore
import Yoga
import SwiftUI


// TODO: Eventually move this to rust for all Android+iOS Support
/// Handles
class RenderViewNode {
    let id = UUID().uuidString
    
    // MARK: Tree Related Props
    var next: RenderViewNode?
    var prev: RenderViewNode?
    var children: [RenderViewNode] = []
    var parent: RenderViewNode?
    var firstChild: RenderViewNode?
    
    // MARK: Layout Node
    let yogaNodeRef = YGNodeNew()!
    /// If key is is here, and doesn't match, you'll need to set it to undefined.
    var prevYogaStyleKeys: Set<String> = []
    // MARK: View State Management
    // TODO: Determine how props are handled
    var props: [String: JSValue] = [:]
    // TODO: Will prob have to be a class eventually for Rust/C++ interop
    let viewWrapper: SolidNativeViewWrapper
    // TODO: Determine what effect this has on SolidJS
    var isTextElement: Bool { viewWrapper.solidNativeViewType.isTextElement }
    
    init(viewWrapper: SolidNativeViewWrapper) {
        self.viewWrapper = viewWrapper
    }
}

// TODO: Manage Yoga layout system Node tree
extension RenderViewNode {
    func updateChildren() {
        var newChildren: [RenderViewNode] = []

        var nextChild = firstChild
        while let child = nextChild {
          newChildren.append(child)
          nextChild = child.next
        }
        children = newChildren
    }
    func setProp(_ name: String, _ value: JSValue) {
        viewWrapper.props.setProp(name, value)
        if name == "style" {
            updateYogaNodeStyle(value)
        }
        // TODO: Updates props on view wrapper
    }
    func removeProp(_ name: String) {
        viewWrapper.props.removeProp(name)
        // TODO: Updates props on view wrapper
    }
    
    // TODO: Use array methods to find index of children for Yoga Compadability
    func removeChild(_ childNode: RenderViewNode) {
        // Link the nodes prev and next of it
        if let childNextSibling = childNode.next,
          let childPrevSibling = childNode.prev
        {
          childPrevSibling.next = childNextSibling
          childNextSibling.prev = childPrevSibling
          // Is first element
        } else if let childNextSibling = childNode.next {
          childNextSibling.prev = nil
          firstChild = childNextSibling
        } else if let childPrevSibling = childNode.prev {
          childPrevSibling.next = nil
        } else {
          firstChild = nil
        }
        childNode.parent = nil
        // TODO: determine how to handle text
        // props.parent = nil
        childNode.next = nil
        childNode.prev = nil
        // Cleanup Yoga Node
        YGNodeFree(childNode.yogaNodeRef)
        // Update the nodes first child
        // Remove as parent
        updateChildren()
    }
    
    // TODO: Use array methods to find index of children for Yoga Compadability
    func insertBefore(_ element: RenderViewNode, _ anchor: RenderViewNode?) {
        // If no anchor set first child to view (make head)
        if let anchor = anchor,
           anchor === firstChild {
            anchor.prev = element
            element.next = anchor
            firstChild = element
        } else if let anchor = anchor,
                  let anchorPrevSibling = anchor.prev {
            anchorPrevSibling.next = element
            element.prev = anchorPrevSibling
            anchor.prev = element
            element.next = anchor
        } else if let firstChild = firstChild {
            // Make it at the end of the list
            var nextChild: RenderViewNode? = firstChild
            while let child = nextChild {
              nextChild = child.next

              if nextChild == nil {
                let lastChild = child
                lastChild.next = element
                element.prev = lastChild
              }
            }
        } else {
            firstChild = element
        }
        element.parent = self

        updateChildren()
    }
    
}

// MARK: Update Yoga Layout Style
extension RenderViewNode {
    
    
    func setYGNodeStyles(from jsValue: JSValue, for node: YGNodeRef) {
        func getFloat(from num: Any) -> Float {
            if let num = num as? NSNumber {
                return Float(truncating: num)
            }
            return 0
        }

        func getYGDirection(from direction: Any) -> YGDirection {
            switch direction as? String {
            case "LTR": return .LTR
            case "RTL": return .RTL
            case "inherit": return .inherit
            default: return .LTR
            }
        }

        func getYGFlexDirection(from direction: Any) -> YGFlexDirection {
            switch direction as? String {
            case "row": return .row
            case "row-reverse": return .rowReverse
            case "column-reverse": return .columnReverse
            default: return .column
            }
        }

        func getYGJustify(from justify: Any) -> YGJustify {
            switch justify as? String {
            case "center": return .center
            case "flex-end": return .flexEnd
            case "space-between": return .spaceBetween
            case "space-around": return .spaceAround
            case "space-evenly": return .spaceEvenly
            default: return .flexStart
            }
        }

        func getYGAlign(from align: Any) -> YGAlign {
            switch align as? String {
            case "center": return .center
            case "flex-end": return .flexEnd
            case "stretch": return .stretch
            case "baseline": return .baseline
            case "space-between": return .spaceBetween
            case "space-around": return .spaceAround
            default: return .auto
            }
        }

        func getYGPositionType(from positionType: Any) -> YGPositionType {
            // guard let positionType = jsValue.toString() else { return .relative }
            if let positionType = positionType as? String,
               positionType == "absolute" {
                return .absolute
            }
            return .relative
        }

        func getYGWrap(from wrap: Any) -> YGWrap {
            if let wrap = wrap as? String,
               wrap == "noWrap" {
                return .noWrap
            }
            return .wrap
        }

        func getYGOverflow(from overflow: Any) -> YGOverflow {
            switch overflow as? String {
            case "hidden": return .hidden
            case "scroll": return .scroll
            default: return .visible
            }
        }

        func getYGDisplay(from display: Any) -> YGDisplay {
            if let display = display as? String,
               display == "none" {
               return .none
            }
            return .flex
        }

        for (key, value) in jsValue.toDictionary() {
            // let value = jsValue
            
            switch key as? String {
            case "direction":
                YGNodeStyleSetDirection(node, getYGDirection(from: value))
            case "flexDirection":
                YGNodeStyleSetFlexDirection(node, getYGFlexDirection(from: value))
            case "justifyContent":
                YGNodeStyleSetJustifyContent(node, getYGJustify(from: value))
            case "alignContent":
                YGNodeStyleSetAlignContent(node, getYGAlign(from: value))
            case "alignItems":
                YGNodeStyleSetAlignItems(node, getYGAlign(from: value))
            case "alignSelf":
                YGNodeStyleSetAlignSelf(node, getYGAlign(from: value))
            case "positionType":
                YGNodeStyleSetPositionType(node, getYGPositionType(from: value))
            case "flexWrap":
                YGNodeStyleSetFlexWrap(node, getYGWrap(from: value))
            case "overflow":
                YGNodeStyleSetOverflow(node, getYGOverflow(from: value))
            case "display":
                YGNodeStyleSetDisplay(node, getYGDisplay(from: value))
            case "flex":
                YGNodeStyleSetFlex(node, getFloat(from: value))
            case "flexGrow":
                YGNodeStyleSetFlexGrow(node, getFloat(from: value))
            case "flexShrink":
                YGNodeStyleSetFlexShrink(node, getFloat(from: value))
            case "flexBasis":
                YGNodeStyleSetFlexBasis(node, getFloat(from: value))
            case "width":
                YGNodeStyleSetWidth(node, getFloat(from: value))
            case "height":
                YGNodeStyleSetHeight(node, getFloat(from: value))
            case "minWidth":
                YGNodeStyleSetMinWidth(node, getFloat(from: value))
            case "minHeight":
                YGNodeStyleSetMinHeight(node, getFloat(from: value))
                // YGNodeStyleSetMinWidth(node, )
                YGNodeCalculateLayout(node, YGUndefined, YGUndefined, .LTR)
            case "maxWidth":
                YGNodeStyleSetMaxWidth(node, getFloat(from: value))
            case "maxHeight":
                YGNodeStyleSetMaxHeight(node, getFloat(from: value))
            case "aspectRatio":
                YGNodeStyleSetAspectRatio(node, getFloat(from: value))
            default:
                continue
            }
        }
    }
    
    func updateNodeLayout() {
        // TODO: Pass layout on view wrapper
        viewWrapper.layoutMetrics = LayoutMetrics()
    }
}
