Creational Patterns are design patterns that deal with how objects are created.

Instead of creating objects directly and carelessly, these patterns help you create them in a controlled, 
flexible, and clean way.

Creational Patterns provide better ways to create objects while hiding the creation logic and reducing tight coupling.

**In Go terms:**
- You control how structs are initialized
- You avoid messy new / struct literals everywhere
- You keep code easy to change and test

**Problem WITHOUT Creational Patterns**
```
type User struct {
	Name  string
	Email string
	Age   int
}
```

Now you create users everywhere like this:
```
u1 := User{"Mayank", "mayank@gmail.com", 25}
u2 := User{"Amit", "", 0}
u3 := User{"", "test@gmail.com", 30}
```