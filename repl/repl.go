package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // Creates a scanner to read input

	for {
		fmt.Fprintf(out, PROMPT)  // Output to screen
		scanned := scanner.Scan() // Returns false when no more tokens to scan
		if !scanned {             // Exit if false
			return
		}

		line := scanner.Text() // Gather last token -> string
		l := lexer.New(line)   // Create new lexer instance based on user input

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // Iterates through entire lexer input
			fmt.Fprintf(out, "%+v\n", tok) // Print out results
		}
	}

}
