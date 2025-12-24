
# DSA – Data Structures & Algorithms in Go

A clean, idiomatic, and job-oriented **Data Structures and Algorithms library implemented in Go**, designed to follow **Go package conventions**, **API discipline**, and **industry-standard repository structure**.

This repository is intended for:

* Backend engineers using Go
* Interview preparation with production-quality code
* Learning Go through real data structure implementations
* Reusable Go packages (not just DSA practice)

---

## Features

* Idiomatic Go package design
* Generic implementations (Go 1.20+)
* Clean separation of algorithms and data structures
* Fully testable and extensible
* No nested modules — single `go.mod`
* Designed like a real Go library, not a playground

---

## Installation

```bash
go get github.com/ak-repo/dsa
```

---

## Project Structure

```text
dsa/
├── algorithms/
│   ├── searching/
│   │   └── binarysearch/
│   └── sorting/
│       ├── bubble/
│       ├── insertion/
│       ├── merge/
│       ├── quick/
│       ├── heap/
│       └── selection/
│
├── structures/
│   ├── list/
│   │   ├── singly/
│   │   └── doubly/
│   ├── stack/
│   ├── queue/
│   ├── heap/
│   │   ├── min/
│   │   └── max/
│   ├── tree/
│   │   ├── binary/
│   │   └── bst/
│   ├── graph/
│   │   ├── adjacencylist/
│   │   └── adjacencymatrix/
│   └── hashtable/
│
├── problems/
│   └── linkedlist/
│
├── internal/
│   └── utils/
│
├── go.mod
└── README.md
```

---

## Usage Examples

### Stack

```go
import "github.com/ak-repo/dsa/structures/stack"

s := stack.New[int]()
s.Push(10)
s.Push(20)

value, ok := s.Pop()
```

---

### Queue

```go
import "github.com/ak-repo/dsa/structures/queue"

q := queue.New[string]()
q.Enqueue("A")
q.Enqueue("B")

v, _ := q.Dequeue()
```

---

### Binary Search

```go
import "github.com/ak-repo/dsa/algorithms/searching/binarysearch"

idx := binarysearch.Search([]int{1, 3, 5, 7}, 5)
```

---

### Quick Sort

```go
import "github.com/ak-repo/dsa/algorithms/sorting/quick"

arr := []int{5, 3, 1, 4}
quick.Sort(arr)
```

---

## Design Philosophy

* **One package = one responsibility**
* **Pure functions for algorithms**
* **Minimal exported surface**
* **Readable over clever**
* **Go standard library style**

This repository intentionally avoids:

* Nested `go.mod` files
* Over-engineering
* Framework-like abstractions

---

## Testing

Run all tests:

```bash
go test ./...
```

Run benchmarks (where available):

```bash
go test -bench=. ./...
```

---

## Versioning

This project follows **Semantic Versioning**:

```
vMAJOR.MINOR.PATCH
```

Breaking API changes will increment `MAJOR`.

---

## Roadmap

* [ ] Add benchmarks for sorting algorithms
* [ ] Add graph traversal algorithms (BFS, DFS)
* [ ] Add advanced tree structures (AVL, Red-Black)
* [ ] Add concurrency-safe variants
* [ ] Improve generic constraints

---

## Who Should Use This?

* Go backend engineers
* Candidates preparing for Go interviews
* Developers learning data structures the Go way
* Engineers building reusable libraries

---

## License

MIT License

---

## Author

**Ananda Krishnan**
Backend Engineer | Golang


