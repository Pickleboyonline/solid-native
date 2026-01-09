//
//  SNSwiftUIRenderer.swift
//  SNSwiftUI
//
//  Renders a ViewNode tree as SwiftUI views
//

import SwiftUI

/// Main renderer view that displays a ViewNode tree
public struct SNSwiftUIRenderer: View {
    @ObservedObject private var viewTree: ViewTree

    public init(viewTree: ViewTree) {
        self.viewTree = viewTree
    }

    public var body: some View {
        Group {
            if let root = viewTree.root {
                NodeView(node: root)
            } else {
                Text("No content")
                    .foregroundColor(.gray)
            }
        }
    }
}

/// Renders a single ViewNode and its children
struct NodeView: View {
    @ObservedObject var node: ViewNode

    var body: some View {
        Group {
            if node.isTextNode {
                Text(node.textContent ?? "")
            } else {
                renderElement()
            }
        }
    }

    @ViewBuilder
    private func renderElement() -> some View {
        switch node.type.lowercased() {
        case "view", "div", "container":
            renderContainer()

        case "vstack":
            VStack(spacing: parseSpacing()) {
                renderChildren()
            }
            .applyModifiers(from: node)

        case "hstack":
            HStack(spacing: parseSpacing()) {
                renderChildren()
            }
            .applyModifiers(from: node)

        case "zstack":
            ZStack {
                renderChildren()
            }
            .applyModifiers(from: node)

        case "scrollview":
            ScrollView {
                renderChildren()
            }
            .applyModifiers(from: node)

        case "button":
            Button(action: handleButtonPress) {
                renderChildren()
            }
            .applyModifiers(from: node)

        case "image":
            renderImage()

        case "spacer":
            Spacer()

        default:
            // Unknown element type - render as a VStack container
            VStack {
                renderChildren()
            }
            .applyModifiers(from: node)
        }
    }

    @ViewBuilder
    private func renderContainer() -> some View {
        VStack(spacing: 0) {
            renderChildren()
        }
        .applyModifiers(from: node)
    }

    @ViewBuilder
    private func renderChildren() -> some View {
        ForEach(node.children) { child in
            NodeView(node: child)
        }
    }

    @ViewBuilder
    private func renderImage() -> some View {
        if let imageName = node.properties["src"] ?? node.properties["name"] {
            Image(systemName: imageName)
                .resizable()
                .applyModifiers(from: node)
        } else {
            EmptyView()
        }
    }

    private func parseSpacing() -> CGFloat? {
        guard let spacingStr = node.properties["spacing"] else {
            return nil
        }
        return Double(spacingStr).map { CGFloat($0) }
    }

    private func handleButtonPress() {
        print("[Button] Pressed: \(node.id)")
        // TODO: Call back to Rust to handle the event
    }
}

// MARK: - View Modifier Extensions

extension View {
    @ViewBuilder
    func applyModifiers(from node: ViewNode) -> some View {
        self
            .applyPadding(from: node)
            .applyBackground(from: node)
            .applyForegroundColor(from: node)
            .applyCornerRadius(from: node)
            .applyFrame(from: node)
    }

    @ViewBuilder
    private func applyPadding(from node: ViewNode) -> some View {
        if let paddingStr = node.properties["padding"],
           let padding = Double(paddingStr) {
            self.padding(CGFloat(padding))
        } else {
            self
        }
    }

    @ViewBuilder
    private func applyBackground(from node: ViewNode) -> some View {
        if let colorStr = node.properties["backgroundColor"] {
            self.background(parseColor(colorStr))
        } else {
            self
        }
    }

    @ViewBuilder
    private func applyForegroundColor(from node: ViewNode) -> some View {
        if let colorStr = node.properties["color"] ?? node.properties["foregroundColor"] {
            self.foregroundColor(parseColor(colorStr))
        } else {
            self
        }
    }

    @ViewBuilder
    private func applyCornerRadius(from node: ViewNode) -> some View {
        if let radiusStr = node.properties["cornerRadius"],
           let radius = Double(radiusStr) {
            self.cornerRadius(CGFloat(radius))
        } else {
            self
        }
    }

    @ViewBuilder
    private func applyFrame(from node: ViewNode) -> some View {
        let width = node.properties["width"].flatMap { Double($0) }.map { CGFloat($0) }
        let height = node.properties["height"].flatMap { Double($0) }.map { CGFloat($0) }

        if width != nil || height != nil {
            self.frame(width: width, height: height)
        } else {
            self
        }
    }

    private func parseColor(_ colorStr: String) -> Color {
        switch colorStr.lowercased() {
        case "red": return .red
        case "blue": return .blue
        case "green": return .green
        case "yellow": return .yellow
        case "orange": return .orange
        case "purple": return .purple
        case "pink": return .pink
        case "gray", "grey": return .gray
        case "black": return .black
        case "white": return .white
        case "clear": return .clear
        default:
            // Try to parse hex color
            if colorStr.hasPrefix("#") {
                return parseHexColor(colorStr) ?? .primary
            }
            return .primary
        }
    }

    private func parseHexColor(_ hex: String) -> Color? {
        let hex = hex.trimmingCharacters(in: CharacterSet.alphanumerics.inverted)
        var int: UInt64 = 0
        Scanner(string: hex).scanHexInt64(&int)
        let a, r, g, b: UInt64
        switch hex.count {
        case 3: // RGB (12-bit)
            (a, r, g, b) = (255, (int >> 8) * 17, (int >> 4 & 0xF) * 17, (int & 0xF) * 17)
        case 6: // RGB (24-bit)
            (a, r, g, b) = (255, int >> 16, int >> 8 & 0xFF, int & 0xFF)
        case 8: // ARGB (32-bit)
            (a, r, g, b) = (int >> 24, int >> 16 & 0xFF, int >> 8 & 0xFF, int & 0xFF)
        default:
            return nil
        }
        return Color(
            .sRGB,
            red: Double(r) / 255,
            green: Double(g) / 255,
            blue:  Double(b) / 255,
            opacity: Double(a) / 255
        )
    }
}
