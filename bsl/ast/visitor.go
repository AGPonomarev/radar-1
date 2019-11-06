package ast

type Visitor struct {
	VisitModule            []InterfaceVisitModule
	LeaveModule            []InterfaceLeaveModule
	VisitVarModListDecl    []InterfaceVisitVarModListDecl
	LeaveVarModListDecl    []InterfaceLeaveVarModListDecl
	VisitVarModDecl        []InterfaceVisitVarModDecl
	VisitVarLocDecl        []InterfaceVisitVarLocDecl
	VisitAutoDecl          []InterfaceVisitAutoDecl
	VisitParamDecl         []InterfaceVisitParamDecl
	LeaveParamDecl         []InterfaceLeaveParamDecl
	VisitMethodDecl        []InterfaceVisitMethodDecl
	LeaveMethodDecl        []InterfaceLeaveMethodDecl
	VisitProcSign          []InterfaceVisitProcSign
	LeaveProcSign          []InterfaceLeaveProcSign
	VisitFuncSign          []InterfaceVisitFuncSign
	LeaveFuncSign          []InterfaceLeaveFuncSign
	VisitBasicLitExpr      []InterfaceVisitBasicLitExpr
	VisitFieldExpr         []InterfaceVisitFieldExpr
	LeaveFieldExpr         []InterfaceLeaveFieldExpr
	VisitIndexExpr         []InterfaceVisitIndexExpr
	LeaveIndexExpr         []InterfaceLeaveIndexExpr
	VisitIdentExpr         []InterfaceVisitIdentExpr
	LeaveIdentExpr         []InterfaceLeaveIdentExpr
	VisitUnaryExpr         []InterfaceVisitUnaryExpr
	LeaveUnaryExpr         []InterfaceLeaveUnaryExpr
	VisitBinaryExpr        []InterfaceVisitBinaryExpr
	LeaveBinaryExpr        []InterfaceLeaveBinaryExpr
	VisitNewExpr           []InterfaceVisitNewExpr
	LeaveNewExpr           []InterfaceLeaveNewExpr
	VisitTernaryExpr       []InterfaceVisitTernaryExpr
	LeaveTernaryExpr       []InterfaceLeaveTernaryExpr
	VisitParenExpr         []InterfaceVisitParenExpr
	LeaveParenExpr         []InterfaceLeaveParenExpr
	VisitNotExpr           []InterfaceVisitNotExpr
	LeaveNotExpr           []InterfaceLeaveNotExpr
	VisitStringExpr        []InterfaceVisitStringExpr
	LeaveStringExpr        []InterfaceLeaveStringExpr
	VisitAssignStmt        []InterfaceVisitAssignStmt
	LeaveAssignStmt        []InterfaceLeaveAssignStmt
	VisitReturnStmt        []InterfaceVisitReturnStmt
	LeaveReturnStmt        []InterfaceLeaveReturnStmt
	VisitBreakStmt         []InterfaceVisitBreakStmt
	VisitContinueStmt      []InterfaceVisitContinueStmt
	VisitRaiseStmt         []InterfaceVisitRaiseStmt
	LeaveRaiseStmt         []InterfaceLeaveRaiseStmt
	VisitExecuteStmt       []InterfaceVisitExecuteStmt
	LeaveExecuteStmt       []InterfaceLeaveExecuteStmt
	VisitCallStmt          []InterfaceVisitCallStmt
	LeaveCallStmt          []InterfaceLeaveCallStmt
	VisitIfStmt            []InterfaceVisitIfStmt
	LeaveIfStmt            []InterfaceLeaveIfStmt
	VisitElsIfStmt         []InterfaceVisitElsIfStmt
	LeaveElsIfStmt         []InterfaceLeaveElsIfStmt
	VisitElseStmt          []InterfaceVisitElseStmt
	LeaveElseStmt          []InterfaceLeaveElseStmt
	VisitWhileStmt         []InterfaceVisitWhileStmt
	LeaveWhileStmt         []InterfaceLeaveWhileStmt
	VisitForStmt           []InterfaceVisitForStmt
	LeaveForStmt           []InterfaceLeaveForStmt
	VisitForEachStmt       []InterfaceVisitForEachStmt
	LeaveForEachStmt       []InterfaceLeaveForEachStmt
	VisitTryStmt           []InterfaceVisitTryStmt
	LeaveTryStmt           []InterfaceLeaveTryStmt
	VisitExceptStmt        []InterfaceVisitExceptStmt
	LeaveExceptStmt        []InterfaceLeaveExceptStmt
	VisitGotoStmt          []InterfaceVisitGotoStmt
	VisitLabelStmt         []InterfaceVisitLabelStmt
	VisitPrepIfInst        []InterfaceVisitPrepIfInst
	LeavePrepIfInst        []InterfaceLeavePrepIfInst
	VisitPrepElsIfInst     []InterfaceVisitPrepElsIfInst
	LeavePrepElsIfInst     []InterfaceLeavePrepElsIfInst
	VisitPrepElseInst      []InterfaceVisitPrepElseInst
	VisitPrepEndIfInst     []InterfaceVisitPrepEndIfInst
	VisitPrepRegionInst    []InterfaceVisitPrepRegionInst
	VisitPrepEndRegionInst []InterfaceVisitPrepEndRegionInst
	VisitPrepExpr          []InterfaceVisitPrepExpr
	LeavePrepExpr          []InterfaceLeavePrepExpr
	VisitPrepBinaryExpr    []InterfaceVisitPrepBinaryExpr
	LeavePrepBinaryExpr    []InterfaceLeavePrepBinaryExpr
	VisitPrepNotExpr       []InterfaceVisitPrepNotExpr
	LeavePrepNotExpr       []InterfaceLeavePrepNotExpr
	VisitPrepSymExpr       []InterfaceVisitPrepSymExpr
	VisitPrepParenExpr     []InterfaceVisitPrepParenExpr
	LeavePrepParenExpr     []InterfaceLeavePrepParenExpr
}

