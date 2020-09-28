grammar Simpl;

ID: [a-z]+; // match lower-case identifiers
NUM: [0-9]+; // match integers
MUL: '*'; // Multiplication
ADD: '+'; // Addition
RPAREN: '('; // Right parentheses
LPAREN: ')'; // Left parentheses
LET: 'let'; // Keyword let
EQ: '='; // Binding operator
IN: 'in'; // Keyword in
END: 'end'; // Keyword end
WS: [ \t\r\n]+ -> skip;
// Whitespace

program: expr EOF;
// a program is just an expression, ending with the end of the input
expr:
	ID															# idExpr // an identifier
	| NUM														# numExpr // a number
	| '(' expr ')'												# parenExpr
	| left = expr '*' right = expr								# multExpr
	| left = expr '+' right = expr								# addExpr
	| 'let' ID '=' value = expr 'in' expression = expr 'end'	# letExpr;