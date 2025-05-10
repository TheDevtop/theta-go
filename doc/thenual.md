# (Θ) Programming Manual

This is the programming manual for the **Theta** list processor.

### Expressions

Symbolic expressions are the essence of LISP.
An expression can be one of several types, and all types have specific symbolic notation:

- Nil `:nil`
- Boolean `:true` | `:false`
- Integer `7`, `-10`
- Floating `3.14`, `4.0`
- String `"Hello, world!"`
- Symbol `+`, `my-symbol`
- Keyword `:ok`, `:name`
- Function `:fn`
- List `()`, `(1 2 3)`, `((1 3) (2 4))`

Symbols and lists are special, as they provide the means for combination and abstraction.
Symbols can abstract any other type;
so can there be a symbol `pi` which abstracts the floating number `3.14`, or a symbol `+` which abstracts the addition function.

Lists are also used to evaluate and applicate functions.
To add some numbers we can evaluate the list `(+ 1 2 3)`, which returns `6`.

### Special forms

**Quote**
Skip evaluation on expression.
- Syntax: `(quote EXP)`
- Example: `(quote foobar)` => `foobar`
- Example: `(quote (foo bar baz))` => `(foo bar baz)`

**Define**
Define new symbolic abstraction.
- Syntax: `(def SYMBOL EXP)`
- Example: `(def deca 10)` => `deca`
- Example: `deca` => `10`

**If Expression**
Perform conditional evaluation.
- Syntax: `(if COND THEN-EXP ELSE-EXP)`
- Example: `(if (= 10 10) "Equal" "Not equal")` => `"Equal"`
- Example: `(if (= 10 9) "Equal" "Not equal")` => `"Not equal"`

**While expression**
While condition evaluate, returns last expression.
- Syntax: `(while COND EXP...)`
- Example: `(def i 0)` => `:ok`
- Example: `(while (< i 10) (def i (+ i 1)))` => `:ok`
- Example: `i` => `10`

**Sequence**
Evaluate multiple expressions, return last expression.
- Syntax: `(seq EXP...)`
- Example: `(seq (def deca 10) (+ deca 2))` => `12`

**Lambda expression**
Construct an anonymous function.
- Syntax: `(fn (SYM...) EXP)`
- Example: `(fn (i) (+ i 1))` => `:fn`
- Example: `(apply (fn (i) (+ i 1)) 4)` => `5`

**Let expression**
Evaluate with local bindings.
- Syntax: `(let ((SYM INIT-EXP)...) EXP...)`
- Example: `(let ((x 10) (y (* 2 2))) (list :result (* x y)))` => `(:result 40)`

### Site/Standard library

**Equal**
Check equality.
- Syntax: `(= X Y)`
- Example: `(= 4 4)` => `:true`
- Example: `(= "Hello" "Goodbye")` => `:false`

**Inequality**
Check equality.
- Syntax: `(!= X Y)`
- Example: `(!= 4 4)` => `:false`
- Example: `(!= "Hello" "Goodbye")` => `:true`

**Add**
Add numbers.
- Syntax: `(+ EXP...)`
- Example: `(+ 8 2)` => `10`

**Subtract**
Subtract numbers.
- Syntax: `(- EXP...)`
- Example: `(- 10 5 3)` => `2`

**Multiply**
Multiply numbers.
- Syntax: `(* EXP...)`
- Example: `(* 3.14 2.2)` => `6.9080`

**Divide**
Divide numbers.
- Syntax: `(/ EXP...)`
- Example: `(/ 7 5)` => `1`
