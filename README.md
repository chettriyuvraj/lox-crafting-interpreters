# lox-crafting-interpreters
Porting the tree-walk interpreter for the Lox programming language to Go. 

I am in the process of porting the Lox programming language, as described in [Part II of Crafting Interpreters](https://craftinginterpreters.com/contents.html) to Go.

- Ported the Scanner.
- Started porting the tokenizer, _(5) != (6)_ spits out an ugly-printed AST as _BinaryExpr { Left GroupingExpr { Expr LiteralExpr { Value 5 } };; Operator Type != Lexeme != Literal <nil> Line 0 ;; Right GroupingExpr { Expr LiteralExpr { Value 6 } } }_

# Usage

- _go run main.go [filename]_ OR _go run main.go_ which starts a prompt for you to input source code.

# Output

- You can look up the [Lox syntax](https://craftinginterpreters.com/the-lox-language.html)
- As of now, on input-ing source code, it will spit out the AST (not for all statements, but basic ones -> completed [Parsing expressions](https://craftinginterpreters.com/parsing-expressions.html))
