- two main strategieds when parsing

  top-down or bottom up

  recursive descent, early, predictive are all top down

  we are writing: recursive descent -> "top down operator precedence" (Pratt Parser)

- ! Our parser is not bullet proof from not recognizing erroneous syntax adn currently has no formal proofs of correctness (this requieres study of theory)
  
- What does it mean to parse statements - speciically let statements correctly ?
  parser shourld pruduce an AST that correctl represents the information contained in the original let statement

- expressions vs statements ?
  this usually depends on the programming language. In Lemon, we are considering function literals as expressions. Some languages might nopt do so.

- when building AST: to hold the identifier of the binding, the 'x' in 'let x= 5;', we
  use the Identifier struct type. Notice that the Identifier structs implements the Expression 
  interface. but the identifier 'x' in a let statement does't produce a value (remember the definition)
  of the expression vs statements for this book) ... so why do we put it in expression ?
  Answer: Just to keep things simple
  
  Also: Identifiers in other parts of the program DO produce values. for ex: let x = valueProducingIdentifier;
