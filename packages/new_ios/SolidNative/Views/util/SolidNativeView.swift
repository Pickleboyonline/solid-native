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

class SolidNativeView {
  dynamic var next: SolidNativeView?
  dynamic var prev: SolidNativeView?
  let id = UUID()
    

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
      let childPrevSibling = element.prev
    {
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
          lastChild.next = element
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
