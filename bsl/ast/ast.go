package ast

import (
	"strings"

	"github.com/tsukanov-as/radar/bsl/tokens"
)

// Value ...
type Value interface{}

// PrepSymbol ...
type PrepSymbol = string

// Node ...
type Node interface {
	//Place() *Place
}

// Place ...
type Place struct {
	Pos     int
	Len     int
	BegLine int
	EndLine int
}

// Decl ...
type Decl interface {
	Node
}

// Stmt ...
type Stmt interface {
	Node
}

// Expr ...
type Expr interface {
	Node
}

// PrepInst ...
type PrepInst interface {
	Node
}

// PrepExpr ...
type PrepExpr interface {
	Node
}

// Module ...
type Module struct {
	Path      string
	Decls     []Decl
	Auto      []*Item
	Body      []Stmt
	Interface []Decl
	Comments  map[int]string
}

// Scope ...
type Scope struct {
	Outer   *Scope
	Objects map[string]*Item
	Auto    []*Item
}

// Insert ...
func (s *Scope) Insert(obj *Item) (alt *Item) {
	if alt = s.Objects[obj.Name]; alt == nil {
		s.Objects[obj.Name] = obj
	}
	return alt
}

// Find ...
func (s *Scope) Find(name string) *Item {
	return s.Objects[strings.ToLower(name)]
}

// NewScope ...
func NewScope(outer *Scope) *Scope {
	const n = 4
	return &Scope{outer, make(map[string]*Item, n), []*Item{}}
}

// Item ...
type Item struct {
	Name string
	Decl Decl
}

// Decls ...
type (
	// VarModListDecl ...
	VarModListDecl struct {
		Directive *tokens.Token
		List      []*VarModDecl
		Place
	}

	// VarModDecl ...
	VarModDecl struct {
		Name      string
		Directive *tokens.Token
		Export    bool
		Place
	}

	// VarLocDecl ...
	VarLocDecl struct {
		Name string
		Place
	}

	// ParamDecl ...
	ParamDecl struct {
		Name  string
		ByVal bool
		Value Expr
		Place
	}

	// MethodDecl ...
	MethodDecl struct {
		Sign Signature
		Vars []*VarLocDecl
		Auto []*Item
		Body []Stmt
		Place
	}

	// Signature ...
	Signature interface {
		// ProcDecl, FuncDecl
	}

	// ProcSign ...
	ProcSign struct {
		Name      string
		Directive *tokens.Token
		Params    []*ParamDecl
		Export    bool
		Place
	}

	// FuncSign ...
	FuncSign struct {
		Name      string
		Directive *tokens.Token
		Params    []*ParamDecl
		Export    bool
		Place
	}
)

// Expr ...
type (
	// BasicLitExpr ...
	BasicLitExpr struct {
		Kind  tokens.Token
		Value Value
		Place
	}

	// FieldExpr ...
	FieldExpr struct {
		Name string
		Args []Expr
		Place
	}

	// IndexExpr ...
	IndexExpr struct {
		Expr Expr
		Place
	}

	// CallExpr ...
	CallExpr struct {
		Args []Expr
		Place
	}

	// IdentExpr ...
	IdentExpr struct {
		Item *Item
		Tail []Expr
		Args []Expr
		Place
	}

	// UnaryExpr ...
	UnaryExpr struct {
		Operator tokens.Token
		Operand  Expr
		Place
	}

	// BinaryExpr ...
	BinaryExpr struct {
		Left     Expr
		Operator tokens.Token
		Right    Expr
		Place
	}

	// NewExpr ...
	NewExpr struct {
		Name *string
		Args []Expr
		Place
	}

	// TernaryExpr ...
	TernaryExpr struct {
		Cond Expr
		Then Expr
		Else Expr
		Tail []Expr
		Place
	}

	// ParenExpr ...
	ParenExpr struct {
		Expr Expr
		Place
	}

	// NotExpr ...
	NotExpr struct {
		Expr Expr
		Place
	}

	// StringExpr ...
	StringExpr struct {
		List []*BasicLitExpr
		Place
	}
)

// Stmt ...
type (
	// AssignStmt ...
	AssignStmt struct {
		Left  *IdentExpr
		Right Expr
		Place
	}

	// ReturnStmt ...
	ReturnStmt struct {
		Expr Expr
		Place
	}

	// BreakStmt ...
	BreakStmt struct {
		Place
	}

	// ContinueStmt ...
	ContinueStmt struct {
		Place
	}

	// RaiseStmt ...
	RaiseStmt struct {
		Expr Expr // TODO: nil
		Place
	}

	// ExecuteStmt ...
	ExecuteStmt struct {
		Expr Expr // TODO: nil
		Place
	}

	// CallStmt ...
	CallStmt struct {
		Ident *IdentExpr
		Place
	}

	// IfStmt ...
	IfStmt struct {
		Cond  Expr
		Then  []Stmt
		ElsIf *[]*ElsIfStmt
		Else  *ElseStmt
		Place
	}

	// ElseStmt ...
	ElseStmt struct {
		Body []Stmt
		Place
	}

	// ElsIfStmt ...
	ElsIfStmt struct {
		Cond Expr
		Then []Stmt
		Place
	}

	// WhileStmt ...
	WhileStmt struct {
		Cond Expr
		Body []Stmt
		Place
	}

	// ForStmt ...
	ForStmt struct {
		Ident *IdentExpr
		From  Expr
		To    Expr
		Body  []Stmt
		Place
	}

	// ForEachStmt ...
	ForEachStmt struct {
		Ident *IdentExpr
		In    Expr
		Body  []Stmt
		Place
	}

	// TryStmt ...
	TryStmt struct {
		Try    []Stmt
		Except ExceptStmt
		Place
	}

	// ExceptStmt ...
	ExceptStmt struct {
		Body []Stmt
		Place
	}

	// GotoStmt ...
	GotoStmt struct {
		Label string
		Place
	}

	// LabelStmt ...
	LabelStmt struct {
		Label string
		Place
	}
)

// PrepInst
type (
	// PrepIfInst ...
	PrepIfInst struct {
		Cond PrepExpr
		Place
	}

	// PrepElsIfInst ...
	PrepElsIfInst struct {
		Cond PrepExpr
		Place
	}

	// PrepElseInst ...
	PrepElseInst struct {
		Place
	}

	// PrepEndIfInst ...
	PrepEndIfInst struct {
		Place
	}

	// PrepRegionInst ...
	PrepRegionInst struct {
		Name string
		Place
	}

	// PrepEndRegionInst ...
	PrepEndRegionInst struct {
		Place
	}
)

// PrepExpr
type (
	// PrepBinaryExpr ...
	PrepBinaryExpr struct {
		Left     PrepExpr
		Operator tokens.Token
		Right    PrepExpr
		Place
	}

	// PrepNotExpr ...
	PrepNotExpr struct {
		Expr PrepExpr
		Place
	}

	// PrepSymExpr ...
	PrepSymExpr struct {
		Symbol PrepSymbol
		Exist  bool
		Place
	}
)
