#Doc

- db in GOLang
- implement parsor to support CREATE, INSERT, SELECT queries
- in memory backend supporting TEXT and INT types 



1. Map SQL source into list of tokens (lexing). Call parse functions 
to find indvidual SQL statements (e.g. SELECT)
2. Parse functions will call helper functions to find patterns of 
recursively parseable chunks, keywords, symbols, identifiers and string/numeric literals
- db in GOLang
- run CREATE, INSERT, SELEimplement parsor to supportqueries
- in memory backend supporting TEXT and INT types 

Process:
1. Map SQL source into list of tokens (lexing)
2. Call parse functions => find individual SQL statements (e.g. SELECT)
3. Parse functions will call own helper functions to find patterns of recursively 
parseable chunks, keywords, symbols, identifiers, numeric/string literals
4. In memory backend to perform ops based on AST. 
5. REPL to accept SQL from CLI and pass to backend to return results








