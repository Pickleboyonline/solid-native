//
//  File.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 12/31/24.
//
import Foundation
import SNLib
import SwiftUI

extension SNCoreStringArray: RandomAccessCollection {
//    public typealias Element = SNCoreStringArray
//    public typealias Index = Int
    
    public var startIndex: Int {
        return 0
    }
    
    public var endIndex: Index {
        return length()
    }
    
    public subscript(position: Int) -> String {
        return get(position)
    }
    
    public func index(after i: Index) -> Index {
        return i + 1
    }
    
    public func index(before i: Index) -> Index {
        return i - 1
    }
}
