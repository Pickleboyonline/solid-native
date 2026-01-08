//
//  ContentView.swift
//  DevApp
//
//  Created by Imran Shitta-Bey on 1/7/26.
//

import SwiftUI
import SNCore

struct ContentView: View {
    @State private var firstNumber: String = ""
    @State private var secondNumber: String = ""
    @FocusState private var isTextFieldFocused: Bool
    
    private var firstUInt32: UInt32? {
        UInt32(firstNumber)
    }
    
    private var secondUInt32: UInt32? {
        UInt32(secondNumber)
    }
    
    private var result: UInt32? {
        guard let first = firstUInt32, let second = secondUInt32 else {
            return nil
        }
        return add(a: first, b: second)
    }
    
    var body: some View {
        NavigationView {
            ScrollView {
                VStack(spacing: 20) {
                    Image(systemName: "plus.calculator")
                        .imageScale(.large)
                        .foregroundStyle(.tint)
                        .padding(.top, 20)
                    
                    Text("UInt32 Adder")
                        .font(.title2)
                        .fontWeight(.semibold)
                    
                    VStack(spacing: 16) {
                        VStack(alignment: .leading, spacing: 8) {
                            Text("First Number:")
                                .font(.headline)
                            TextField("Enter first number", text: $firstNumber)
                                .textFieldStyle(.roundedBorder)
                                .keyboardType(.numberPad)
                                .focused($isTextFieldFocused)
                                .submitLabel(.done)
                                .onSubmit {
                                    isTextFieldFocused = false
                                }
                            
                            if !firstNumber.isEmpty && firstUInt32 == nil {
                                Text("Please enter a valid UInt32 (0 to 4,294,967,295)")
                                    .font(.caption)
                                    .foregroundColor(.red)
                            }
                        }
                        
                        VStack(alignment: .leading, spacing: 8) {
                            Text("Second Number:")
                                .font(.headline)
                            TextField("Enter second number", text: $secondNumber)
                                .textFieldStyle(.roundedBorder)
                                .keyboardType(.numberPad)
                                .focused($isTextFieldFocused)
                                .submitLabel(.done)
                                .onSubmit {
                                    isTextFieldFocused = false
                                }
                            
                            if !secondNumber.isEmpty && secondUInt32 == nil {
                                Text("Please enter a valid UInt32 (0 to 4,294,967,295)")
                                    .font(.caption)
                                    .foregroundColor(.red)
                            }
                        }
                    }
                    
                    // Done button for keyboard dismissal
                    if isTextFieldFocused {
                        Button("Done") {
                            isTextFieldFocused = false
                        }
                        .buttonStyle(.borderedProminent)
                        .padding(.top, 10)
                    }
                    
                    Divider()
                        .padding(.vertical, 10)
                    
                    VStack(spacing: 8) {
                        Text("Result:")
                            .font(.headline)
                        
                        if let result = result {
                            Text("\(result)")
                                .font(.title)
                                .fontWeight(.bold)
                                .foregroundColor(.green)
                        } else if !firstNumber.isEmpty && !secondNumber.isEmpty {
                            Text("Invalid input")
                                .font(.title)
                                .fontWeight(.bold)
                                .foregroundColor(.red)
                        } else {
                            Text("Enter both numbers")
                                .font(.title)
                                .fontWeight(.medium)
                                .foregroundColor(.secondary)
                        }
                    }
                    
                    Spacer(minLength: 50)
                }
                .padding()
                .frame(maxWidth: .infinity, maxHeight: .infinity)
            }
            .navigationTitle("")
            .navigationBarHidden(true)
            .onTapGesture {
                // Dismiss keyboard when tapping outside text fields
                isTextFieldFocused = false
            }
        }
        .navigationViewStyle(StackNavigationViewStyle()) // Ensures full screen on all devices
    }
}

#Preview {
    ContentView()
}
