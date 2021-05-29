package builtins

import (
	"fmt"
	"os"
	"strconv"
)

//BuiltinTest is a `test` builtin implementation for giosh.
func BuiltinTest(expr ...string) error {
	if len(expr) < 1 {
		return fmt.Errorf("operator missing")
	} else if len(expr) == 1 {
		switch expr[0] {
		default:
			return fmt.Errorf("unknown option, %s", expr[0])
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
		case "-z":
			if !CheckZero(target) {
				return fmt.Errorf("")
			}
		case "-n":
			if !CheckNonZero(target) {
				return fmt.Errorf("")
			}
		default:
			return fmt.Errorf("unknown option: %s", expr[0])
		}
	} else if len(expr) == 3 {
		op1, op2 := expr[0], expr[2]
		operator := expr[1]
		var res bool
		switch operator {
		case "=":
			res = CheckIsEqual(op1, op2)
		case "!=":
			res = CheckIsNotEqual(op1, op2)
		case "<":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				res = CheckIsLessThan(op1, op2)
			}
		case "<=":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				res = CheckIsLessEqual(op1, op2)
			}
		case ">":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				res = CheckIsGreaterThan(op1, op2)
			}
		case ">=":
			if op1, op2, err := evaluateAsciiAsInt(op1, op2); err != nil {
				return err
			} else {
				res = CheckIsGreaterEqual(op1, op2)
			}
		default:
			return fmt.Errorf("unknown operator %s", expr[1])
		}
		if res != true {
			return fmt.Errorf("")
		}
	} else {
		return fmt.Errorf("too long arguments specified for test")
	}
	return nil
}

//CheckFile checks a file existence with a filePath given
func CheckFile(filePath string) error {
	if _, err := os.Open(filePath); err != nil {
		return err
	} else {
		return nil
	}
}

//CheckDirectory checks a directory existence with a dirPath given
func CheckDirectory(dirPath string) error {
	if _, err := os.ReadDir(dirPath); err != nil {
		return err
	} else {
		return nil
	}
}

//CheckNonZero checks a string is not a blank string slice
func CheckNonZero(s string) bool {
	if s == "" {
		return false
	} else {
		return true
	}
}

//CheckZero checks a string IS a blank string slice
func CheckZero(s string) bool {
	if s != "" {
		return false
	} else {
		return true
	}
}

func CheckSymbolicLink(sln string) error {
	if _, err := os.Readlink(sln); err != nil {
		return err
	} else {
		return nil
	}
}

func evaluateAsciiAsInt(op1 string, op2 string) (opp1 int, opp2 int, err error) {
	if opp1, err = strconv.Atoi(op1); err != nil {
		return 0, 0, err
	} else if opp2, err = strconv.Atoi(op2); err != nil {
		return 0, 0, err
	} else {
		return opp1, opp2, nil
	}
}

func CheckIsEqual(op1, op2 string) bool {
	if op1 == op2 {
		return true
	} else {
		return false
	}
}
func CheckIsNotEqual(op1, op2 string) bool {
	if op1 != op2 {
		return true
	} else {
		return false
	}
}

func CheckIsLessThan(op1, op2 int) bool {
	if op1 < op2 {
		return true
	} else {
		return false
	}
}

func CheckIsLessEqual(op1, op2 int) bool {
	if op1 <= op2 {
		return true
	} else {
		return false
	}
}

func CheckIsGreaterThan(op1, op2 int) bool {
	if op1 > op2 {
		return true
	} else {
		return false
	}
}

func CheckIsGreaterEqual(op1, op2 int) bool {
	return op1 >= op2
}
