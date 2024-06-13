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


/// Temp protocol for eventual implimention
/// Prob would have to be a class, because we need to get the definitions of them all
/// Need to support "refs" basically js representation of class
/// Callbacks are just refs
protocol SolidNativeViewP {
    var props: [String: Any] {get}
    static var name: String {get}
     var isTextElement: Bool {get}
    associatedtype V: View
    func render() -> V
    init()
}

extension SolidNativeViewP {
    var isTextElement: Bool {false}
}

class BaseSolidNativeView: SolidNativeViewP {
    var props: [String : Any] = [:]
    
    required init() {
        
    }
    static var name: String {"BaseSolidNativeView"}
    
    func render() -> some View {
        HStack {
            Text("Hi")
        }
    }
}


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
    
    // MARK: View State Management
    // TODO: Determine how props are handled
    var props: [String: JSValue] = [:]
    // TODO: Will prob have to be a class eventually for Rust/C++ interop
    let view: any SolidNativeViewP
    // TODO: Determine what effect this has on SolidJS
    var isTextElement: Bool {view.isTextElement}
    
    init(view: any SolidNativeViewP) {
        self.view = view
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
    func setProp(_ name: String, _ value: Any) {
        
    }
    func removeProp(_ name: String) {
        
    }
    
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
    func updateYogaNodeStyle() {
        
    }
    
    func updateNodeLayout() {
        
    }
}
