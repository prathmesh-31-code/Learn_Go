# Memory Structure in Go (Golang)

In Go, memory management is designed to be simple, efficient, and safe. The language automatically handles allocation and deallocation using a combination of stack allocation, heap allocation, and garbage collection.

---

##  1. Overview of Go Memory Model

Go divides memory into several key regions:

* **Stack** → Fast, function-local variables
* **Heap** → Dynamically allocated memory
* **Global/Data Segment** → Global and package-level variables
* **Code Segment** → Compiled instructions

---

##  2. Stack Memory (Fast & Automatic)

### What it is:

* Stores **local variables** inside functions
* Managed automatically (push/pop)

### Characteristics:

* Very fast allocation/deallocation
* Each goroutine has its own stack (starts small, grows dynamically)
* No garbage collection needed

### Example:

```go
func add(a int, b int) int {
    sum := a + b // stored on stack
    return sum
}
```

 `a`, `b`, and `sum` are typically stored on the **stack**

---

##  3. Heap Memory (Dynamic Allocation)

### What it is:

* Stores data that **outlives function scope**
* Managed by Go’s **garbage collector (GC)**

### When used:

* When variables are referenced outside their scope
* Large objects
* When compiler decides via **escape analysis**

### Example:

```go
func getPointer() *int {
    x := 10
    return &x // moved to heap
}
```

 `x` escapes the function → stored on the **heap**

---

##  4. Escape Analysis (Key Concept)

What does "escape" mean?

In Go, a variable escapes when it cannot safely stay inside the function where it was created.

In simple terms:

* If a variable is used outside its function, it escapes
* If Go cannot guarantee it will be used only locally, it moves it to the heap

Why does this matter?
Stack memory is fast but temporary
Heap memory is slower but survives longer

So if something needs to live longer than the function → it must go to the heap

Go compiler decides:
 “Stack or Heap?”

### Rule:

* If variable **does NOT escape** → stack
* If variable **escapes** → heap

### Example:

```go
func foo() *int {
    x := 42
    return &x // escapes → heap
}
```

### Check escape analysis:

```bash
go build -gcflags="-m"
```

---

##  5. Global / Data Segment

### Stores:

* Global variables
* Package-level variables
* Constants (sometimes optimized)

### Example:

```go
var globalVar = 100
```

 Stored in **data segment**, not stack/heap

---

##  6. Garbage Collector (GC)

Go uses a **concurrent, tri-color mark-and-sweep GC**
The Garbage Collector (GC) in Go is responsible for automatically managing heap memory—so you don’t have to manually free memory like in C/C++.

### Responsibilities:

* Frees unused heap memory
* Prevents memory leaks

### Key Features:

* Runs concurrently with your program
* Low pause times
* Automatically triggered

---

##  7. Goroutine Stack Model

Each goroutine has:

* Independent stack
* Starts small (~2 KB)
* Grows/shrinks dynamically

This is different from traditional threads (which use fixed-size stacks)

---

##  8. Memory Flow Example

```go
func main() {
    x := 10          // stack
    y := &x          // pointer to stack
    z := new(int)    // heap
    *z = 20
}
```

| Variable | Location            |
| -------- | ------------------- |
| `x`      | Stack               |
| `y`      | Stack (points to x) |
| `z`      | Stack (pointer)     |
| `*z`     | Heap                |

---

##  9. Stack vs Heap (Comparison)

| Feature    | Stack          | Heap              |
| ---------- | -------------- | ----------------- |
| Speed      | Fast           | Slower            |
| Management | Automatic      | Garbage collected |
| Size       | Limited        | Large             |
| Lifetime   | Function scope | Flexible          |

---

##  10. Key Takeaways

* Go **prefers stack allocation** whenever possible
* **Escape analysis** decides memory placement
* Heap is managed by the **garbage collector**
* Goroutines use **lightweight dynamic stacks**
* Developers usually **don’t manually manage memory**
