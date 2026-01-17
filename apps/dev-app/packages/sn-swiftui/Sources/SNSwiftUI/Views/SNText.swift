//
//  SNText.swift
//  SNSwiftUI
//
//  Text component with comprehensive styling support.
//

import Foundation
import SwiftUI
import SNCore

class SNText: SolidNativeView {
    static var name: String { "sn_text" }
    static var isTextElement: Bool { true }

    func render() -> some View {
        // Try to use text descriptors first (for nested text)
        if let descriptors = wrapper.textDescriptors, !descriptors.isEmpty {
            return AnyView(renderFromDescriptors(descriptors))
        }

        // Fallback to simple text prop
        if let textValue = props["text"], case .string(let text) = textValue {
            var textView = Text(text)

            if let style = props["style"], case .object(let styleProps) = style {
                // Apply text transform first
                if let transform = styleProps["textTransform"], case .string(let transformStr) = transform {
                    switch transformStr {
                    case "uppercase":
                        textView = Text(text.uppercased())
                    case "lowercase":
                        textView = Text(text.lowercased())
                    case "capitalize":
                        textView = Text(text.capitalized)
                    default:
                        break
                    }
                }

                textView = applyTextStyles(textView, styleProps: styleProps)
            }

            return AnyView(textView)
        }

        return AnyView(EmptyView())
    }

    private func renderFromDescriptors(_ descriptors: [TextDescriptor]) -> some View {
        var result = Text("")

        for descriptor in descriptors {
            var segment = Text(descriptor.text)

            // Apply styles from the descriptor
            segment = applyTextStyles(segment, styleProps: descriptor.styles)

            result = result + segment
        }

        return result
    }

    private func applyTextStyles(_ text: Text, styleProps: [String: JsValue]) -> Text {
        var styledText = text

        // Font
        styledText = styledText.font(textFont(from: styleProps))

        // Color
        if let colorVal = styleProps["color"], case .string(let colorHex) = colorVal {
            styledText = styledText.foregroundColor(Color(hex: colorHex))
        }

        // Letter spacing / kerning
        if let letterSpacing = styleProps["letterSpacing"], case .number(let spacing) = letterSpacing {
            styledText = styledText.kerning(CGFloat(spacing))
        }

        // Text decoration
        if let decoration = styleProps["textDecorationLine"], case .string(let decorationStr) = decoration {
            switch decorationStr {
            case "underline":
                styledText = styledText.underline()
            case "line-through":
                styledText = styledText.strikethrough()
            case "underline line-through":
                styledText = styledText.underline().strikethrough()
            default:
                break
            }
        }

        return styledText
    }

    private func textFont(from styleProps: [String: JsValue]) -> Font {
        var fontSize: CGFloat = 14
        var fontFamily: String?
        var fontWeight: Font.Weight = .regular
        var isItalic = false

        // Font size
        if let sizeVal = styleProps["fontSize"], case .number(let size) = sizeVal {
            fontSize = CGFloat(size)
        }

        // Font family
        if let familyVal = styleProps["fontFamily"], case .string(let family) = familyVal {
            fontFamily = family
        }

        // Font weight
        if let weightVal = styleProps["fontWeight"], case .string(let weight) = weightVal {
            switch weight {
            case "bold": fontWeight = .bold
            case "100": fontWeight = .ultraLight
            case "200": fontWeight = .thin
            case "300": fontWeight = .light
            case "400": fontWeight = .regular
            case "500": fontWeight = .medium
            case "600": fontWeight = .semibold
            case "700": fontWeight = .bold
            case "800": fontWeight = .heavy
            case "900": fontWeight = .black
            default: fontWeight = .regular
            }
        }

        // Font style (italic)
        if let styleVal = styleProps["fontStyle"], case .string(let style) = styleVal {
            isItalic = style == "italic"
        }

        // Build the font
        var font: Font
        if let family = fontFamily {
            font = Font.custom(family, size: fontSize)
        } else {
            font = Font.system(size: fontSize)
        }

        font = font.weight(fontWeight)

        if isItalic {
            font = font.italic()
        }

        return font
    }
}
