# Project Rules

## Go Generic Struct Assignments
- Do NOT attempt to assign or convert generic struct types with different type parameters (e.g., assigning `DatabaseOptions[any]` to `DatabaseOptions[T]`).
- When extracting generic values from an `any`-typed container, use explicit type assertions on the underlying value field (`if val, ok := container.Value.(T); ok`).
