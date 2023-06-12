//
//  SolidNativeView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation
import JavaScriptCore
import SwiftUI

// Not meant to be exteneded
class Props: ObservableObject {
    @Published var updateCount = 0;
    @Published var values: [String:Any] = [:];
    
    // TODO: Figure out how to deal with events and refs
    
    func getInt(_ name: String) -> Int? {
        values[name] as? Int
    }
    
    func getProp<T>(name: String, `default`: T) -> T {
        if let prop = values[name] as? T {
            return prop
        }
        return `default`
    }
    
    func getChildren() -> [SolidNativeView] {
        if let values = values["children"] as? [SolidNativeView] {
            return values
        }
        return []
    }
}

typealias TOnRef<T> = (T) -> Void

/**
 Internalized state thats accessable
 Meant to be extended
 Make JSExport.
 TODO: If using hermes, maybe consider something else
 TODO: Expose methods as needed
 */
@objc public class Ref: NSObject, ObservableObject  {
    @Published var updateCount = 0;
    
    func setRef(_ onRef: TOnRef<Ref>) {
        // Needs to be JS Serialized
        onRef(self)
    }
}

struct SolidRegularView: View {
    // Should provide implimentation for this
    @ObservedObject var props: Props
    
    @StateObject var ref = Ref()
    
    func getOnRef() -> TOnRef<Ref> {
        let onRef: TOnRef<Ref>  = props.getProp(name: "onRef", default: {
            (_ r: Ref) -> Void in
        })
        return onRef
    }

    // You're free!
    var body: some View {
        let message: String = props.getProp(name: "message", default: "")
        VStack {
            Text("Update Count: \(message)")
            Button("Inc") {
                // props.triggerUpdate()
            }
        }.onAppear() {
            ref.setRef(getOnRef())
        }
    }
}


protocol SolidNativeViewJSExport: JSExport {
    // Expose methods needed.
}

/**
 Later, we can have it to where you only need to pass in the SwiftUI View that takes in props.
 */
@objc public class SolidNativeView: NSObject, ObservableObject, SolidNativeViewJSExport {
    var next: SolidNativeView?
    var prev: SolidNativeView?
    var parent: SolidNativeView?
    
    var firstChild: SolidNativeView?
    var childrenCount = 0
    
    private let props = Props()
    
    func setProp(_ name: String, _ value: Any) {
        assert(name != "children", "Err: User `removeChild` or `insertBefore` to update children!")
        props.values[name] = value
    }
    
    private func _setProp(_ name: String, _ value: Any) {
        props.values[name] = value
    }
    
    var children: [SolidNativeCore] = []
    
    // Iterate over first child prop
    // O(n)
    private func updateChildren() {
        // Make new array of childrenCount;
        
        //        for i in 0...childrenCount {
        //
        //        }
        _setProp("children", self.children)
    }
   
    // O(1)
    func removeChild() {
        // Link the nodes prev and next of it
        // Update the nodes first child
        // Remove as parent
        childrenCount -= 1
        updateChildren()
    }
    
    // O(1)
    func insertBefore(view: SolidNativeView, anchor: SolidNativeView?) {
        // If no anchor set first child to view (make head)
        //
        
        childrenCount += 1
        updateChildren()
    }
    
    @ViewBuilder
    func render() -> some View {
        SolidRegularView(props: props)
    }
}



class SolidNativeModule {
    // You can basically do anything here
}


//protocol AnySolidNativeView {
//
//}
