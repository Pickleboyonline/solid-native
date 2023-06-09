package com.example.solid_native_core

/**
 * This module will be used by a JSExportProtocal wrapper that will take a
 * JSValue;
 *
 *
 */
class Module {
    /**
     * A dictionary of AnyFunctions based on their name.
     */
    // val functions;

    /**
     * Dictionary of any properties.
     */
    // val properties;
    /**
     * Dictionary of any properties
     */
    // val constants;
}

/**
 * Takes in module and exposes a
 * lookup function for JS land.
 *
 * JS tells it what its looking for amd this method returns
 * a method that unwraps the JSValue args based on its specifications,
 * converts that into a AnyArgument (which is made of AnyTypes, then passes that anytype to the function
 * which returns an AnyReturn, then parses the selected value.
 *
 * The point here is that we need the same type sig on everything to look up a function,
 * because I'm 99% sure there isnt a way to have nested, dynamic properties that contain function blocks.
 * If so, thats cool.
 *
 *
 */
final class JSObjectModule {

}

