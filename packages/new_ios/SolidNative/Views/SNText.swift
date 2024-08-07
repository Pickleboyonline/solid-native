import Foundation
import SwiftUI
import Snmobile

struct SNText: SolidNativeView {
    static var name: String {
        "sn_text"
    }
    static var isTextElement: Bool {
        true
    }
    
    static func measureNode(nodeId: String) -> SNSnmobileSize {
        let textNode = SharedSolidNativeCore.viewWrapperRegistry[nodeId]!
        // TODO: Need to make function that:
        // Grabs View Wrapper from node ID
        // TODO: Make some cache any type on the wrapper for state
        
        var fontSize = UIFont.systemFontSize
        if let styles = textNode.props["style"] {
            if let size = styles.getForKey("fontSize"), size.isNumber() {
                fontSize = size.getNumber()
            }
        }
        // For text, it has to make a UIKit text element tfor this.
        /*
         Attributes that matter:
         - Font (.font)
         - Paragraph Style (.paragraphStyle)
         - Kern (.kern)
         */
        let attibutes = [NSAttributedString.Key.font: fontSize]
        
        let displayedText = {
            if let text = textNode.props["text"], text.isString() {
                return text.getString()
            }
            return ""
        }()
        
        let size = sizeOfString(displayedText, withAttributes: attibutes)
        
        return SNSnmobileSize(Float(size.width), height: Float(size.height))!
    }
    
    var props: SolidNativeProps
    var children: SolidNativeChildren

    func textStyle(from style: SNSnmobileJSValue) -> Font {
        var font = Font.system(size: 14)

        if let fontFamily = style.getForKey("fontFamily"), fontFamily.isString() {
            let family = fontFamily.getString()
            if let fontSize = style.getForKey("fontSize"), fontSize.isNumber() {
                font = Font.custom(family, size: CGFloat(fontSize.getNumber()))
            } else {
                font = Font.custom(family, size: 14)
            }
        } else if let fontSize = style.getForKey("fontSize"), fontSize.isNumber() {
            font = Font.system(size: CGFloat(fontSize.getNumber()))
        }

        if let fontStyle = style.getForKey("fontStyle"), fontStyle.isString() {
            if fontStyle.getString() == "italic" {
                font = font.italic()
            }
        }

        if let fontWeight = style.getForKey("fontWeight"), fontWeight.isString() {
            switch fontWeight.getString() {
            case "bold":
                font = font.weight(.bold)
            case "100":
                font = font.weight(.ultraLight)
            case "200":
                font = font.weight(.thin)
            case "300":
                font = font.weight(.light)
            case "400":
                font = font.weight(.regular)
            case "500":
                font = font.weight(.medium)
            case "600":
                font = font.weight(.semibold)
            case "700":
                font = font.weight(.bold)
            case "800":
                font = font.weight(.heavy)
            case "900":
                font = font.weight(.black)
            default:
                break
            }
        }

        return font
    }

    func color(from style: SNSnmobileJSValue) -> Color {
        if let colorValue = style.getForKey("color"), colorValue.isString() {
            return Color(hex: colorValue.getString())
        }
        return Color.black
    }

    func applyTextStyles(_ text: Text, style: SNSnmobileJSValue) -> Text {
        var styledText = text.font(textStyle(from: style))
        
        styledText = styledText.foregroundColor(color(from: style))
        
        if let letterSpacing = style.getForKey("letterSpacing"), letterSpacing.isNumber() {
            styledText = styledText.kerning(CGFloat(letterSpacing.getNumber()))
        }
        
        if let lineHeight = style.getForKey("lineHeight"), lineHeight.isNumber() {
            styledText = styledText.lineSpacing(CGFloat(lineHeight.getNumber())) as! Text
        }
        
        if let textAlign = style.getForKey("textAlign"), textAlign.isString() {
            switch textAlign.getString() {
            case "left":
                styledText = styledText.multilineTextAlignment(.leading) as! Text
            case "right":
                styledText = styledText.multilineTextAlignment(.trailing) as! Text
            case "center":
                styledText = styledText.multilineTextAlignment(.center) as! Text
            default:
                break
            }
        }
        


        return styledText
    }
    
    

    var body: some View {
        if let txt = props["text"],
           txt.isString() {
            var textView = Text(txt.getString())
            
            if let style = props["style"], style.isObject() {

                if let textTransform = style.getForKey("textTransform"), textTransform.isString() {
                    switch textTransform.getString() {
                    case "uppercase":
                        textView = Text(txt.getString().uppercased())
                    case "lowercase":
                        textView = Text(txt.getString().lowercased())
                    case "capitalize":
                        textView = Text(txt.getString().capitalized(with: nil))
                    default:
                        break
                    }
                }
                
                textView = applyTextStyles(textView, style: style)
            }
            
            return AnyView(textView)
        } else {
            return AnyView(EmptyView())
        }
    }
}


private func sizeOfString(_ string: String, withAttributes attributes: [NSAttributedString.Key: Any], constrainedToWidth width: CGFloat? = nil) -> CGSize {
    var size = (string as NSString).size(withAttributes: attributes)
    
    if let width = width {
        let constraintRect = CGSize(width: width, height: .greatestFiniteMagnitude)
        let boundingBox = string.boundingRect(
            with: constraintRect,
            options: [.usesLineFragmentOrigin, .usesFontLeading],
            attributes: attributes,
            context: nil
        )
        size = boundingBox.size
    }
    
    return size
}
