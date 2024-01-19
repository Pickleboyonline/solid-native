//
//  JSValueBuilder.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/19/24.
//

import Foundation
import JavaScriptCore


class JSValueBuilder {
    let value = JSValue(newObjectIn: SolidNativeCore.shared.jsContext)!
    
    // Helper function to process JSValue to Swift type
    private func processValue<T>(_ jsValue: JSValue, asType type: T.Type) -> T? {
        switch jsValue {
        case let value where value.isString:
            return value.toString() as? T
        case let value where value.isNumber:
            return value.toNumber() as? T
        case let value where value.isBoolean:
            return value.toBool() as? T
        default:
            return value as? T
        }
    }
    
    // No arguments
    func addSyncFunction<U>(_ name: String, fn: @escaping () -> U?) {
        let objcFunc: @convention(block) (JSValue) -> JSValue? = { jsValue in
            let result = fn()
            print(result ?? "nil", name)
            return JSValue(object: result, in: SolidNativeCore.shared.jsContext)
            
        }
        value.setObject(objcFunc, forKeyedSubscript: name as NSString)
    }
    
    // Single argument
    func addSyncFunction<T, U>(_ name: String, fn: @escaping (T) -> U?) {
        let objcFunc: @convention(block) (JSValue) -> JSValue? = { jsValue in
            guard let arg1 = self.processValue(jsValue, asType: T.self) else {
                return nil
            }
            
            if let result = fn(arg1) {
                return JSValue(object: result, in: SolidNativeCore.shared.jsContext)
            }
            return nil
        }
        value.setObject(objcFunc, forKeyedSubscript: name as NSString)
    }
    
    // Two arguments
    func addSyncFunction<T, U, V>(_ name: String, fn: @escaping (T, U) -> V?) {
        let objcFunc: @convention(block) (JSValue, JSValue) -> JSValue? = { jsValue1, jsValue2 in
            guard let arg1 = self.processValue(jsValue1, asType: T.self),
                  let arg2 = self.processValue(jsValue2, asType: U.self) else {
                return nil
            }
            
            if let result = fn(arg1, arg2) {
                return JSValue(object: result, in: SolidNativeCore.shared.jsContext)
            }
            return nil
        }
        value.setObject(objcFunc, forKeyedSubscript: name as NSString)
    }
    
    // Three arguments
    func addSyncFunction<T, U, V, W>(_ name: String, fn: @escaping (T, U, V) -> W?) {
        let objcFunc: @convention(block) (JSValue, JSValue, JSValue) -> JSValue? = { jsValue1, jsValue2, jsValue3 in
            guard let arg1 = self.processValue(jsValue1, asType: T.self),
                  let arg2 = self.processValue(jsValue2, asType: U.self),
                  let arg3 = self.processValue(jsValue3, asType: V.self) else {
                return nil
            }
            
            if let result = fn(arg1, arg2, arg3) {
                return JSValue(object: result, in: SolidNativeCore.shared.jsContext)
            }
            return nil
        }
        value.setObject(objcFunc, forKeyedSubscript: name as NSString)
    }
    
    // Four arguments
    func addSyncFunction<T, U, V, W, X>(_ name: String, fn: @escaping (T, U, V, W) -> X?) {
        let objcFunc: @convention(block) (JSValue, JSValue, JSValue, JSValue) -> JSValue? = { jsValue1, jsValue2, jsValue3, jsValue4 in
            guard let arg1 = self.processValue(jsValue1, asType: T.self),
                  let arg2 = self.processValue(jsValue2, asType: U.self),
                  let arg3 = self.processValue(jsValue3, asType: V.self),
                  let arg4 = self.processValue(jsValue4, asType: W.self) else {
                return nil
            }
            
            if let result = fn(arg1, arg2, arg3, arg4) {
                return JSValue(object: result, in: SolidNativeCore.shared.jsContext)
            }
            return nil
        }
        value.setObject(objcFunc, forKeyedSubscript: name as NSString)
    }
    
    func addProperty(_ name: String, prop: Any) {
        value.setObject(prop, forKeyedSubscript: name as NSString)
    }
}
