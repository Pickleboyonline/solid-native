////
////  NewSolidNativeCore.swift
////  SolidNative
////
////  Created by Imran Shitta-Bey on 6/12/24.
////
//
//import Foundation
//import JavaScriptCore
//
///// Need to:
///// -
/////
//class NewSolidNativeCore {
//
//    let jsContext = JSContext()!
//
//    let moduleRegistry: [String: SolidNativeView] = [:]
//
//    
//    func injectRuntimeIntoJSContext() {
//        // Place timers, console.log, core lib etc.. in here.
//        // Core must have:
//        //  - get native modules
//        //  - render functionality (for now)
//        //  - potentially function to restart lol
//    }
//
//    func downloadAndRunJsBundleSync() {
//        
//    }
//    
//    func downloadAndRunJsBundleAsync() {
//        
//    }
//    
//    // Need to grab root render function for system
//    
//    
//}
//
//class SolidNativeRenderer {
//    // TODO: Expose methods for class
//    
//    // TODO: Somehow, we need a registry to know what class to make.
//    
//    // There needs to be a map of view name => class def
//    
//    // We need way to manage layout shadow tree
//    // Also each node has an associated yoga layout node to caclulate the nodes.
//    
//    //
//}
//
//class _SolidNativeShadowView {
//    // Graph structure.
//    
//    // Contains render function. Takes in view def
//    // (must have protocal to accept props
//}
//
