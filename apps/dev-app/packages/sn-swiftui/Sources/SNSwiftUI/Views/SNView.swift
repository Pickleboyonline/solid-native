//
//  SNView.swift
//  SNSwiftUI
//
//  Container view component that renders children using Yoga-SwiftUI Flex layout.
//

import Foundation
import SwiftUI
import YogaSwiftUI
import yoga
import SNCore

class SNView: SolidNativeView {
    static var name: String { "sn_view" }

    func render() -> some View {
        let flexProps = extractFlexProps(from: props)

        return Flex(
            direction: flexProps.direction,
            justifyContent: flexProps.justifyContent,
            alignItems: flexProps.alignItems,
            wrap: flexProps.wrap,
            rowGap: flexProps.rowGap,
            columnGap: flexProps.columnGap
        ) {
            ForEach(children, id: \.self) { nodeId in
                if let childWrapper = self.wrapper.hostReceiver?.viewWrapperRegistry[nodeId] {
                    childWrapper.render()
                        .applyFlexChildProps(from: childWrapper.props)
                }
            }
        }
    }

    private struct FlexProps {
        var direction: YGFlexDirection = .column
        var justifyContent: YGJustify = .flexStart
        var alignItems: YGAlign = .stretch
        var wrap: YGWrap = .noWrap
        var rowGap: CGFloat = 0
        var columnGap: CGFloat = 0
    }

    private func extractFlexProps(from props: SolidNativeProps) -> FlexProps {
        var result = FlexProps()

        guard let style = props["style"], case .object(let styleProps) = style else {
            return result
        }

        // flexDirection
        if let fd = styleProps["flexDirection"], case .string(let dir) = fd {
            switch dir {
            case "row": result.direction = .row
            case "column": result.direction = .column
            case "row-reverse": result.direction = .rowReverse
            case "column-reverse": result.direction = .columnReverse
            default: break
            }
        }

        // justifyContent
        if let jc = styleProps["justifyContent"], case .string(let justify) = jc {
            switch justify {
            case "flex-start": result.justifyContent = .flexStart
            case "flex-end": result.justifyContent = .flexEnd
            case "center": result.justifyContent = .center
            case "space-between": result.justifyContent = .spaceBetween
            case "space-around": result.justifyContent = .spaceAround
            case "space-evenly": result.justifyContent = .spaceEvenly
            default: break
            }
        }

        // alignItems
        if let ai = styleProps["alignItems"], case .string(let align) = ai {
            switch align {
            case "flex-start": result.alignItems = .flexStart
            case "flex-end": result.alignItems = .flexEnd
            case "center": result.alignItems = .center
            case "stretch": result.alignItems = .stretch
            case "baseline": result.alignItems = .baseline
            default: break
            }
        }

        // flexWrap
        if let fw = styleProps["flexWrap"], case .string(let wrap) = fw {
            switch wrap {
            case "wrap": result.wrap = .wrap
            case "nowrap": result.wrap = .noWrap
            case "wrap-reverse": result.wrap = .wrapReverse
            default: break
            }
        }

        // gap
        if let gap = styleProps["gap"], case .number(let gapVal) = gap {
            result.rowGap = CGFloat(gapVal)
            result.columnGap = CGFloat(gapVal)
        }
        if let rowGap = styleProps["rowGap"], case .number(let gapVal) = rowGap {
            result.rowGap = CGFloat(gapVal)
        }
        if let colGap = styleProps["columnGap"], case .number(let gapVal) = colGap {
            result.columnGap = CGFloat(gapVal)
        }

        return result
    }
}

// MARK: - View Extension for Flex Child Props

