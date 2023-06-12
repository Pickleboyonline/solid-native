//
//  SolidNativeProps.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/11/23.
//

import Foundation


// Not meant to be exteneded
class SolidNativeProps: ObservableObject {
    @Published var updateCount = 0;
    @Published var values: [String:Any] = [:];
    @Published var text = ""
    
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
    
    func getChildren() -> [SolidNativeElement] {
        if let values = values["children"] as? [SolidNativeElement] {
            return values
        }
        return []
    }
}
