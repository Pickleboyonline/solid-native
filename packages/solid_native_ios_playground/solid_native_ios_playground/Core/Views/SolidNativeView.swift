//
//  SolidNativeView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI
import JavaScriptCore

typealias TOnRef<T> = (T) -> Void

@objc protocol RefJSExport: JSExport {
    
}

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

struct SolidNativeView: View {
    // Should provide implimentation for this
    @ObservedObject var props: SolidNativeProps
    
    @StateObject var ref = Ref()
    
    func getOnRef() -> TOnRef<Ref> {
        let onRef: TOnRef<Ref>  = props.getProp(name: "onRef", default: {
            (_ r: Ref) -> Void in
        })
        return onRef
    }

    var body: some View {
        // Could be cleaner
        let message: String = props.getProp(name: "message", default: "")
        VStack {
            Text("Update Count: \(message)")
            Button("Inc") {
            }
        }.onAppear() {
            ref.setRef(getOnRef())
        }
    }
}
