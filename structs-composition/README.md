# Go Structs & Composition

A complete, hands-on learning repository covering **Structs and Composition in Go**, from beginner fundamentals to production-ready backend design.

Our learning approach:

**Learn → Understand → Implement → Test → Review → Apply**

## Learning Roadmap

| Phase | Modules | Focus |
|---|---|---|
| Phase 1 | 01–05 | Struct Fundamentals |
| Phase 2 | 06–10 | Struct Composition |
| Phase 3 | 11–14 | Design Patterns |
| Phase 4 | 15–16 | Real-world Application |

## Repository Structure

```text
structs-composition/
├── README.md
├── go.mod
├── 01-struct-basics/
├── 02-struct-fields-and-tags/
├── 03-struct-pointers/
├── 04-struct-methods/
├── 05-constructors/
├── 06-nested-structs/
├── 07-embedding/
├── 08-promoted-fields-methods/
├── 09-composition-vs-inheritance/
├── 10-method-promotion-override/
├── 11-mixins/
├── 12-interfaces-and-composition/
├── 13-encapsulation/
├── 14-dependency-injection/
├── 15-domain-modeling/
├── 16-real-world-application/
├── assignments/
└── docs/
```

## Phase 1 — Struct Fundamentals

1. **Struct Basics:** Declaration, initialization, zero values.
2. **Struct Fields and Tags:** Exported fields, unexported fields, JSON tags, marshaling and unmarshaling.
3. **Struct Pointers:** Value semantics, pointer semantics and copying.
4. **Struct Methods:** Value receivers, pointer receivers and method behaviour.
5. **Constructors:** Constructor functions, validation and private struct initialization.

## Phase 2 — Struct Composition

6. **Nested Structs:** Structs inside structs and complex models.
7. **Embedding:** Anonymous fields, named fields and embedding.
8. **Promoted Fields and Methods:** Field and method promotion.
9. **Composition vs Inheritance:** Composition over traditional inheritance.
10. **Method Promotion and Override Trap:** Method shadowing, promotion and explicit embedded method calls.

## Phase 3 — Design Patterns

11. **Mixins:** Reusable behaviour through composition.
12. **Interfaces and Composition:** Implicit implementation and accept interfaces, return structs.
13. **Encapsulation:** Exported/unexported fields and package boundaries.
14. **Dependency Injection:** Interface-based dependencies and testability.

## Phase 4 — Real-world Application

15. **Domain Modelling:** User, Account, Order, Payment and BaseEntity.
16. **Real-world Application:** Build a backend-style application using composition, interfaces and constructors.

## Practical Assignments

| # | Assignment | Objective |
|---|---|---|
| 01 | The JSON Modeler | Convert nested JSON into Go structs |
| 02 | The Promoted Field | Embed BaseEntity into User |
| 03 | The Constructor Pattern | Build a private Server with NewServer |
| 04 | The Override Trap | Understand method promotion and shadowing |
| 05 | The Mixin | Combine Drivable and Flyable |
| 06 | The Pointer Trap | Understand copying and mutation |
| 07 | The Nested Model | Model Order, Customer and Address |
| 08 | The Behaviour Model | Attach business methods to structs |
| 09 | The Interface Bridge | Implement interfaces implicitly |
| 10 | The Repository Pattern | Inject a repository into a service |
| 11 | The Payment Model | Compose payment and transaction metadata |
| 12 | The Package Boundary | Practise encapsulation |
| 13 | The Configuration Builder | Build and validate configuration |
| 14 | The Mini Backend | Combine all concepts |

## Learning Progress

### Phase 1 — Fundamentals
- [ ] Struct declaration and initialization
- [ ] Fields and struct tags
- [ ] Value vs pointer semantics
- [ ] Methods and receivers
- [ ] Constructor pattern

### Phase 2 — Composition
- [ ] Nested structs
- [ ] Embedding fields
- [ ] Promoted fields and methods
- [ ] Composition vs inheritance
- [ ] Method promotion and override trap

### Phase 3 — Design Patterns
- [ ] Mixins
- [ ] Accept interfaces, return structs
- [ ] Encapsulation
- [ ] Dependency injection

### Phase 4 — Real-world Application
- [ ] Domain modelling
- [ ] Production-style composition

## How to Run

Check your Go installation:

```bash
go version
```

Run an individual module:

```bash
go run .
```

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Learning Philosophy

Every module should include:

1. Concept explanation.
2. Small code examples.
3. Practical exercises.
4. Expected output and test cases.
5. Common mistakes and edge cases.
6. Interview questions.
7. Completion checklist.

**Don't just copy the code. Understand why it works.**

## Final Goal

By completing this repository, you should be able to confidently design Go applications using structs and composition, understand embedding and method promotion, and build maintainable backend systems using Go's native design principles.

Happy Learning!
