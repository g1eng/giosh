package builtins

import (
	"errors"
	"strconv"
)

//BuiltinTest is a `test` builtin implementation for giosh.
func BuiltinTest(expr ...string) error {
	if len(expr) < 1 {
		return errors.New("operator missing")
	} else if len(expr) == 1 {
		switch expr[0] {
		case "-t":
			return HasActiveTTY()
		default:
			return errors.New("unknown option" + expr[0])
		}
	} else if len(expr) == 2 {
		target := expr[1]
		switch expr[0] {
		case "-f":
			return CheckFile(target)
		case "-d":
			return CheckDirectory(target)
		case "-h":
			return CheckSymbolicLink(target)
		case "-s":
			return CheckSocket(target)
		case "-z":
			return CheckZero(target)
		case "-n":
			return CheckNonZero(target)
		case "=":
		case "<":
		case "<=":
		case ">":
		case ">=":
			return errors.New("operand missing: " + expr[1])
		default:
			return errors.New("unknown option " + expr[0])
		}
	} else if len(expr) == 2 {
		switch expr[1] {
		case "=":
		case "<":
		case "<=":
		case ">":
		case ">=":
			return errors.New("right hand operand missing: " + expr[1])
		default:
			return errors.New("unknown syntax " + expr[0] + " " + expr[1])
		}
	} else if len(expr) == 3 {
		op1, op2 := expr[0], expr[2]
		operator := expr[1]
		switch operator {
		case "=":
			return CheckIsEqual(op1, op2)
		case "!=":
			return CheckIsNotEqual(op1, op2)
		case "<":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				return CheckIsLessThan(op1, op2)
			}
		case "<=":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				return CheckIsLessEqual(op1, op2)
			}
		case ">":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				return CheckIsGreaterThan(op1, op2)
			}
		case ">=":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				return CheckIsGreaterEqual(op1, op2)
			}
		default:
			return errors.New("unknown operator " + expr[1])
		}

	} else {
		return errors.New("too long arguments specified for test")
	}
	return nil
}

//stub functions

func HasActiveTTY() error              { return nil }
func CheckFile(_ string) error         { return nil }
func CheckDirectory(_ string) error    { return nil }
func CheckSocket(_ string) error       { return nil }
func CheckSymbolicLink(_ string) error { return nil }
func CheckNonZero(_ string) error      { return nil }
func CheckZero(_ string) error         { return nil }

func evaluateAsciiAsInt(op1 string, op2 string) (opp1 int, opp2 int, err error) {
	if opp1, err = strconv.Atoi(op1); err != nil {
		return 0, 0, err
	} else if opp2, err = strconv.Atoi(op2); err != nil {
		return 0, 0, err
	} else {
		return opp1, opp2, nil
	}
}

func CheckIsEqual(op1, op2 string) error {
	if op1 == op2 {
		return nil
	} else {
		return errors.New("")
	}
}
func CheckIsNotEqual(op1, op2 string) error {
	if op1 != op2 {
		return nil
	} else {
		return errors.New("")
	}
}

func CheckIsLessThan(op1, op2 int) error {
	if op1 < op2 {
		return nil
	} else {
		return errors.New("")
	}
}

func CheckIsLessEqual(op1, op2 int) error {
	if op1 <= op2 {
		return nil
	} else {
		return errors.New("")
	}
}

func CheckIsGreaterThan(op1, op2 int) error {
	if op1 > op2 {
		return nil
	} else {
		return errors.New("")
	}
}

func CheckIsGreaterEqual(op1, op2 int) error {
	if op1 > op2 {
		return nil
	} else {
		return errors.New("")
	}
}