extension View {
    @ViewBuilder
    func applyFlexChildProps(from props: SolidNativeProps) -> some View {
        var view = AnyView(self)

        guard let style = props["style"], case .object(let styleProps) = style else {
            view
            return
        }

        var result: AnyView = view

        // flex (shorthand for flexGrow)
        if let flex = styleProps["flex"], case .number(let val) = flex {
            result = AnyView(result.flexGrow(CGFloat(val)))
        }

        // flexGrow
        if let fg = styleProps["flexGrow"], case .number(let val) = fg {
            result = AnyView(result.flexGrow(CGFloat(val)))
        }

        // flexShrink
        if let fs = styleProps["flexShrink"], case .number(let val) = fs {
            result = AnyView(result.flexShrink(CGFloat(val)))
        }

        // width
        if let w = styleProps["width"] {
            result = AnyView(result.width(yogaDimension(from: w)))
        }

        // height
        if let h = styleProps["height"] {
            result = AnyView(result.height(yogaDimension(from: h)))
        }

        // minWidth
        if let mw = styleProps["minWidth"] {
            result = AnyView(result.minWidth(yogaDimension(from: mw)))
        }

        // maxWidth
        if let mxw = styleProps["maxWidth"] {
            result = AnyView(result.maxWidth(yogaDimension(from: mxw)))
        }

        // minHeight
        if let mh = styleProps["minHeight"] {
            result = AnyView(result.minHeight(yogaDimension(from: mh)))
        }

        // maxHeight
        if let mxh = styleProps["maxHeight"] {
            result = AnyView(result.maxHeight(yogaDimension(from: mxh)))
        }

        // margin
        if let m = styleProps["margin"] {
            let val = ygValue(from: m)
            result = AnyView(result
                .marginTop(val)
                .marginRight(val)
                .marginBottom(val)
                .marginLeft(val))
        }

        // marginTop, marginRight, marginBottom, marginLeft
        if let mt = styleProps["marginTop"] {
            result = AnyView(result.marginTop(ygValue(from: mt)))
        }
        if let mr = styleProps["marginRight"] {
            result = AnyView(result.marginRight(ygValue(from: mr)))
        }
        if let mb = styleProps["marginBottom"] {
            result = AnyView(result.marginBottom(ygValue(from: mb)))
        }
        if let ml = styleProps["marginLeft"] {
            result = AnyView(result.marginLeft(ygValue(from: ml)))
        }

        // padding
        if let p = styleProps["padding"] {
            let val = ygValue(from: p)
            result = AnyView(result
                .paddingTop(val)
                .paddingRight(val)
                .paddingBottom(val)
                .paddingLeft(val))
        }

        // paddingTop, paddingRight, paddingBottom, paddingLeft
        if let pt = styleProps["paddingTop"] {
            result = AnyView(result.paddingTop(ygValue(from: pt)))
        }
        if let pr = styleProps["paddingRight"] {
            result = AnyView(result.paddingRight(ygValue(from: pr)))
        }
        if let pb = styleProps["paddingBottom"] {
            result = AnyView(result.paddingBottom(ygValue(from: pb)))
        }
        if let pl = styleProps["paddingLeft"] {
            result = AnyView(result.paddingLeft(ygValue(from: pl)))
        }

        // alignSelf
        if let aself = styleProps["alignSelf"], case .string(let align) = aself {
            var ygAlign: YGAlign = .auto
            switch align {
            case "auto": ygAlign = .auto
            case "flex-start": ygAlign = .flexStart
            case "flex-end": ygAlign = .flexEnd
            case "center": ygAlign = .center
            case "stretch": ygAlign = .stretch
            case "baseline": ygAlign = .baseline
            default: break
            }
            result = AnyView(result.alignSelf(ygAlign))
        }

        result
    }
}

// MARK: - Helper Functions

private func yogaDimension(from value: JsValue) -> YogaDimension {
    switch value {
    case .number(let num):
        return .point(Float(num))
    case .string(let str):
        if str == "auto" {
            return .auto
        } else if str.hasSuffix("%"), let num = Float(str.dropLast()) {
            return .percent(num / 100.0)
        } else if let num = Float(str) {
            return .point(num)
        }
        return .auto
    default:
        return .auto
    }
}

private func ygValue(from value: JsValue) -> YGValue {
    switch value {
    case .number(let num):
        return YGValue(value: Float(num), unit: .point)
    case .string(let str):
        if str == "auto" {
            return YGValue(value: 0, unit: .auto)
        } else if str.hasSuffix("%"), let num = Float(str.dropLast()) {
            return YGValue(value: num, unit: .percent)
        } else if let num = Float(str) {
            return YGValue(value: num, unit: .point)
        }
        return YGValue(value: 0, unit: .auto)
    default:
        return YGValue(value: 0, unit: .auto)
    }
}
