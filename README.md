# My own little programming language :sparkles:

Evaluated using ANTLR and Golang :rocket:

## Features

```
// "Functions"
a = 5
f = b -> b * a
f(6)
// 30

f = x -> x(2)
f(x -> x * 2)
// 4

// "Let Expression"
let x = 2 + 3 in x * x
// 25

// "Lists"
l = [1,2,3]
l[1]
// 2

m = [[1,2,3],[4,5,6]]
m[1][1]
// 5

// "Passing lists"
x = [1,2,3]
f = l -> l[0] + l[1] + l[2]
f(x)
// 6

l = [4, x -> x + 2,[1,2,3]]
times = x ->
  let c = x[0] in
  let f = x[1] in
  let l = x[2] in
  if c == 0
    then l
  else
    times([c - 1, f, [f(l[0]),f(l[1]),f(l[2])]])
times(l)
// [9,10,11]

l = [0, x -> x * 2,[1,2,3]]
map = x ->
  let c = x[0] in
  let f = x[1] in
  let l = x[2] in
  if c == len l
    then l
  else
    map([c + 1, f, l[c] = (f(l[c]))])
map(l)
// [2,4,6]

// "Structs"
s = { A: 5 }
s.A = 6
s.A
// 6

// "Encapsulation"
obj = {
  i: 2
  GetI: _ -> i
  SetI: x -> (i = x)
}

obj.i
// i is not a field of obj
obj.GetI(_)
// 2
i
// i is not defined
i = 7
obj.SetI(5)
obj.GetI(_)
// 5
i
// 7

g = {
  count: 0
  Count: _ -> count
  UpdateCount: _ -> (count = count + 1)
}

g.UpdateCount(_)
g.Count(_)
// 1
g.UpdateCount(_)
g.Count(_)
// 2
g.UpdateCount(_)
g.Count(_)
// 3

// "Passing structs"
f = x -> x.SetP(x.L[1] + x.F(x.B - x.A))
s = {
  A: 10
  B: 20
  F: x -> x / 2
  L: [1,2,3]
  p: _
  GetP: _ -> p
  SetP: x -> (p = x)
}
f(s)
s.GetP(_)
// 7

f = x -> x[0].F(x[1])
s = {
  a: 5
  F: x -> x + a
}
f([s, 5])
// 10

// "Lexical scoping"
a = 1
fun = x -> x + a
a = 2
fun(3)
// 4

// "Fibonacci"
fib = x -> if x < 2 then 1 else fib(x-1) + fib(x-2)
fib(10)
// 89

// "Error messages"
a = b
// b is not defined
c = 5
c(5)
// c is not a function
c[0]
// c is not a list
l = [1,2,3]
l[3]
// list index out of bounds
l + 5
// l is not a number
```

## How to run

```
go get github.com/andreasgarvik/garvik-lang
garvik-lang *file*
```
