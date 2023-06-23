// Revamp prop Swift object to take JS values instead of Any values

// Add get<Type> functions for all types

// Cleanup APIs

// Live reloading

// Get more SwiftUI components

// Figure out styling, how its done.

/**
 * Some changes:
 * I want there to be a a separate directory for iOS and Android.
 * The reason being is to not encourage huge abstractions between the
 * two platforms and make it more obvious as to what you're dealing with.
 *
 * The idea is to have two entry points but a shared package between them
 * of different modules. This way one is encouraged to split up dev concerns.
 *
 * For the linker, I want to make it very light weight. Maybe some autogen
 * Swift/Kotlin file that can resolve the dependencies for you and put them
 * in an explicit file.
 *
 * An overall goal for this project is to provide a lightweight abstraction
 * so that TS devs can be productive but not so far abstracted to the point
 * where fixing issues, debugging, profiling, upgrading become a troublesome
 * like react native or expo.
 */
