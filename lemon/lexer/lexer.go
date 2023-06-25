package lexer

import "go/token"

type Lexer struct {
	input        string
	position     int  //current position in input (points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) NextToken() token.Token{

}

func New(input string) *Lexer {
	l := &Lexer{input: input}
    l.readChar()
	return l
}

// TODO: add support for utf,unicode

func (l *Lexer) readChar(){
    if l.readPosition >= len(l.input){
        // 0 is ASCII code for "NULL"
        l.ch = 0
    } else{
        l.ch = l.input[l.readPosition] 
    } 

    l.position = l.readPosition
    l.readPosition += 1
}
