//
//  SolidNativeApp.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/20/24.
//

import SwiftUI
import SNLib

@main
struct SolidNativeApp: App {
    
    let hostReceiver: HostReceiver
    
    init() {
        let core = SNCoreCore()!
        
        hostReceiver = HostReceiver()
        
        let renderer = SNRendererRenderer(hostReceiver)!
        
        let nodeId = renderer.createRootNode(SNView.name)
        hostReceiver.onNodeCreated(nodeId, nodeType: SNView.name)
        hostReceiver.rootNodeId = nodeId
        
        core.registerGo(renderer)
        let url = "http://localhost:8080/"
        do {
//            try core.start(fromJS: """
//            console.log("Hello World from QuickJS!")
//            console.log("Hello World from QuickJS!")
//            """)
            try core.start(fromServer: url)
        } catch {
            print("Unexpected error: \(error).")
        }
    }
    
    var body: some Scene {
        WindowGroup {
            hostReceiver.getRootNode().render()
        }
    }
}
