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

Problems here - 
- Hard to remember field order
- Invalid users possible (empty name)
- Creation logic repeated everywhere
- If you add a new field → all code breaks

**WITH Creational Pattern (Factory – Simple)**
We centralize how User is created.

Factory function - 
```
func NewUser(name, email string, age int) User {
	if name == "" {
		panic("name cannot be empty")
	}
	return User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
```

Usage
```
u1 := NewUser("Mayank", "mayank@gmail.com", 25)
u2 := NewUser("Amit", "", 0)
```

Benefits - 
- Field order hidden
- Validation in one place
- Safer and cleaner
- Easy to change later
