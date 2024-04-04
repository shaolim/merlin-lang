package parser

import (
	"fmt"
	"github/shaolim/merlin-lang/ast"
	"github/shaolim/merlin-lang/lexer"
	"github/shaolim/merlin-lang/token"
)

type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  token.Token
	peerToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peerToken
	p.peerToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	fmt.Printf("program %+v", program)

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		result := p.parserLetStatement()
		fmt.Printf("curr token Let %+v \n", result)
		return result
	default:
		return nil
	}
}

func (p *Parser) parserLetStatement() *ast.LetStatement {

	stmt := &ast.LetStatement{Token: p.curToken}

	fmt.Printf("token literal %+v  --- %+v \n", stmt, stmt.TokenLiteral())

	if !p.expectPeek(token.IDENT) {
		fmt.Println("ident 1 ")
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		fmt.Println("ident 2")
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	if !p.curTokenIs(token.SEMICOLON) {
		fmt.Println("assign 3")
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peerToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}
