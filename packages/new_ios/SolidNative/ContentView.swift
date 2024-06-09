//
//  ContentView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/12/24.
//

import SwiftUI
import YogaSwiftUI

struct ContentView: View {
    var body: some View {
        Flex(direction: .column)  {
            Flex(direction: .column) {
                Flex(direction: .row) {
                    Text("Row 1 Item 1")
                    Text("Row 1 Item 2")
                }
                Flex(direction: .row) {
                    Text("Row 1 Item 1")
                    Text("Row 1 Item 2")
                }
                
                Text("Heslslo")
            }.background(Color.green)
        }.background(Color.gray)
    }
    
}

#Preview {
    ContentView()
}
