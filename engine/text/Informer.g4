// Informer.g4
grammar Informer;

// Tokens
ARTICLE: ('an'|'a');
GERUND: [\p{L}]'ing';
WORD: [\p{L}-]+;
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
   ARTICLE 'rulebook'                                    # Rulebook
   | ARTICLE designator 'based' 'rulebook'               # ObjectBasedRulebook
   | ARTICLE 'activity'                                  # Activity
   | certainty designator                                # CertaintyOfAttribute
   | ARTICLE 'kind' 'of' designator                      # ObjectKind
   | ARTICLE 'kind' 'of' 'value' 'with' 'values' values  # ValueKind ;

certainty:
   'usually' | 'always' | 'never';

designator:
   WORD
   | ARTICLE
   | GERUND
   | 'of'
   | designator designator;

values:
   WORD 
   | values ',' values
   | values 'and' values;