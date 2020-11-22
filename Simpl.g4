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
	| fun = expr '(' arg = expr ')'						# callExpr
	| left = expr '/' right = expr						# divExpr
	| left = expr '*' right = expr						# multExpr
	| left = expr '-' right = expr						# subExpr
	| left = expr '+' right = expr						# addExpr
	| arg = expr '->' body = expr						# lambdaExpr
	| '//' expr											# commentExpr
	| '(' expr ')'										# parenExpr
	| 'var' id = expr '=' value = expr					# varExpr
	| 'if' con = expr 'then' t = expr 'else' f = expr	# ifElseExpr
	| BOOL												# boolExpr
	| ID												# idExpr
	| NUM												# numExpr
	| STR												# strExpr;