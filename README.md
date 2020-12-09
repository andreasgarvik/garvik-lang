# Programming language creation project

My own little programming language :sparkles:

## Features

```
// "Functions"
f = x -> x(2)
f(x -> x * 2)
// 4

// "Lists"
l = [1,2,3]
l[1]
// 2

// "Passing lists
// x = [1,2,3]
// f = l -> l[0] + l[1] + l[2]
// f(x)
// 6

// "Structs"
s = { a = 5 }
s.a
// 5

// "Passing structs
f = x -> x.b - x.a
s = {
  a = 2
  b = 3
}
f(s)
// 1

// "Fibonacci"
fib = x -> if x == 0 then 0 else if x == 1 then 1 else fib(x-1) + fib(x-2)
fib(10)
// 55
```

## How to run

```
go run main.go *file*
```
