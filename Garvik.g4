grammar Garvik;

BOOL: 'true' | 'false'; // match boolean
ID: [a-zA-Z_]+; // match identifiers
NUM: [0-9]+; // match integers
STR:
	'"' ('\\' [btnfr"'\\] | ~[\r\n\\"])* '"'; // match string literal
WS: [ \t\r\n]+ -> skip; // Whitespace

program: expr*;

expr:
	left = expr '==' right = expr								# equalExpr
	| left = expr '<' right = expr								# lessExpr
	| left = expr '>' right = expr								# greaterExpr
	| id = expr '.' field = expr '=' value = expr				# dotAssignExpr
	| id = expr '.' field = expr								# dotExpr
	| id = expr '[' key = expr ']' '=' value = expr				# lookupAssignExpr
	| id = expr '[' key = expr ']'								# lookupExpr
	| fun = expr '(' arg = expr ')'								# callExpr
	| left = expr '/' right = expr								# divExpr
	| left = expr '*' right = expr								# multExpr
	| left = expr '-' right = expr								# subExpr
	| left = expr '+' right = expr								# addExpr
	| param = expr '->' body = expr								# lambdaExpr
	| id = expr ':' field = expr								# fieldExpr
	| id = expr '=' value = expr								# assignExpr
	| '//' expr													# commentExpr
	| '[' expr (',' expr)* ']'									# listExpr
	| '(' expr ')'												# parenExpr
	| '{' expr* '}'												# structExpr
	| 'if' con = expr 'then' t = expr 'else' f = expr			# ifElseExpr
	| 'let' id = expr '=' value = expr 'in' expression = expr	# letExpr
	| 'len' id = expr											# lenExpr
	| BOOL														# boolExpr
	| ID														# idExpr
	| NUM														# numExpr
	| STR														# strExpr;