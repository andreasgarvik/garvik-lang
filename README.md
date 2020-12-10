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
m = [[1,2,3],[4,5,6]]
m[1][1]
// 5

// "Passing lists
// x = [1,2,3]
// f = l -> l[0] + l[1] + l[2]
// f(x)
// 6

// l = [2, x -> x * 2,[1,2,3]]
// times = x ->
    if x[0] == 0
      then x[2]
    else
      times([x[0] - 1, x[1], [x[1](x[2][0]),x[1](x[2][1]),x[1](x[2][2])]])
// times(l)
// [4,8,12]

// l = [0, x -> x * 2,[1,2,3]]
// map = x ->
    if x[0] == len x[2]
      then x[2]
    else
      map([x[0] + 1, x[1], x[2][x[0]] = (x[1](x[2][x[0]]))])
// map(l)
// [2,4,6]

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

f = x -> x[0].f(x[1])
s = { a = 2 f = x -> x + a}
f([s, 5])
// 7

// "Fibonacci"
fib = x -> if x == 0 then 0 else if x == 1 then 1 else fib(x-1) + fib(x-2)
fib(10)
// 55
```

## How to run

```
go run main.go *file*
```
