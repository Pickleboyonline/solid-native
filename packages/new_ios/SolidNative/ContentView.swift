//
//  ContentView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/20/24.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        ZStack(alignment: .topLeading) {
            
            AnyView(Text("Hello World")).frame(
                width: 10,
                height: 40,
                alignment: .topLeading
            )
            
        }.edgesIgnoringSafeArea(.all).frame(
            width: 300,
            height: 500,
            alignment: .topLeading
        ).background(Color.red)
    }
}

#Preview {
    ContentView()
}
