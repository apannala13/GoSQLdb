// *** Imports ***
package gosql 


import (
	"fmt" 
	"strings"
)

// *** App struct and methods ***
type location struct {
	line uint
	col uint
}

type keyword string

const (
	selectKeyword keyword = "select"
	fromKeyWord keyword = "from"
	asKeyWord keyword = "as"
	tableKeyword keyword = "table"
	createKeyWord keyword = "create"
	insertKeyWord keyword = "insert"
	intoKeyWord keyword = "into"
	valuesKeyWord keyword = "values"
	intKeyWord keyword = "int"
	textKeyWord keyword = "text"
)

 
type symbol string 

const (
	semicolonSymbol symbol = ";"
	asteriskSymbol symbol = "*"
	commaSymbol symbol = ","
	leftparenSymbol symbol = "("
	rightparenSymbol symbol = ")"
)

type tokenKind uint 

//enumerated constants for different kinds of tokens, using 'iota' for automatic incrementation
const (
	keywordKind tokenkind = iota 
	symbolKind 
	identifierKind 
	stringKind 
	numericKind	 
)
type token struct{
	value string
	kind tokenKind
	loc location
}
type cursor struct{
	pointer uint 
	loc location 
}

func(t *token) equals(other *token) bool{
	return t.value == other.value && t.kind == other.kind
}
//lexer as func signature
//returns token, cursor and boolean indicating success
type lexer func(string, cursor) (*token, cursor, bool)
