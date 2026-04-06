The Singleton Pattern is a creational design pattern that ensures only ONE instance of a struct exists in the entire application and provides a global way to access it.

Singleton ensures a struct is created only once and reused everywhere.

**Why Do We Need Singleton?**
Imagine:
- Multiple DB connections created accidentally ❌
- Multiple config objects with different values ❌
- Multiple loggers writing inconsistently ❌