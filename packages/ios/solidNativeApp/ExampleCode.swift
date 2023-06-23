//
//  ExampleCode.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/21/23.
//

import Foundation


//sharedSolidNativeCore.registerElements()
//let root = sharedSolidNativeCore.createRootElement(name: "sn_v_stack_view")
//let button =  sharedSolidNativeCore.createRootElement(name: "sn_button_view")
//let text = sharedSolidNativeCore.createRootElement(name: "sn_text_view")
//rootElement = root
//
//var inc = 0
//
//root.insertBefore(element: button, anchor: nil)
//
//root.insertBefore(element: text, anchor: button)
//
//
//text.setProp("text", "Count: 0")
//
//
//button.setProp("title", "Increment!")
//
//button.setProp("onPress", {
//    inc += 1
//    text.setProp("text", "Count \(inc)")
//})
//
//var toggle = true
//
//root.setProp("value", toggle)
//
//root.setProp("onChange") {
//    (_ newValue: Bool) -> Void in
//    toggle = newValue
//    root.setProp("value", newValue)
//}
//
//Timer.scheduledTimer(withTimeInterval: 1.0, repeats: true) { _ in
//    toggle = !toggle
//    root.setProp("value", toggle)
//}
//
//let jsPrint: @convention(block) (_ contents: String) -> Void = { str in
//    print(str)
//}
//
//jsContext.setObject(jsPrint, forKeyedSubscript: "print" as NSString)
//
//Task {
//    if let url = URL(string: "http://localhost:8080") {
//        do {
//            let (data, res) = try await URLSession.shared.data(from: url)
//
//            guard let res = res as? HTTPURLResponse,
//                  res.statusCode == 200 /* OK */ else {
//                throw URLError(.badServerResponse)
//            }
//
//            let str = String(decoding: data, as: UTF8.self)
//
//
//
//            print("Result: " + jsContext.evaluateScript(str).toString()!)
//
//        } catch {
//            // contents could not be loaded
//            print("Url was bad!")
//        }
//    } else {
//        // the URL was bad!
//        print("ERROR: Url was bad!")
//    }
//}
//
//
//print("JSValue: " + jsContext.evaluateScript("const global = this; Object.keys(global)").toString()!)
