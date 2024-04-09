
package gosql 

//types and constraints
import (
	"fmt" //formatting and printing
	"strings" //string manipulation
)

//ytacking location within source 
type location struct {
	line uint
	col uint
}

//custom type for representing sql keywords as strings
type keyword string

//sql keywords
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

//custom type for representing symbols 
type symbol string 

//sql symbols
const (
	semicolonSymbol symbol = ";"
	asteriskSymbol symbol = "*"
	commaSymbol symbol = ","
	leftparenSymbol symbol = "("
	rightparenSymbol symbol = ")"
)

//categorize tokens in SQL
type tokenKind uint 

//enumerated constants for different kinds of tokens, using 'iota' for automatic incrementation
const (
	keywordKind tokenkind = iota //starts at 0, increment by 1 for each line
	symbolKind // becomes 1
	identifierKind //2
	stringKind //3
	numericKind	 //4
)
//represent token (SQL text with an associated type and location)
type token struct{
	value string
	kind tokenKind
	loc location
}
//keep track of current position
type cursor struct{
	pointer uint //index of cur char
	loc location //line and column at current pointer position
}

//compare two tokens, check if same value and type (kind)
func(t *token) equals(other *token) bool{
	return t.value == other.value && t.kind == other.kind
}
//lexer as func signature
//returns token, cursor and boolean indicating success
type lexer func(string, cursor) (*token, cursor, bool)
