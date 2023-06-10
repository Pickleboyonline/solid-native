//
//  SNProps.swift
//  solid_native_core_ios
//
//  Created by Imran Shitta-Bey on 6/9/23.
//  Copyright © 2023 orgName. All rights reserved.
//

import Foundation


class SNProps: ObservableObject {
    @Published var data: [String: Any?] = [:]
    
    func setProp(name: String, value: Any) {
        data[name] = value
    }
    
    // TODO: Create typed versions of this.
    func getProp<T>(name: String, defaultValue: T) -> T  {
        let value = data[name] as? T
        return value ?? defaultValue
    }
}
