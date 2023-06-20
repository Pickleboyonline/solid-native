//
//  SolidNativeView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore
import SwiftUI

protocol AnySolidNativeElementJSExport: JSExport {
    // Expose methods needed.
}

/**
 Later, we can have it to where you only need to pass in the SwiftUI View that takes in props.
 Contains a a solid native view
 */
@objc public class AnySolidNativeElement: NSObject, AnySolidNativeElementJSExport {
    var next: AnySolidNativeElement?
    var prev: AnySolidNativeElement?
    let id = UUID()
    
    class var name: String {
        "any_element"
    }
    
    required override init() {
    }
    
    
    let props = SolidNativeProps()
    
    func setProp(_ name: String, _ value: Any) {
        assert(name != "children", "Err: User `removeChild` or `insertBefore` to update children!")
        props.values[name] = value
    }
    
    private func _setProp(_ name: String, _ value: Any) {
        props.values[name] = value
    }
    
    var children: [AnySolidNativeElement] = []
    
    // Can be getter
    var firstChild: AnySolidNativeElement?
    
    var parentElement: AnySolidNativeElement?
    
    // Iterate over first child prop
    // O(n)
    private func updateChildrenInProps() {
        var newChildren: [AnySolidNativeElement] = []
        
        var nextChild = firstChild
        while let child = nextChild {
            newChildren.append(child)
            nextChild = child.next
        }
        
        children = newChildren
        _setProp("children", newChildren)
    }
   
    // O(1)
    func removeChild(element: AnySolidNativeElement) {
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
        }
        element.parentElement = nil
        element.next = nil
        element.prev = nil
        // Update the nodes first child
        // Remove as parent
        updateChildrenInProps()
    }
    
    // O(1)
    func insertBefore(element: AnySolidNativeElement, anchor: AnySolidNativeElement?) {
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
            firstChild.prev = element
            element.next = firstChild
            self.firstChild = element
        } else {
            firstChild = element
        }
        
        element.parentElement = self
    
        updateChildrenInProps()
    }
    
    // TODO: You need to override this!
    // This is how SolidJS will deliver a text prop.
    var isTextElement: Bool {
        false
    }
    
    func setText(text: String) {
        props.text = text
    }
    
    
    
    func render() -> AnyView {
        AnyView(EmptyView())
    }
    // @ViewBuilder
//    func render() -> some View {
//        // view.init(isTextView: <#Bool#>, props: <#SolidNativeProps#>)
//        Text("")
//    }
}

class SolidNativeModule {
    // You can basically do anything here
}


//protocol AnySolidNativeView {
//
//}
