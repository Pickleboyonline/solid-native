//
//  ContentView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 1/12/24.
//

import SwiftUI
import YogaSwiftUI
import Yoga

struct ContentView: View {
    var body: some View {
        Flex(direction: .column, alignItems: .stretch)  {
            Flex(direction: .column, alignItems: .flexStart) {
                Group {
                    Text("Row 1 Item 1").background(Color.purple)
                }
                
                
                Flex(direction: .row, alignItems: .stretch) {
                    Text("Row 2 Item 1")
                       //  .alignItems(.)
                    Text("Row 2 Item 2")
                }.padding(.top, 100)
                
                Text("Heslslo")
            }.background(Color.green).alignItems(.flexStart)
        }.background(Color.gray)
    }
    
}

#Preview {
    ContentView()
}
