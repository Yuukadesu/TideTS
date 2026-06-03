grammar TideSQL;

// 最简 TideTS SQL：INSERT + SELECT（单测点、单设备路径）。
// 示例：
//   INSERT INTO root.sg1.d1(temperature) VALUES (100, 25.5);
//   SELECT temperature FROM root.sg1.d1 WHERE time >= 100 AND time <= 200 LIMIT 10;

options {
}

statement
    : insertStmt
    | selectStmt
    ;

insertStmt
    : INSERT INTO path LPAREN measurement RPAREN VALUES LPAREN timestamp COMMA value RPAREN
    ;

selectStmt
    : SELECT measurement FROM path whereClause? limitClause?
    ;

whereClause
    : WHERE timePredicate (AND timePredicate)*
    ;

timePredicate
    : TIME cmpOp INTEGER
    ;

cmpOp
    : GTE
    | LTE
    | GT
    | LT
    | EQ
    ;

limitClause
    : LIMIT INTEGER
    ;

path
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

measurement
    : IDENTIFIER
    ;

timestamp
    : INTEGER
    ;

value
    : BOOLEAN
    | INTEGER
    | FLOAT
    | STRING
    ;

INSERT   : I N S E R T ;
INTO     : I N T O ;
SELECT   : S E L E C T ;
FROM     : F R O M ;
WHERE    : W H E R E ;
AND      : A N D ;
VALUES   : V A L U E S ;
LIMIT    : L I M I T ;
TIME     : T I M E ;

GTE : '>=' ;
LTE : '<=' ;
GT  : '>' ;
LT  : '<' ;
EQ  : '=' ;

LPAREN : '(' ;
RPAREN : ')' ;
COMMA  : ',' ;
DOT    : '.' ;
SEMI   : ';' ;

BOOLEAN
    : 'true'
    | 'false'
    | 'TRUE'
    | 'FALSE'
    ;

INTEGER : [0-9]+ ;

FLOAT
    : [0-9]+ '.' [0-9]+
    ;

STRING
    : '\'' (~['\\\r\n] | '\\' .)* '\''
    ;

IDENTIFIER
    : [a-zA-Z_] [a-zA-Z0-9_]*
    ;

WS : [ \t\r\n]+ -> skip ;

fragment A : [aA] ;
fragment B : [bB] ;
fragment C : [cC] ;
fragment D : [dD] ;
fragment E : [eE] ;
fragment F : [fF] ;
fragment G : [gG] ;
fragment H : [hH] ;
fragment I : [iI] ;
fragment J : [jJ] ;
fragment K : [kK] ;
fragment L : [lL] ;
fragment M : [mM] ;
fragment N : [nN] ;
fragment O : [oO] ;
fragment P : [pP] ;
fragment Q : [qQ] ;
fragment R : [rR] ;
fragment S : [sS] ;
fragment T : [tT] ;
fragment U : [uU] ;
fragment V : [vV] ;
fragment W : [wW] ;
fragment X : [xX] ;
fragment Y : [yY] ;
fragment Z : [zZ] ;
