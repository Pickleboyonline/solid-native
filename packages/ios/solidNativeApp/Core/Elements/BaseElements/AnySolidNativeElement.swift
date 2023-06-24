//
//  SolidNativeView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore
import SwiftUI

@objc protocol AnySolidNativeElementJSExport: JSExport {
    // Expose methods needed.
    var firstChild: AnySolidNativeElement? { get }
    var parentElement: AnySolidNativeElement? {get}
    func setProp(_ name: String, _ value: JSValue?)
    var isTextElement: Bool {get}
    func removeChild(_ element: AnySolidNativeElement)
    func insertBefore(_ element: AnySolidNativeElement, _ anchor: AnySolidNativeElement?)
    var next: AnySolidNativeElement? {get}
    var prev: AnySolidNativeElement? {get}
}


/**
 Later, we can have it to where you only need to pass in the SwiftUI View that takes in props.
 Contains a a solid native view
 */
@objc public class AnySolidNativeElement: NSObject, AnySolidNativeElementJSExport {
    dynamic var next: AnySolidNativeElement?
    dynamic var prev: AnySolidNativeElement?
    let id = UUID()
    
    class var name: String {
        "any_element"
    }
    
    required override init() {
    }
    
    
    let props = SolidNativeProps()
    
    @objc public func setProp(_ name: String, _ value: JSValue?) {
        // print("JS value type: " + String(value!.isString))
        assert(name != "children", "Err: User `removeChild` or `insertBefore` to update children!")
        props.values[name] = value
    }
    
    
    var children: [AnySolidNativeElement] = []
    
    // Can be getter
    dynamic var firstChild: AnySolidNativeElement?
    
    dynamic var parentElement: AnySolidNativeElement?
    
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
        props.children = newChildren
    }
    
   
    // O(1)
    func removeChild(_ element: AnySolidNativeElement) {
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
    func insertBefore(_ element: AnySolidNativeElement, _ anchor: AnySolidNativeElement?) {
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
            var nextChild: AnySolidNativeElement? = firstChild
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

class SolidNativeModule {
    // You can basically do anything here
}


//protocol AnySolidNativeView {
//
//}
