### What is AST

Abstract Syntax Tree is kind of an intermediate data structure that is used to represent what the source code means..in this instance , the source code , using the lexer has been turned into tokens and with thte help of these tokens we buiild out the Abstarct Syntax Tree , while also checing the correctness of the source code along the way...thus it also kinda does static analysis of the source code too.

### What do parsers do ?

take source code as input and produce a data structure that represents the source code. 

- context free grammar is used by parser generators. context free grammar is a set of rules that describe how to correct sentences in a language.

- two main strategieds when parsing

  top-down or bottom up

  recursive descent, early, predictive are all top down

  we are writing: recursive descent -> "top down operator precedence" (Pratt Parser)

- ! Our parser is not bullet proof from not recognizing erroneous syntax adn currently has no formal proofs of correctness (this requieres study of theory)
  
- What does it mean to parse statements - speciically let statements correctly ?
  parser shourld pruduce an AST that correctl represents the information contained in the original let statement

- expressions vs statements ?
  this usually depends on the programming language. In Lemon, we are considering function literals as expressions. Some languages might nopt do so.

- **when building AST**: to hold the identifier of the binding, the 'x' in 'let x= 5;', we
  use the Identifier struct type. Notice that the Identifier structs implements the Expression 
  interface. but the identifier 'x' in a let statement does't produce a value (remember the definition
  of the expression vs statements for this book ... expression produces value and statement doesn't.....this definition varies but we are considering it to be true) ... so why do we put it in expression ?
  Answer: Just to keep things simple...because Identifiers in other parts of the monkey/lemon lang do produce values. for ex: let x = valueProducingIdentifier;
  
  Also: Identifiers in other parts of the program DO produce values. for ex: let x = valueProducingIdentifier;

  - after reading recursive descent parser pseudocode:
    - basic ideaa behind the recursive descent parser is evident in the pseudocode...the entry point is the parseProgram where it constructs the root node of Abstract Syntax Tree (AST) ...then the child nodes (statements) are built and then statements are built and this cycle goes on...
  

### Parsing expressions:

- Pratt Parsing/ Top down operator precedence parsing as alternative to parsers using context free grammar /BNF

 **Our interpreter won't have postfix operators in order to keep scope limited.**

### Statements, expressions & expression statement

In our current understanding in this book, expression produce value, statements do not. This can vary from language to language but for now we will consider it to be true.

- programs in lemon/monkey are made up of statements (let, return). now we are adding expression statements.

Expression Statement: a statement that consists of only single expression. 

`let x = 5;` -> this is a statement

`x + 10;` -> this is an expression statement


### TODO: Later
 Look more into expression vs statements
 We are currently doing prefix and infix parsing ... what about postfix parsing? implement later

 - also learn about buffers in detail in golang