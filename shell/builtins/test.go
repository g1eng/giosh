package builtins

import (
	"errors"
	"strconv"
)

func test(expr ...string) (err error) {
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
		var op1, op2 int
		if op1, err = strconv.Atoi(expr[0]); err != nil {
			return err
		} else if op2, err = strconv.Atoi(expr[2]); err != nil {
			return err
		} else {
			switch expr[1] {
			case "=":
				return func() error {
					if op1 == op2 {
						return nil
					} else {
						return errors.New("")
					}
				}()
			case "<":
				return func() error {
					if op1 < op2 {
						return nil
					} else {
						return errors.New("")
					}
				}()
			case "<=":
				return func() error {
					if op1 <= op2 {
						return nil
					} else {
						return errors.New("")
					}
				}()
			case ">":
				return func() error {
					if op1 > op2 {
						return nil
					} else {
						return errors.New("")
					}
				}()
			case ">=":
				return func() error {
					if op1 >= op2 {
						return nil
					} else {
						return errors.New("")
					}
				}()
			default:
				return errors.New("unknown operator " + expr[1])
			}
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
