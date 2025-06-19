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
	ASSIGN = "="
	PLUS   = "PLUS"

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
)
