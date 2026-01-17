//
//  SNImage.swift
//  SNSwiftUI
//
//  Image component with URL loading via AsyncImage.
//

import Foundation
import SwiftUI
import SNCore

class SNImage: SolidNativeView {
    static var name: String { "sn_image" }

    func render() -> some View {
        guard let source = props["source"],
              case .object(let sourceProps) = source,
              let uri = sourceProps["uri"],
              case .string(let urlString) = uri,
              let url = URL(string: urlString) else {
            return AnyView(EmptyView())
        }

        return AnyView(
            AsyncImage(url: url) { phase in
                switch phase {
                case .success(let image):
                    image
                        .resizable()
                        .aspectRatio(contentMode: self.contentMode)
                case .failure:
                    Color.gray.opacity(0.3)
                case .empty:
                    ProgressView()
                @unknown default:
                    EmptyView()
                }
            }
        )
    }

    private var contentMode: ContentMode {
        // Check top-level resizeMode first
        if let mode = props["resizeMode"], case .string(let modeStr) = mode {
            return parseResizeMode(modeStr)
        }

        // Check style.resizeMode
        if let style = props["style"], case .object(let styleProps) = style {
            if let mode = styleProps["resizeMode"], case .string(let modeStr) = mode {
                return parseResizeMode(modeStr)
            }
        }

        return .fit
    }

    private func parseResizeMode(_ mode: String) -> ContentMode {
        switch mode {
        case "cover": return .fill
        case "contain", "center", "stretch": return .fit
        default: return .fit
        }
    }
}
