package token

type TokenType string

// Token 词法单元
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 未知的类型
	ILLEGAL = "ILLEGAL"
	// 文件结尾
	EOF = "EOF"

	// 标识符和字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func New(tokenType TokenType, ch byte) Token {
	return Token{
		tokenType,
		string(ch),
	}
}
