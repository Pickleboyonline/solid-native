//
//  ContentView.swift
//  SolidNative
//
//  Created by Imran Shitta-Bey on 6/20/24.
//

import SwiftUI

struct ContentView: View {
    @State private var isHidden = false
    
    var body: some View {
        VStack {
            Button("Toggle Hidden") {
                isHidden.toggle()
            }
            
            Text("Hello")
                .hidden()
                .onAppear {
                    print("View appeared") // Only prints once when view is first created
                }
                .onDisappear {
                    print("View disappeared") // Won't print when using .hidden()
                }
        }
    }
}

#Preview {
    ContentView()
}
