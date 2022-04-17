// Informer.g4
grammar Informer;

// Tokens
ARTICLE: ('an'|'a');
WORD: [\p{L}]+;
PUNCT: [\\.,:-];
COMMENT: '[' .*? ']\n' -> skip;
WHITESPACE: [ \r\t]+ -> skip;
EOL: [\n]+;


// Rules
start : (statement '.'? EOL)+ EOF;

statement: 
   rulebook
   | activity;

rulebook:
   designator 'is' ARTICLE 'rulebook'
   | designator 'is' ARTICLE designator 'based' 'rulebook';

activity:
   designator 'is' ARTICLE 'activity';

designator:
   WORD
   | ARTICLE
   | designator '-' designator
   | designator designator;
