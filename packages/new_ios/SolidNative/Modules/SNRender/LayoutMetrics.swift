//
//  LayoutMetrics.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/13/24.
//

import Foundation


/// Custom Data type to set layout stuff for view wrapper.
/// Will take in view  props
struct LayoutMetrics {
    var x: CGFloat = 0
    var y: CGFloat = 0
    var width: CGFloat = 0
    var height: CGFloat = 0
    var contentLeftInset: CGFloat = 0
    var contentTopInset: CGFloat = 0
    var contentRightInset: CGFloat = 0
    var contentBottomInset: CGFloat = 0
    
    var contentFrame: CGRect {
      return CGRect(
        x: contentLeftInset,
        y: contentTopInset,
        width: width - contentLeftInset - contentRightInset,
        height: height - contentTopInset - contentBottomInset
      )
    }
}
