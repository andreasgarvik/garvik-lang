grammar Simpl;

BOOL: 'true' | 'false'; // match boolean
ID: [a-z]+; // match identifiers
NUM: [0-9]+; // match integers
STR:
	'"' ('\\' [btnfr"'\\] | ~[\r\n\\"])* '"'; // match string literal
WS: [ \t\r\n]+ -> skip; // Whitespace

program: expr*;

expr:
	left = expr '==' right = expr						# equalExpr
	| id = expr '[' key = expr ']'						# lookupExpr
	| fun = expr '(' arg = expr ')'						# callExpr
	| left = expr '/' right = expr						# divExpr
	| left = expr '*' right = expr						# multExpr
	| left = expr '-' right = expr						# subExpr
	| left = expr '+' right = expr						# addExpr
	| arg = expr '->' body = expr						# lambdaExpr
	| id = expr '=' value = expr						# assignExpr
	| id = expr '.' field = expr						# dotExpr
	| left = expr ',' right = expr						# commaExpr
	| '//' expr											# commentExpr
	| '[' expr ']'										# listExpr
	| '(' expr ')'										# parenExpr
	| '{' expr* '}'										# structExpr
	| 'if' con = expr 'then' t = expr 'else' f = expr	# ifElseExpr
	| BOOL												# boolExpr
	| ID												# idExpr
	| NUM												# numExpr
	| STR												# strExpr;