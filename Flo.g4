// Flo.g4
grammar Flo;

options {
    superClass=FloParserBase;
}

// Keywords
TRUE                   : 'true';
FALSE                  : 'false';
PRINT                  : 'print';
IF                     : 'if';
FOR                    : 'for';
ELSE                   : 'else';
ELIF                   : 'elif';
LOGICAL_OR             : 'or';
LOGICAL_AND            : 'and';
FUNC                   : 'func';
RETURN                 : 'return';

// Relational ops

EQUALS                 : '==';
NOT_EQUALS             : '!=';
LESS                   : '<';
LESS_OR_EQUALS         : '<=';
GREATER                : '>';
GREATER_OR_EQUALS      : '>=';

// Arithmetic ops

OR                     : '|';
MOD                    : '%';
LSHIFT                 : '<<';
RSHIFT                 : '>>';
BIT_CLEAR              : '&^';

// Unary ops
NOT                    : 'not';

// Other ops
MULTIPLY               : '*';
POWER                  : '**';
DIVIDE                 : '/';
PLUS_PLUS              : '++';
MINUS_MINUS            : '--';
ADDITION               : '+';
SUBTRACTION            : '-';

// Hidden tokens
WS                     : [ \t]+             -> channel(HIDDEN);
COMMENT                : '/*' .*? '*/'      -> channel(HIDDEN);
TERMINATOR             : [\r\n]+            -> channel(HIDDEN);
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

// Other tokens
L_CURLY                : '{';
R_CURLY                : '}';
L_BRACKET              : '[';
R_BRACKET              : ']';
L_PAREN                : '(';
R_PAREN                : ')';
SEMI                   : ';';
INTEGER                : [0-9]+;
CHARS                  : [']~[']*['];
FLOATING_POINT         : [0-9]+[.][0-9]*;
IDENTIFIER             : [_a-zA-Z]+[_a-zA-Z0-9]*;

// Rules
start : multi_stmts EOF
      ;

multi_stmts : (stmt eos)+
            ;

stmt : simple_stmt
     | if_stmt
     | for_stmt
     | func_decl
     | print_stmt
     | return_stmt
     | block
     ;

simple_stmt : expression
            | assign_stmt
            ;

assign_stmt : IDENTIFIER '=' expression
            ;

if_stmt     : IF condition_block ( ELIF condition_block )* ( ELSE '{' multi_stmts? '}' )?
            ;

for_stmt    : FOR for_clause_block
            ;

func_decl   : FUNC IDENTIFIER '(' parameters? ')' '{' multi_stmts? '}'
               ;

parameters  : IDENTIFIER ( ',' IDENTIFIER)*
            ;

for_clause_block  : simple_stmt SEMI expression SEMI simple_stmt '{' multi_stmts? '}'
            ;

condition_block: expression '{' multi_stmts? '}'
               ;

print_stmt  : PRINT expression
            ;

return_stmt : RETURN expression?
            ;

block : '{' multi_stmts? '}'
      ;

expression
   : '(' expression ')'                                      # ExpressionGroup
   | expression '(' expression_list? ')'                     # CallExpression
   | expression '**' expression                              # Power
   | op=('+'|'-') expression                                 # Unary
   | IDENTIFIER op=('++'|'--')                               # UnaryIncDec
   | expression L_BRACKET expression R_BRACKET               # GetItem
   | expression op=('*'|'/'|'%') expression                  # MulDiv
   | expression op=('+'|'-') expression                      # AddSub
   | expression op=('<<'|'>>') expression                    # BitShift
   | expression '&' expression                               # BitwiseAnd
   | expression '^' expression                               # XOR
   | expression '|' expression                               # BitwiseOr
   | expression op=('<'|'<='|'>'|'>='|'=='|'!=') expression  # Relational
   | NOT expression                                          # Not
   | expression LOGICAL_AND expression                       # And
   | expression LOGICAL_OR expression                        # Or
   | L_BRACKET expression_list? R_BRACKET                    # List
   | INTEGER                                                 # Integer
   | FLOATING_POINT                                          # Float
   | IDENTIFIER                                              # ReadIdentifier
   | TRUE                                                    # Boolean
   | FALSE                                                   # Boolean
   | CHARS                                                   # String
   ;

expression_list : expression ( ',' expression )* 
                ;

eos
    : ';'
    | EOF
    | {p.lineTerminatorAhead()}?
    | {p.checkPreviousTokenText("}")}?
    ;