# Program examples

### Increment
```lisp
(def ++ (fn (n) (+ n 1)))
```

### Fibonacci
```lisp
(def fib (fn (n) (if (< n 2) n (+ (fib (- n 1)) (fib (- n 2))))))
```
