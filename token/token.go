package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//INTEGER + IDENTIFIER
	IDENT = "IDENT"
	INT   = "INT"

	//OPERATOR
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LT     = "<"
	GT     = ">"

	EQ     = "=="
	NOT_EQ = "!="

	//KEY WORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	WHILE = "WHILE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,

	"while": WHILE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// 区分关键字和标识符
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
