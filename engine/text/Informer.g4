// Informer.g4
grammar Informer;

// Tokens
WORD: [\p{L}-]+;
PUNCT: [\\.];
COMMENT: '[' .*? ']\n' -> skip;
WHITESPACE: [ \r\t]+ -> skip;
EOL: [\n]+;



// Rules
start : statement+;

statement: 
   definition '.'? EOL;

definition: 
   identifier 'is a rulebook' # Rulebook
   | identifier 'is an object based rulebook' # Rulebook
   | identifier 'is an activity' # Activity;

identifier:
   WORD
   | WORD identifier;