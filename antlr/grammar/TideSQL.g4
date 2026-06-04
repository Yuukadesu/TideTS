grammar TideSQL;

// TideTS SQL：INSERT / SELECT / DELETE / CREATE / SHOW。
// 示例：
//   INSERT INTO root.sg1.d1(temperature) VALUES (100, 25.5), (101, 26.0);
//   SELECT temperature FROM root.sg1.d1 WHERE time >= 100 AND time <= 200 LIMIT 10;
//   SELECT COUNT(temperature) FROM root.sg1.d1 WHERE time >= 100 AND time <= 200;
//   DELETE FROM root.sg1.d1(temperature) WHERE time >= 100 AND time <= 200;
//   CREATE TIMESERIES root.sg1.d1(temperature) WITH DATATYPE=DOUBLE;
//   SHOW DEVICES root.sg1.**;
//   SHOW TIMESERIES root.sg1.d1;

options {
}

statement
    : insertStmt
    | selectStmt
    | deleteStmt
    | createTimeseriesStmt
    | showDevicesStmt
    | showTimeseriesStmt
    ;

insertStmt
    : INSERT INTO path LPAREN measurement RPAREN VALUES valueRow (COMMA valueRow)*
    ;

valueRow
    : LPAREN timestamp COMMA value RPAREN
    ;

selectStmt
    : SELECT measurement FROM path whereClause? limitClause?
    | SELECT COUNT LPAREN measurement RPAREN FROM path whereClause?
    ;

deleteStmt
    : DELETE FROM path LPAREN measurement RPAREN whereClause
    ;

createTimeseriesStmt
    : CREATE TIMESERIES path LPAREN measurement RPAREN WITH DATATYPE EQ dataTypeName
    ;

showDevicesStmt
    : SHOW DEVICES showPattern?
    ;

showTimeseriesStmt
    : SHOW TIMESERIES path
    ;

showPattern
    : path (DOT STAR STAR)?
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

dataTypeName
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

INSERT     : I N S E R T ;
DELETE     : D E L E T E ;
INTO       : I N T O ;
SELECT     : S E L E C T ;
COUNT      : C O U N T ;
FROM       : F R O M ;
WHERE      : W H E R E ;
AND        : A N D ;
VALUES     : V A L U E S ;
LIMIT      : L I M I T ;
TIME       : T I M E ;
CREATE     : C R E A T E ;
TIMESERIES : T I M E S E R I E S ;
WITH       : W I T H ;
DATATYPE   : D A T A T Y P E ;
SHOW       : S H O W ;
DEVICES    : D E V I C E S ;

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
STAR   : '*' ;

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
