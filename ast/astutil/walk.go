// +build !appengine

package astutil

import (
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// WalkFunc is used in Walk to walk the AST
type WalkFunc func(interface{}) error

// Walk walks the ASTs associated with a statement list generated by parser.ParseSrc
// each expression and/or statement is passed to the WalkFunc function.
// If the WalkFunc returns an error the walk is aborted and the error is returned
func Walk(stmts []ast.Stmt, f WalkFunc) error {
	for _, stmt := range stmts {
		if err := walkStmt(stmt, f); err != nil {
			return err
		}
	}
	return nil
}

func walkExprs(exprs []ast.Expr, f WalkFunc) error {
	for _, exp := range exprs {
		if err := walkExpr(exp, f); err != nil {
			return err
		}
	}
	return nil
}

func walkStmt(stmt ast.Stmt, f WalkFunc) error {
	//short circuit out if there are no functions
	if stmt == nil || f == nil {
		return nil
	}
	if err := callFunc(stmt, f); err != nil {
		return err
	}
	switch stmt := stmt.(type) {
	case *ast.BreakStmt:
	case *ast.ContinueStmt:
	case *ast.ReturnStmt:
		return walkExprs(stmt.Exprs, f)
	case *ast.ExprStmt:
		return walkExpr(stmt.Expr, f)
	case *ast.VarStmt:
		return walkExprs(stmt.Exprs, f)
	case *ast.LetsStmt:
		if err := walkExprs(stmt.Rhss, f); err != nil {
			return err
		}
		return walkExprs(stmt.Lhss, f)
	case *ast.IfStmt:
		if err := walkExpr(stmt.If, f); err != nil {
			return err
		}
		if err := Walk(stmt.Then, f); err != nil {
			return err
		}
		if err := Walk(stmt.ElseIf, f); err != nil {
			return err
		}
		if err := Walk(stmt.Else, f); err != nil {
			return err
		}
	case *ast.TryStmt:
		if err := Walk(stmt.Try, f); err != nil {
			return err
		}
		if err := Walk(stmt.Catch, f); err != nil {
			return err
		}
		if err := Walk(stmt.Finally, f); err != nil {
			return err
		}
	case *ast.LoopStmt:
		if err := walkExpr(stmt.Expr, f); err != nil {
			return err
		}
		if err := Walk(stmt.Stmts, f); err != nil {
			return err
		}
	case *ast.ForStmt:
		if err := walkExpr(stmt.Value, f); err != nil {
			return err
		}
		if err := Walk(stmt.Stmts, f); err != nil {
			return err
		}
	case *ast.CForStmt:
		if err := walkExpr(stmt.Expr1, f); err != nil {
			return err
		}
		if err := walkExpr(stmt.Expr2, f); err != nil {
			return err
		}
		if err := walkExpr(stmt.Expr3, f); err != nil {
			return err
		}
		if err := Walk(stmt.Stmts, f); err != nil {
			return err
		}
	case *ast.ThrowStmt:
		if err := walkExpr(stmt.Expr, f); err != nil {
			return err
		}
	case *ast.ModuleStmt:
		if err := Walk(stmt.Stmts, f); err != nil {
			return err
		}
	case *ast.SwitchStmt:
		if err := walkExpr(stmt.Expr, f); err != nil {
			return err
		}
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				if err := Walk(ssd.Stmts, f); err != nil {
					return err
				}
				continue
			}
			if err := walkExpr(ss.(*ast.CaseStmt).Expr, f); err != nil {
				return err
			}
			if err := Walk(ss.(*ast.CaseStmt).Stmts, f); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown statement %v", reflect.TypeOf(stmt))
	}
	return nil
}

func walkExpr(expr ast.Expr, f WalkFunc) error {
	//short circuit out if there are no functions
	if expr == nil || f == nil {
		return nil
	}
	if err := callFunc(expr, f); err != nil {
		return err
	}
	switch expr := expr.(type) {
	case *ast.LenExpr:
	case *ast.NumberExpr:
	case *ast.IdentExpr:
	case *ast.MemberExpr:
		return walkExpr(expr.Expr, f)
	case *ast.StringExpr:
	case *ast.ItemExpr:
		if err := walkExpr(expr.Value, f); err != nil {
			return err
		}
		return walkExpr(expr.Index, f)
	case *ast.SliceExpr:
		if err := walkExpr(expr.Value, f); err != nil {
			return err
		}
		if err := walkExpr(expr.Begin, f); err != nil {
			return err
		}
		return walkExpr(expr.End, f)
	case *ast.ArrayExpr:
		return walkExprs(expr.Exprs, f)
	case *ast.MapExpr:
		for _, expr := range expr.MapExpr {
			if err := walkExpr(expr, f); err != nil {
				return err
			}
		}
	case *ast.DerefExpr:
		return walkExpr(expr.Expr, f)
	case *ast.AddrExpr:
		return walkExpr(expr.Expr, f)
	case *ast.UnaryExpr:
		return walkExpr(expr.Expr, f)
	case *ast.ParenExpr:
		return walkExpr(expr.SubExpr, f)
	case *ast.FuncExpr:
		return Walk(expr.Stmts, f)
	case *ast.AssocExpr:
		return walkExpr(expr.Lhs, f)
	case *ast.LetsExpr:
		if err := walkExprs(expr.Lhss, f); err != nil {
			return err
		}
		return walkExprs(expr.Rhss, f)
	case *ast.NewExpr:
	case *ast.BinOpExpr:
		if err := walkExpr(expr.Lhs, f); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, f)
	case *ast.ConstExpr:
	case *ast.AnonCallExpr:
		if err := walkExpr(expr.Expr, f); err != nil {
			return err
		}
		return walkExpr(&ast.CallExpr{Func: reflect.Value{}, SubExprs: expr.SubExprs, VarArg: expr.VarArg, Go: expr.Go}, f)
	case *ast.CallExpr:
		return walkExprs(expr.SubExprs, f)
	case *ast.TernaryOpExpr:
		if err := walkExpr(expr.Expr, f); err != nil {
			return err
		}
		if err := walkExpr(expr.Lhs, f); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, f)
	case *ast.MakeChanExpr:
		return walkExpr(expr.SizeExpr, f)
	case *ast.MakeExpr:
		if err := walkExpr(expr.LenExpr, f); err != nil {
			return err
		}
		return walkExpr(expr.CapExpr, f)
	case *ast.ChanExpr:
		if err := walkExpr(expr.Rhs, f); err != nil {
			return err
		}
		return walkExpr(expr.Lhs, f)
	default:
		return fmt.Errorf("unknown expression %v", reflect.TypeOf(expr))
	}
	return nil
}

func callFunc(x interface{}, f WalkFunc) error {
	if x == nil || f == nil {
		return nil
	}
	return f(x)
}
