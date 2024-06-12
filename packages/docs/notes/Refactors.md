Need to improve the way types are marshaled between JS and Swift. Main types are usually some form of type conversion.

An article about how optionals work in Swift: [https://www.hackingwithswift.com/articles/136/the-complete-guide-to-optionals-in-swift]()

Side Note: Swift options are weird because the type information is not clearly shown and you have to use unwrap syntax to get the compiler to check it properly. 

Values are either numbers or string. They can also be a dictionary as well.

Also need to understand when a props is exictly undefined.

Could also use closures with generics. Basically when you have a view, you can build a system for the props to update.

So if a prop updates, theres a hook for it.

something like
```ts
addProp("name", (propValue: Dictionary) => {

})
```

Now, Swift can auto marshal types, the question is what happens with undefined and null behavior.

Also, we need refs. (need to figure out how solidjs handles this).

Expo Does a similar thing with their module system. I could change up my system to more mimic this now that i am more framework. 

