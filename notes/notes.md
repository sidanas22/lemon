- two main strategieds when parsing

  top-down or bottom up

  recursive descent, early, predictive are all top down

  we are writing: recursive descent -> "top down operator precedence" (Pratt Parser)

- ! Our parser is not bullet proof from not recognizing erroneous syntax adn currently has no formal proofs of correctness (this requieres study of theory)
  
- What does it mean to parse statements - speciically let statements correctly ?
  parser shourld pruduce an AST that correctl represents the information contained in the original let statement

- expressions vs statements ?
  this usually depends on the programming language. In Lemon, we are considering function literals as expressions. Some languages might nopt do so.
