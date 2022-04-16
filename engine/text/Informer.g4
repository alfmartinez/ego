// Informer.g4
grammar Informer;

// Tokens
NUMBER: DIGIT+;
PUNCT: [\\.,:];
STRING: '"' (ESC|.)*? '"';
COMMENT: '[' .*? ']\n' -> skip;
WHITESPACE: [ \r\t]+ -> skip;
EOL: [\n]+;
WORD: LETTER+;
ANY: . ;

fragment ESC: '\\"' | '\\\\';
fragment LETTER: [\p{L}];
fragment DIGIT: [0-9];

// Rules
start : (definition)+ EOF;

definition
   : activityDeclaration 
   | rulebookDeclaration 
   | actionDeclaration 
   | aliasDeclaration
   | title  
   ;

title: STRING EOL;

rulebookDeclaration
   : designator 'is' 'a' 'rulebook' '.' EOL  # GenericRulebook
   | designator 'is' ('an'|'a') WORD 'based' 'rulebook' '.' EOL # SpecificRulebook
   ;

activityDeclaration: designator 'is' 'an' 'activity' '.' EOL;

actionDeclaration: designator 'is' 'an' 'action' 'applying' 'to' 'a' WORD '.' EOL;

aliasDeclaration: 'Understand' STRING 'as' WORD '.' EOL;

designator: 
   designator WORD
   | designator '-' WORD
   | WORD;