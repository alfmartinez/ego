// Informer.g4
grammar Informer;

// Tokens
ARTICLE: ('an'|'a');
GERUND: [\p{L}]'ing';
WORD: [\p{L}]+;
PUNCT: [\\.,:-];
COMMENT: '[' .*? ']\n' -> skip;
WHITESPACE: [ \r\t]+ -> skip;
EOL: [\n]+;


// Rules
start : (statement '.'? EOL)+ EOF;

statement: 
   definition;

definition: 
   designator 'is' definitionType;

definitionType: 
   ARTICLE 'rulebook' 
   | ARTICLE designator 'based' 'rulebook' 
   | ARTICLE 'activity'
   | certainty designator 
   | ARTICLE 'kind' 'of' designator; 

certainty:
   'usually' | 'always' | 'never';

designator:
   WORD
   | ARTICLE
   | GERUND
   | 'of'
   | designator '-' designator
   | designator designator;