func (v *Visitor) HookUp(plugins []interface{}) {

	for _, p := range plugins {
		if tp, ok := p.(InterfaceVisitModule); ok {
			v.VisitModule = append(v.VisitModule, tp)
		}
		if tp, ok := p.(InterfaceLeaveModule); ok {
			v.LeaveModule = append(v.LeaveModule, tp)
		}
		if tp, ok := p.(InterfaceVisitVarModListDecl); ok {
			v.VisitVarModListDecl = append(v.VisitVarModListDecl, tp)
		}
		if tp, ok := p.(InterfaceLeaveVarModListDecl); ok {
			v.LeaveVarModListDecl = append(v.LeaveVarModListDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitVarModDecl); ok {
			v.VisitVarModDecl = append(v.VisitVarModDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitVarLocDecl); ok {
			v.VisitVarLocDecl = append(v.VisitVarLocDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitAutoDecl); ok {
			v.VisitAutoDecl = append(v.VisitAutoDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitParamDecl); ok {
			v.VisitParamDecl = append(v.VisitParamDecl, tp)
		}
		if tp, ok := p.(InterfaceLeaveParamDecl); ok {
			v.LeaveParamDecl = append(v.LeaveParamDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitMethodDecl); ok {
			v.VisitMethodDecl = append(v.VisitMethodDecl, tp)
		}
		if tp, ok := p.(InterfaceLeaveMethodDecl); ok {
			v.LeaveMethodDecl = append(v.LeaveMethodDecl, tp)
		}
		if tp, ok := p.(InterfaceVisitProcSign); ok {
			v.VisitProcSign = append(v.VisitProcSign, tp)
		}
		if tp, ok := p.(InterfaceLeaveProcSign); ok {
			v.LeaveProcSign = append(v.LeaveProcSign, tp)
		}
		if tp, ok := p.(InterfaceVisitFuncSign); ok {
			v.VisitFuncSign = append(v.VisitFuncSign, tp)
		}
		if tp, ok := p.(InterfaceLeaveFuncSign); ok {
			v.LeaveFuncSign = append(v.LeaveFuncSign, tp)
		}
		if tp, ok := p.(InterfaceVisitBasicLitExpr); ok {
			v.VisitBasicLitExpr = append(v.VisitBasicLitExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitFieldExpr); ok {
			v.VisitFieldExpr = append(v.VisitFieldExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveFieldExpr); ok {
			v.LeaveFieldExpr = append(v.LeaveFieldExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitIndexExpr); ok {
			v.VisitIndexExpr = append(v.VisitIndexExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveIndexExpr); ok {
			v.LeaveIndexExpr = append(v.LeaveIndexExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitIdentExpr); ok {
			v.VisitIdentExpr = append(v.VisitIdentExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveIdentExpr); ok {
			v.LeaveIdentExpr = append(v.LeaveIdentExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitUnaryExpr); ok {
			v.VisitUnaryExpr = append(v.VisitUnaryExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveUnaryExpr); ok {
			v.LeaveUnaryExpr = append(v.LeaveUnaryExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitBinaryExpr); ok {
			v.VisitBinaryExpr = append(v.VisitBinaryExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveBinaryExpr); ok {
			v.LeaveBinaryExpr = append(v.LeaveBinaryExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitNewExpr); ok {
			v.VisitNewExpr = append(v.VisitNewExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveNewExpr); ok {
			v.LeaveNewExpr = append(v.LeaveNewExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitTernaryExpr); ok {
			v.VisitTernaryExpr = append(v.VisitTernaryExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveTernaryExpr); ok {
			v.LeaveTernaryExpr = append(v.LeaveTernaryExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitParenExpr); ok {
			v.VisitParenExpr = append(v.VisitParenExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveParenExpr); ok {
			v.LeaveParenExpr = append(v.LeaveParenExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitNotExpr); ok {
			v.VisitNotExpr = append(v.VisitNotExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveNotExpr); ok {
			v.LeaveNotExpr = append(v.LeaveNotExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitStringExpr); ok {
			v.VisitStringExpr = append(v.VisitStringExpr, tp)
		}
		if tp, ok := p.(InterfaceLeaveStringExpr); ok {
			v.LeaveStringExpr = append(v.LeaveStringExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitAssignStmt); ok {
			v.VisitAssignStmt = append(v.VisitAssignStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveAssignStmt); ok {
			v.LeaveAssignStmt = append(v.LeaveAssignStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitReturnStmt); ok {
			v.VisitReturnStmt = append(v.VisitReturnStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveReturnStmt); ok {
			v.LeaveReturnStmt = append(v.LeaveReturnStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitBreakStmt); ok {
			v.VisitBreakStmt = append(v.VisitBreakStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitContinueStmt); ok {
			v.VisitContinueStmt = append(v.VisitContinueStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitRaiseStmt); ok {
			v.VisitRaiseStmt = append(v.VisitRaiseStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveRaiseStmt); ok {
			v.LeaveRaiseStmt = append(v.LeaveRaiseStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitExecuteStmt); ok {
			v.VisitExecuteStmt = append(v.VisitExecuteStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveExecuteStmt); ok {
			v.LeaveExecuteStmt = append(v.LeaveExecuteStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitCallStmt); ok {
			v.VisitCallStmt = append(v.VisitCallStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveCallStmt); ok {
			v.LeaveCallStmt = append(v.LeaveCallStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitIfStmt); ok {
			v.VisitIfStmt = append(v.VisitIfStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveIfStmt); ok {
			v.LeaveIfStmt = append(v.LeaveIfStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitElsIfStmt); ok {
			v.VisitElsIfStmt = append(v.VisitElsIfStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveElsIfStmt); ok {
			v.LeaveElsIfStmt = append(v.LeaveElsIfStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitElseStmt); ok {
			v.VisitElseStmt = append(v.VisitElseStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveElseStmt); ok {
			v.LeaveElseStmt = append(v.LeaveElseStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitWhileStmt); ok {
			v.VisitWhileStmt = append(v.VisitWhileStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveWhileStmt); ok {
			v.LeaveWhileStmt = append(v.LeaveWhileStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitForStmt); ok {
			v.VisitForStmt = append(v.VisitForStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveForStmt); ok {
			v.LeaveForStmt = append(v.LeaveForStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitForEachStmt); ok {
			v.VisitForEachStmt = append(v.VisitForEachStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveForEachStmt); ok {
			v.LeaveForEachStmt = append(v.LeaveForEachStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitTryStmt); ok {
			v.VisitTryStmt = append(v.VisitTryStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveTryStmt); ok {
			v.LeaveTryStmt = append(v.LeaveTryStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitExceptStmt); ok {
			v.VisitExceptStmt = append(v.VisitExceptStmt, tp)
		}
		if tp, ok := p.(InterfaceLeaveExceptStmt); ok {
			v.LeaveExceptStmt = append(v.LeaveExceptStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitGotoStmt); ok {
			v.VisitGotoStmt = append(v.VisitGotoStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitLabelStmt); ok {
			v.VisitLabelStmt = append(v.VisitLabelStmt, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepIfInst); ok {
			v.VisitPrepIfInst = append(v.VisitPrepIfInst, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepIfInst); ok {
			v.LeavePrepIfInst = append(v.LeavePrepIfInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepElsIfInst); ok {
			v.VisitPrepElsIfInst = append(v.VisitPrepElsIfInst, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepElsIfInst); ok {
			v.LeavePrepElsIfInst = append(v.LeavePrepElsIfInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepElseInst); ok {
			v.VisitPrepElseInst = append(v.VisitPrepElseInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepEndIfInst); ok {
			v.VisitPrepEndIfInst = append(v.VisitPrepEndIfInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepRegionInst); ok {
			v.VisitPrepRegionInst = append(v.VisitPrepRegionInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepEndRegionInst); ok {
			v.VisitPrepEndRegionInst = append(v.VisitPrepEndRegionInst, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepExpr); ok {
			v.VisitPrepExpr = append(v.VisitPrepExpr, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepExpr); ok {
			v.LeavePrepExpr = append(v.LeavePrepExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepBinaryExpr); ok {
			v.VisitPrepBinaryExpr = append(v.VisitPrepBinaryExpr, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepBinaryExpr); ok {
			v.LeavePrepBinaryExpr = append(v.LeavePrepBinaryExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepNotExpr); ok {
			v.VisitPrepNotExpr = append(v.VisitPrepNotExpr, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepNotExpr); ok {
			v.LeavePrepNotExpr = append(v.LeavePrepNotExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepSymExpr); ok {
			v.VisitPrepSymExpr = append(v.VisitPrepSymExpr, tp)
		}
		if tp, ok := p.(InterfaceVisitPrepParenExpr); ok {
			v.VisitPrepParenExpr = append(v.VisitPrepParenExpr, tp)
		}
		if tp, ok := p.(InterfaceLeavePrepParenExpr); ok {
			v.LeavePrepParenExpr = append(v.LeavePrepParenExpr, tp)
		}
	}

}

type InterfaceVisitModule interface {
	VisitModule(Node)
}
type InterfaceLeaveModule interface {
	LeaveModule(Node)
}
type InterfaceVisitVarModListDecl interface {
	VisitVarModListDecl(Node)
}
type InterfaceLeaveVarModListDecl interface {
	LeaveVarModListDecl(Node)
}
type InterfaceVisitVarModDecl interface {
	VisitVarModDecl(Node)
}
type InterfaceVisitVarLocDecl interface {
	VisitVarLocDecl(Node)
}
type InterfaceVisitAutoDecl interface {
	VisitAutoDecl(Node)
}
type InterfaceVisitParamDecl interface {
	VisitParamDecl(Node)
}
type InterfaceLeaveParamDecl interface {
	LeaveParamDecl(Node)
}
type InterfaceVisitMethodDecl interface {
	VisitMethodDecl(Node)
}
type InterfaceLeaveMethodDecl interface {
	LeaveMethodDecl(Node)
}
type InterfaceVisitProcSign interface {
	VisitProcSign(Node)
}
type InterfaceLeaveProcSign interface {
	LeaveProcSign(Node)
}
type InterfaceVisitFuncSign interface {
	VisitFuncSign(Node)
}
type InterfaceLeaveFuncSign interface {
	LeaveFuncSign(Node)
}
type InterfaceVisitBasicLitExpr interface {
	VisitBasicLitExpr(Node)
}
type InterfaceVisitFieldExpr interface {
	VisitFieldExpr(Node)
}
type InterfaceLeaveFieldExpr interface {
	LeaveFieldExpr(Node)
}
type InterfaceVisitIndexExpr interface {
	VisitIndexExpr(Node)
}
type InterfaceLeaveIndexExpr interface {
	LeaveIndexExpr(Node)
}
type InterfaceVisitIdentExpr interface {
	VisitIdentExpr(Node)
}
type InterfaceLeaveIdentExpr interface {
	LeaveIdentExpr(Node)
}
type InterfaceVisitUnaryExpr interface {
	VisitUnaryExpr(Node)
}
type InterfaceLeaveUnaryExpr interface {
	LeaveUnaryExpr(Node)
}
type InterfaceVisitBinaryExpr interface {
	VisitBinaryExpr(Node)
}
type InterfaceLeaveBinaryExpr interface {
	LeaveBinaryExpr(Node)
}
type InterfaceVisitNewExpr interface {
	VisitNewExpr(Node)
}
type InterfaceLeaveNewExpr interface {
	LeaveNewExpr(Node)
}
type InterfaceVisitTernaryExpr interface {
	VisitTernaryExpr(Node)
}
type InterfaceLeaveTernaryExpr interface {
	LeaveTernaryExpr(Node)
}
type InterfaceVisitParenExpr interface {
	VisitParenExpr(Node)
}
type InterfaceLeaveParenExpr interface {
	LeaveParenExpr(Node)
}
type InterfaceVisitNotExpr interface {
	VisitNotExpr(Node)
}
type InterfaceLeaveNotExpr interface {
	LeaveNotExpr(Node)
}
type InterfaceVisitStringExpr interface {
	VisitStringExpr(Node)
}
type InterfaceLeaveStringExpr interface {
	LeaveStringExpr(Node)
}
type InterfaceVisitAssignStmt interface {
	VisitAssignStmt(Node)
}
type InterfaceLeaveAssignStmt interface {
	LeaveAssignStmt(Node)
}
type InterfaceVisitReturnStmt interface {
	VisitReturnStmt(Node)
}
type InterfaceLeaveReturnStmt interface {
	LeaveReturnStmt(Node)
}
type InterfaceVisitBreakStmt interface {
	VisitBreakStmt(Node)
}
type InterfaceVisitContinueStmt interface {
	VisitContinueStmt(Node)
}
type InterfaceVisitRaiseStmt interface {
	VisitRaiseStmt(Node)
}
type InterfaceLeaveRaiseStmt interface {
	LeaveRaiseStmt(Node)
}
type InterfaceVisitExecuteStmt interface {
	VisitExecuteStmt(Node)
}
type InterfaceLeaveExecuteStmt interface {
	LeaveExecuteStmt(Node)
}
type InterfaceVisitCallStmt interface {
	VisitCallStmt(Node)
}
type InterfaceLeaveCallStmt interface {
	LeaveCallStmt(Node)
}
type InterfaceVisitIfStmt interface {
	VisitIfStmt(Node)
}
type InterfaceLeaveIfStmt interface {
	LeaveIfStmt(Node)
}
type InterfaceVisitElsIfStmt interface {
	VisitElsIfStmt(Node)
}
type InterfaceLeaveElsIfStmt interface {
	LeaveElsIfStmt(Node)
}
type InterfaceVisitElseStmt interface {
	VisitElseStmt(Node)
}
type InterfaceLeaveElseStmt interface {
	LeaveElseStmt(Node)
}
type InterfaceVisitWhileStmt interface {
	VisitWhileStmt(Node)
}
type InterfaceLeaveWhileStmt interface {
	LeaveWhileStmt(Node)
}
type InterfaceVisitForStmt interface {
	VisitForStmt(Node)
}
type InterfaceLeaveForStmt interface {
	LeaveForStmt(Node)
}
type InterfaceVisitForEachStmt interface {
	VisitForEachStmt(Node)
}
type InterfaceLeaveForEachStmt interface {
	LeaveForEachStmt(Node)
}
type InterfaceVisitTryStmt interface {
	VisitTryStmt(Node)
}
type InterfaceLeaveTryStmt interface {
	LeaveTryStmt(Node)
}
type InterfaceVisitExceptStmt interface {
	VisitExceptStmt(Node)
}
type InterfaceLeaveExceptStmt interface {
	LeaveExceptStmt(Node)
}
type InterfaceVisitGotoStmt interface {
	VisitGotoStmt(Node)
}
type InterfaceVisitLabelStmt interface {
	VisitLabelStmt(Node)
}
type InterfaceVisitPrepIfInst interface {
	VisitPrepIfInst(Node)
}
type InterfaceLeavePrepIfInst interface {
	LeavePrepIfInst(Node)
}
type InterfaceVisitPrepElsIfInst interface {
	VisitPrepElsIfInst(Node)
}
type InterfaceLeavePrepElsIfInst interface {
	LeavePrepElsIfInst(Node)
}
type InterfaceVisitPrepElseInst interface {
	VisitPrepElseInst(Node)
}
type InterfaceVisitPrepEndIfInst interface {
	VisitPrepEndIfInst(Node)
}
type InterfaceVisitPrepRegionInst interface {
	VisitPrepRegionInst(Node)
}
type InterfaceVisitPrepEndRegionInst interface {
	VisitPrepEndRegionInst(Node)
}
type InterfaceVisitPrepExpr interface {
	VisitPrepExpr(Node)
}
type InterfaceLeavePrepExpr interface {
	LeavePrepExpr(Node)
}
type InterfaceVisitPrepBinaryExpr interface {
	VisitPrepBinaryExpr(Node)
}
type InterfaceLeavePrepBinaryExpr interface {
	LeavePrepBinaryExpr(Node)
}
type InterfaceVisitPrepNotExpr interface {
	VisitPrepNotExpr(Node)
}
type InterfaceLeavePrepNotExpr interface {
	LeavePrepNotExpr(Node)
}
type InterfaceVisitPrepSymExpr interface {
	VisitPrepSymExpr(Node)
}
type InterfaceVisitPrepParenExpr interface {
	VisitPrepParenExpr(Node)
}
type InterfaceLeavePrepParenExpr interface {
	LeavePrepParenExpr(Node)
}
