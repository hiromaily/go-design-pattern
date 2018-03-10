# go-design-pattern
Go Design Patterns

## Creational Patterns
It groups common practices for creating objects.

### Singleton
It has a unique instance of a type in the entire program.

### Builder
It reuses an algorithm to create many implementations of an interface.
However I'd not use this patterns because of lengthy code.

### Factory method
It delegates the creation of different types of objects and defines an interface for creating an object, but let subclasses decide which class to instantiate. 
Factory Method lets a class defer instantiation to subclasses.
It'd useful to handle several categorized data which have same functions(actions) but different content.
