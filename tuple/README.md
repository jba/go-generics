This package introduces generic tuples. Here I justify why they are useful.

Tuples are somewhat useful their own. For instance, they would come in handy
when parallelizing a function with more than one return value:

```
func f(int) (int, error) ...

var ints []int = ...
c := make(chan tuple.T2[int, error], len(ints))
for _, i := range ints {
    i := i
    go func() { c <- tuple.New2(f(i)) }
}
```

However, I don't think they carry their weight when viewed in isolation.

But they will turn out to be extremely useful when defining other generic
packages. They can help reduce or eliminate "arity copying," defining multiple
types or functions to handle different numbers of arguments.

The generic metrics package included in the prototype translator is a good
example. It has types Metric1, Metric2 and so on, to support metrics with
different numbers of dimensions. Each MetricN type has copies of all the Metric
methods with only slight variations. You might argue that as bad as this is, it
is at least confined to the package itself; users don't feel the pain. But what
if a user wanted to write a function that took a Metric as an argument, or
returned one? They would need a separate function for each MetricN type. Arity
copying is infectious.

Now consider the metrics package in this repo, built with tuples in mind. There
is only one type, `Metric[T]`, and one set of methods. `T` can be any comparable
type, or a tuple of two or more comparable types--which, thanks to Go's rules on
struct equality, is itself comparable.

The metrics package is practically a best case for tuples. In other cases,
tuples can't eliminate arity copying, but they can mitigate it. The file
examples_test.go contains the outline of a property-testing framework, greatly
simplified from github.com/leanovate/gopter. The basic idea is to generate a lot
of random arguments for a boolean function (the predicate) and see if it ever
returns false on them. The gopter package handles functions of any arity via
reflection. To build a type-safe version using generics, arity copying is needed
somewhere, but we can at least define a single Check function that takes a
one-argument predicate and a single generator of random values:

```
func Check[T any](max int, pred func(T) bool, gen Gen[T])
```

The file presents two designs that start from there, both built on the
fact that the `T` parameter of `Check` can be a tuple.

In the first design, users must adapt their multi-argument predicate
to take a tuple, which is easily done via the Spread method defined on
each tuple type:

```
func pred(int, string) bool ...

func wrapper(t tuple.T2[int, string]) bool { return pred(t.Spread()) }
```

Combining N individual generators into one that generates an N-tuple is harder,
so the package provides a series of CombineGensN functions for that purpose. By
means of a small local helper and the tuples.JoinN functions, the Nth
CombineGens function can be defined in terms of the (N-1)st in one line:

```
func CombineGens3[type A, B, C](ga Gen[A], gb Gen[B], gc Gen[C]) Gen(tuple.T3[A, B, C]) {
    return joinGen(CombineGens2(ga, gb), gc, tuple.Join3(A, B, C))
}
```

This is possible because instead of the obvious representation of an N-tuple as
N separate fields, the tuples package defines it as an (N-1)-tuple with a single
new field. You can also see this pseudo-recursion at work in the `Nth` methods
defined on each tuple type.

In the second design of the property tester, a number of CheckN functions are
defined as small wrappers around the one-argument `Check`:

```
func Check2[type A, B](max int, pred func(A, B) bool, genA Gen[A], genB Gen[B]) (A, B, bool) {
    p := func(t tuple.T2[A, B]) bool { return pred(t.Spread()) }
    g := func() tuple.T2[A, B] { return tuples.New2(genA(), genB()) }
    r, ok := Check(max, p, g)
    return r.V0(), r.V1(), ok
}
```

This is facilitated by the Spread method and the NewN constructor functions.

In general, whenever there is a need to support functions of varying arity in
arguments or return values, a single function signature `func(T) R` will
suffice, with tuples providing the glue.
