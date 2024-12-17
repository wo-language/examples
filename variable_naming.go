// from src/go/scanner/scanner.go
// as you can see, they used a comment for every variable
type Scanner struct {
	// immutable state
	file *token.File  // source file handle
	dir  string       // directory portion of file.Name()
	src  []byte       // source
	err  ErrorHandler // error reporting; or nil
	mode Mode         // scanning mode

	// scanning state
	ch         rune      // current character
	offset     int       // character offset
	rdOffset   int       // reading offset (position after current character)
	lineOffset int       // current line offset
	insertSemi bool      // insert a semicolon before next newline
	nlPos      token.Pos // position of newline in preceding comment

	// public state - ok to modify
	ErrorCount int // number of errors encountered
}

// with Wo's preferred full variable names, those are eliminated like so
// as I was scanning through the code, I had to repeatedly scroll back to read the comments
// but this would have sufficed
type Scanner struct {
	// immutable state
	sourceFile          *token.File
  sourceFileDirectory string
	source              []byte
	err                 ErrorHandler // nullable
	scanningMode        Mode   

	// scanning state
	currentChar       rune
	charOffset        int
	readingOffset     int       // (position after current character)
	currentLineOffset int
	insertSemicolon   bool      //before next newline
	newLinePosition   token.Pos // pin preceding comment

	// public state - ok to modify
	ErrorCount        int
}
