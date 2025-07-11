# (Θ) Programming Manual

This is the programming manual for the **Theta** list processor.

### Expressions

Symbolic expressions are the essence of LISP.
An expression can be one of several types, and all types have specific symbolic notation:

- Nil `nil`
- Boolean `true` | `false`
- Integer `7`, `-10`
- Floating `3.14`, `4.0`
- String `"Hello, world!"`
- Symbol `+`, `my-symbol`
- Keyword `:ok`, `:name`
- Procedure or function `:bin`
- List `()`, `(1 2 3)`, `((1 3) (2 4))`

Symbols and lists are special, as they provide the means for combination and abstraction.
Symbols can abstract any other type;
so can there be a symbol `pi` which abstracts the floating number `3.14`, or a symbol `+` which abstracts the addition function.

Lists are also used to evaluate and applicate functions abd procedures.
To add some numbers we can evaluate the list `(+ 1 2 3)`, which returns `6`.
Functions and procedures are distinct types but can be used interchangeably;
Procedures are builtin, while functions are written by the user.

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
Construct a function.
- Syntax: `(fn (SYM...) EXP)`, `(lambda (SYM...) EXP)`
- Example: `(fn (i) (+ i 1))` => `:bin`
- Example: `(apply (fn (i) (+ i 1)) 4)` => `5`

**Let expression**
Evaluate with local bindings.
- Syntax: `(let ((SYM INIT-EXP)...) EXP...)`
- Example: `(let ((x 10) (y (* 2 2))) (list :result (* x y)))` => `(:result 40)`

### Site/Standard library

**Equal**
Check equality.
- Syntax: `(= X Y)`
- Example: `(= 4 4)` => `true`
- Example: `(= "Hello" "Goodbye")` => `false`

**Inequality**
Check equality.
- Syntax: `(!= X Y)`
- Example: `(!= 4 4)` => `false`
- Example: `(!= "Hello" "Goodbye")` => `true`

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

**List**
Construct a list.
- Syntax: `(list EXP...)`
- Example: `(list :one 1 :ten 10)` => `(:one 1 :ten 10)`

**Car**
Get the first part of the list.
- Syntax: `(car EXP)`
- Example: `(car (quote (foo bar baz bat)))` => `foo`

**Cdr**
Get the rest part of the list.
- Syntax: `(cdr EXP)`
- Example: `(cdr (quote (foo bar baz bat)))` => `(bar baz bat)`

**Apply**
Apply a function to its arguments.
- Syntax: `(apply FUNCTION EXP...)`
- Example: `(apply + 2 3 4)` => `9`

**Map**
Map list with function.
- Syntax: `(map FUNCTION LIST)`
- Example: `(def ++ (fn (i) (+ i 1)))` => `++`
- Example: `(map ++ (list 2 4 6 8 10))` => `(3 5 7 9 11)`

**Filter**
Filter list with function.
- Syntax: `(filter FUNCTION LIST)`
- Example: `(filter key? (list 1 :two 3 :four 5))` => `(:two :four)`

**Concat**
Concatenate strings.
- Syntax: `(concat STRING...)`
- Example: `(concat "eeni" "meeni" "mini" "mo")` => `"eeni meeni mini mo"`

**Printf**
Print formatted strings.
- Syntax: `(printf FORMAT-STRING EXP...)`
- Example: `(printf "%s was released in %d" "Linux 1.0" 1994)` => `"Linux 1.0 was released in 1994"`

**Unfunction**
Deconstruct a function.
- Syntax: `(!fn FUNCTION)`
- Example: `(!fn ++)` => `(fn (i) (+ i 1))`
