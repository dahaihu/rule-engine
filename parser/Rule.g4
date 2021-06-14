grammar Rule;

WHITESPACE
    : [ \r\n\t]+ -> skip
    ;

INT
    : [0-9]+
    ;

FLOAT
    : INT '.' [0-9]+
    ;

STRING
   : '"' [a-zA-Z] [a-zA-Z0-9]+ '"'
   ;

AND
    :'and'
    ;
OR
    : 'or'
    ;
MUL
    : '*'
    ;
DIV
    : '/'
    ;
ADD
    : '+'
    ;
SUB
    : '-'
    ;
EQ
    : '=='
    ;
NE
    : '!='
    ;
GT
    : '>'
    ;
LT
    : '<'
    ;
GE
    : '>='
    ;
LE
    : '<='
    ;
LEFTBRACKET
    : '('
    ;
RIGHTBRACKET
    : ')'
    ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_'  | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

start
    : logical EOF;

arithmetic
    : INT                                               # Int
    | FLOAT                                             # Float
    | attrPath                                          # Attr
    | SUB arithmetic                                    # NegativeValue
    | left=arithmetic op=(MUL | DIV) right=arithmetic   # MulDiv
    | left=arithmetic op=(ADD | SUB) right=arithmetic   # AddSub
    | LEFTBRACKET arithmetic RIGHTBRACKET               # ArithmeticWithParenthesis
    ;

compare
    : LEFTBRACKET compare RIGHTBRACKET # CompareWithParenthesis
    | left=arithmetic op=(EQ|NE|GT|LT|GE|LE) right=arithmetic # BasicCompare
    ;

logical
    : compare #BasicLogical
    | left=logical op=(AND|OR) right=logical #TwoLogical
    | LEFTBRACKET logical RIGHTBRACKET # LogicalWithParenthesis
    ;