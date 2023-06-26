//
//  SNTextView.swift
//  solid_native_ios_playground
//
//  Created by Imran Shitta-Bey on 6/12/23.
//

import Foundation
import SwiftUI
import JavaScriptCore



class SNTextElement: AnySolidNativeElement {
    
    class override var name: String {
        "sn_text"
    }
    
    override var isTextElement: Bool {
        true
    }
    
    struct SNTextView: View {
        
        @ObservedObject var props: SolidNativeProps
        
        
        typealias Modifyer = (_ view: Text, _ value: JSValue?) -> any View
        
        typealias TAcceptedProps = [String : Modifyer]
        
        func applyModifiers(@ViewBuilder _ content:  () -> Text) -> Text {
            // Essentially, every prop will have a value and data. Depending
            // on this information, we execute a function based on its value and pass in that view
            // Thus, we need a subscript like class to execute based on the accepted view modifiers
            // Data values are always wrapped in JS file
            var startingView = content()
            for (name, jsValue) in props.values {
                // TODO: Use has map
                let acceptedProps: TAcceptedProps = [
                    "fontWeight": applyFontWeight
                ]
                if let applier = acceptedProps[name] ?? nil {
                    startingView = applier(startingView, jsValue) as! Text
                }
            }
            return startingView
        }
        
        func applyFontWeight(_ view: Text, _ value: JSValue?) -> Text {
            view.fontWeight(toFontWeight(value?.toString()))
        }
        
        func applyFont(_ view: Text, _ value: JSValue?) -> Text {
            view.font(.body)
        }
        
        var body: some View {
            // Our base view will take care of
            applyModifiers {
                let text = props.getString(name: "text")
                Text(text)
            }
        }
        
        func toFontWeight(_ fontWeight: String?) -> Font.Weight {
          switch fontWeight {
          case "ultraLight":
            return .ultraLight
          case "thin":
            return .thin
          case "light":
            return .light
          case "regular":
            return .regular
          case "medium":
            return .medium
          case "semibold":
            return .semibold
          case "bold":
            return .bold
          case "heavy":
            return .heavy
          case "black":
            return .black
          default:
            return .regular
          }
        }
    }
    
    override func render() -> AnyView
    {
        AnyView(SNTextView(props: self.props))
    }
    
}
