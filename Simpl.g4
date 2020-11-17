grammar Simpl;

BOOL: 'true' | 'false'; // match boolean
ID: [a-z]+; // match identifiers
NUM: [0-9]+; // match integers
STR:
	'"' ('\\' [btnfr"'\\] | ~[\r\n\\"])* '"'; // match string literal
WS: [ \t\r\n]+ -> skip; // Whitespace

program: expr EOF;
// a program is just an expression, ending with the end of the input
expr:
	BOOL																		# boolExpr
	| ID																		# idExpr
	| NUM																		# numExpr
	| STR																		# strExpr
	| '(' expr ')'																# parenExpr
	| left = expr '==' right = expr												# equalExpr
	| fun = expr '(' arg = expr ')'												# applyExpr
	| left = expr '/' right = expr												# divExpr
	| left = expr '*' right = expr												# multExpr
	| left = expr '-' right = expr												# subExpr
	| left = expr '+' right = expr												# addExpr
	| arg = expr '->' body = expr												# lamdaExpr
	| 'if' con = expr 'then' t = expr 'else' f = expr							# ifElseExpr
	| 'let' id = expr '(' arg = expr ')' '=' body = expr 'in' call = expr 'end'	# funExpr
	| 'let' id = expr '=' value = expr 'in' expression = expr 'end'				# letExpr;