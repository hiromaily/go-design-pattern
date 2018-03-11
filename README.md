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
It delegates the creation of different types of objects and defines an interface for creating an object,
 but let subclasses decide which class to instantiate. 
Factory Method lets a class defer instantiation to subclasses.
It'd useful to handle several categorized data which have same functions(actions) but different content.

### Abstract Factory
It's a factory of factories. It provides a way to encapsulate a group of individual factories 
that have a common theme without specifying their concrete classes.

### Prototype
It is used when the type of objects to create is determined by a prototypical instance, 
which is cloned to produce new objects.  
The main objective for the Prototype design pattern is to avoid repetitive object creation.


## Structural Patterns
It eases the design by identifying a simple way to realize relationships between entities.

### Composite
It's partitioning design pattern.

