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





// *** Entry Point ***


// tkes a source string and returns a list of tokens or an error.
func lex(source string) ([]*token, error){
    // empty list of tokens.
    tokens := []*token{}
    //  cursor to keep track of the current position in the source string.
    cur := cursor{}

lex:
    for cur.pointer < uint(len(source)){
        // a list of lexer functions to try in sequence.
        // Each lexer function recognizes different kinds of tokens.
        lexers := []lexer{lexKeyWord, lexSymbol, lexString, lexNumeric, lexIdentifier}
        
        // Iterate over each lexer function.
        for _, l := range lexers{
            // lex token using function
            // If successful, `ok` will be true, and `token` and `newCursor` will be initialized.
            if token, newCursor, ok := l(source, cur); ok{
                // Update the cursor to the new position after lexing the token.
                cur = newCursor

                // If a token was produced, add it to list of tokens.
                if token != nil{
                    tokens = append(tokens, token)
                }

                // Continue from the start of the lex label, skipping the rest of the lexer functions.
                continue lex
            }
        }

        // If no lexer function could lex a token, construct error message. Last successful token
        hint := ""
        if len(tokens) > 0{
            hint = " after " + tokens[len(tokens) - 1].value
        }
        return nil, fmt.Errorf("Unable to lex token%s at %d:%d", hint,cur.loc.line, cur.loc.col )
    }
    // Return list of lexed tokens and no error once entire source string has been lexed.
    return tokens, nil
}
