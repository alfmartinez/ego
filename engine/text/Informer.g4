// Informer.g4
grammar Informer;

// Tokens
NUMBER: [0-9]+;
IDENT: [a-zA-Z]+;
PUNCT: [\\.,:];
STRING: '"' (ESC|.)*? '"';
COMMENT: '[' .*? ']' -> skip;
WHITESPACE: [ \r\n\t]+ -> skip;

fragment
ESC: '\\"' | '\\\\';

// Rules
start : (definition)* EOF;

definition
   : rulebook                             # RulebookDeclaration
   ;

rulebook
   : designator 'is' 'a' 'rulebook' '.';

designator
   : IDENT 
   | IDENT designator;