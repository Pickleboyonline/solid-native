//
//  File.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 12/31/24.
//
import Foundation
import Snmobile
import SwiftUI

extension SNSnmobileStringArray: RandomAccessCollection {
    public typealias Element = SolidNativeViewWrapper
    public typealias Index = Int
    
    public var startIndex: Index {
        return 0
    }
    
    public var endIndex: Index {
        return length()
    }
    
    public subscript(position: Index) -> Element {
        let nodeId = get(position)
        return SharedSolidNativeCore.viewWrapperRegistry[nodeId]!
    }
    
    public func index(after i: Index) -> Index {
        return i + 1
    }
    
    public func index(before i: Index) -> Index {
        return i - 1
    }
}
