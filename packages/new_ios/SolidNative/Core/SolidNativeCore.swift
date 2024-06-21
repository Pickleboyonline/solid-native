//
//  SolidNativeCore.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/20/24.
//

import Foundation
import Snmobile
import UIKit

var SharedSolidNativeCore: SolidNativeCore!

@objc public class SolidNativeCore: NSObject, SNSnmobileHostReceiverProtocol {

    public override init() {
        super.init()
        if SharedSolidNativeCore == nil {
            SharedSolidNativeCore = self
        }
    }
    
    var snmobile: SNSnmobileSolidNativeMobile!
    
    var viewWrapperRegistry: [String:SolidNativeViewWrapper] = [:]
    
    let viewTypeRegistry: [String: any SolidNativeView.Type] = [
        SNView.name:SNView.self
    ]
    
    var rootNodeId: String?
    
    func start(jsUrl: String) throws {
        snmobile = SNSnmobileSolidNativeMobile(self)!
        
        let nodeId = snmobile.createRootNode("sn_view")
        
        viewWrapperRegistry[nodeId] = SolidNativeViewWrapper(viewType: SNView.self)
        
        rootNodeId = nodeId
        
        snmobile.registureModules()
        
        try snmobile.runJs(fromServer: jsUrl)
        
        
        
    }
    
    func getRootNode() -> SolidNativeViewWrapper {
        viewWrapperRegistry[rootNodeId!]!
    }
    
}


// Conform to Protocol
extension SolidNativeCore {
    public func doesNodeRequireMeasuring(_ nodeType: String?) -> Bool {
        viewTypeRegistry[nodeType!]!.doesRequireMeasuring
    }
    
    public func getDeviceScreenSize() -> SNSnmobileSize? {
        // Get the width and height of the device screen
        let screenWidth = UIScreen.main.bounds.width
        let screenHeight = UIScreen.main.bounds.height
        return SNSnmobileSize(Float(screenWidth), height: Float(screenHeight))
    }
    
    public func isTextElement(_ nodeId: String?) -> Bool {
        viewWrapperRegistry[nodeId!]!.solidNativeViewType.isTextElement
    }
    
    /// TODO: Impliment
    public func measureNode(_ nodeId: String?) -> SNSnmobileSize? {
        return SNSnmobileSize(0, height: 0)
    }
    
    public func onChildrenChange(_ nodeId: String?, nodeIds: SNSnmobileStringArray?) {
        print("Node Ids: \(String(describing: nodeIds?.length()))")
        viewWrapperRegistry[nodeId!]!.children = nodeIds!
    }
    
    public func onLayoutChange(_ nodeId: String?, layoutMetrics: SNSnmobileLayoutMetrics?) {
        viewWrapperRegistry[nodeId!]!.layoutMetrics = layoutMetrics!
    }
    
    public func onNodeCreated(_ nodeId: String?, nodeType: String?) {
        let viewType = viewTypeRegistry[nodeType!]!
        
        viewWrapperRegistry[nodeId!] = SolidNativeViewWrapper(viewType: viewType)
    }
    
    public func onPropUpdated(_ nodeId: String?, key: String?, value: SNSnmobileJSValue?) {
        viewWrapperRegistry[nodeId!]!.props[key!] = value!
    }
    
    public func onUpdateRevisionCount(_ nodeId: String?) {
        viewWrapperRegistry[nodeId!]!.updateRevisionCount()
    }
    
}
