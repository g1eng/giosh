package builtins

import (
	. "gopkg.in/check.v1"
)

//test() with no argument exits with error
func (s *TestSuite) TestFailsWithErrorOnNoArgument(c *C) {
	c.Check(BuiltinTest(), NotNil)
}

//test() exits with unknown single option, like -c
func (s *TestSuite) TestFailsWithUnknownSingleOption(c *C) {
	c.Check(BuiltinTest("-c"), NotNil)
}

//test exits with unknown option with an argumnt
func (s *TestSuite) TestFailsWithUnknownOptionWithArgument(c *C) {
	c.Check(BuiltinTest("-c", "hoge"), NotNil)
}

//test exits for invalid options with three or more arguments
func (s *TestSuite) TestFailsWithNonOptionThreeArguments(c *C) {
	c.Check(BuiltinTest("a", "b", "c"), NotNil)
}
func (s *TestSuite) TestFailsWithNonOptionMoreThanThreeArguments(c *C) {
	c.Check(BuiltinTest("a", "b", "c", "1", "2"), NotNil)
}

//exits with error on single =
func (s *TestSuite) TestFailsWithSingleOperator(c *C) {
	c.Check(BuiltinTest("="), NotNil)
}

//exits with error on missing operand for =
func (s *TestSuite) TestFailsWithMissingOperandEquation(c *C) {
	c.Check(BuiltinTest("asa93", "="), NotNil)
}

//exits with missing operand for <
func (s *TestSuite) TestFailsWithMissingOperandLessThan(c *C) {
	c.Check(BuiltinTest("asa93", "<"), NotNil)
}

//exits with validated Equal
func (s *TestSuite) TestIsInEqual(c *C) {
	c.Check(BuiltinTest("asa", "=", "asa"), IsNil)
}

//exits with validated Mathematical Equal
func (s *TestSuite) TestIsInMathematicalEqual(c *C) {
	c.Check(BuiltinTest("1283829", "=", "1283829"), IsNil)
}

//exits for stderr with NotEqual // STDERR
func (s *TestSuite) TestIsNotEqual(c *C) {
	c.Check(BuiltinTest("asa", "=", "ase"), NotNil)
}

// < operand for non-mathematical
func (s *TestSuite) TestFailsForNonMathematicalGt(c *C) {
	c.Check(BuiltinTest("15", ">", "junior"), NotNil)
}

// <= operand
func (s *TestSuite) TestFailsForNonMathematicalGe(c *C) {
	c.Check(BuiltinTest("15", ">=", "junior"), NotNil)
}

// > operand for non-mathematical
func (s *TestSuite) TestFailsForNonMathematicalLt(c *C) {
	c.Check(BuiltinTest("65", "<", "senior"), NotNil)
}

// >= operand for non-mathematical
func (s *TestSuite) TestFailsForNonMathematicalLe(c *C) {
	c.Check(BuiltinTest("65", "<=", "senior"), NotNil)
}

// valid greater than
func (s *TestSuite) TestIsGt(c *C) {
	c.Check(BuiltinTest("65", ">", "15"), IsNil)
}

// valid less than
func (s *TestSuite) TestIsLt(c *C) {
	c.Check(BuiltinTest("314159265", ">", "65536"), IsNil)
}

// test cannot parse and evaluate float number
func (s *TestSuite) TestFailsOnFloatNumberLt(c *C) {
	c.Check(BuiltinTest("3.14159265", ">", "65536"), NotNil)
}

// invalid operator for the argument
func (s *TestSuite) TestEqualInvalid(c *C) {
	c.Check(BuiltinTest("="), NotNil)
}
func (s *TestSuite) TestGreaterThanInvalid(c *C) {
	c.Check(BuiltinTest(">"), NotNil)
}
func (s *TestSuite) TestGreaterEqualInvalid(c *C) {
	c.Check(BuiltinTest(">="), NotNil)
}
func (s *TestSuite) TestLessThanInvalid(c *C) {
	c.Check(BuiltinTest("<"), NotNil)
}
func (s *TestSuite) TestLessEqualInvalid(c *C) {
	c.Check(BuiltinTest("<="), NotNil)
}

// missing operand for operators =, >, and so on.
func (s *TestSuite) TestEqualMissingOperand(c *C) {
	c.Check(BuiltinTest("g", "="), NotNil)
}
func (s *TestSuite) TestGreaterThanMissingOperand(c *C) {
	c.Check(BuiltinTest("5", ">"), NotNil)
}
func (s *TestSuite) TestGreaterEqualMissingOperand(c *C) {
	c.Check(BuiltinTest("5", ">="), NotNil)
}
func (s *TestSuite) TestLessThanMissingOperand(c *C) {
	c.Check(BuiltinTest("5", "<"), NotNil)
}
func (s *TestSuite) TestLessEqualMissingOperand(c *C) {
	c.Check(BuiltinTest("5", "<="), NotNil)
}

// = operator

func (s *TestSuite) TestCheckIsEqualForStrings(c *C) {
	c.Check(CheckIsEqual("asa", "asa"), Equals, true)
}
func (s *TestSuite) TestCheckIsEqualForEquation(c *C) {
	c.Check(CheckIsEqual("123", "123"), Equals, true)
}
func (s *TestSuite) TestCheckIsEqualForNotEqual(c *C) {
	c.Check(CheckIsEqual("123", "1"), Equals, false)
}

func (s *TestSuite) TestEqual(c *C) {
	c.Check(BuiltinTest("asa", "=", "asa"), IsNil)
}

// != operator

func (s *TestSuite) TestCheckIsNotEqualForStrings(c *C) {
	c.Check(CheckIsNotEqual("asa", "asu"), Equals, true)
}
func (s *TestSuite) TestCheckIsNotEqualForMathematical(c *C) {
	c.Check(CheckIsNotEqual("132", "231"), Equals, true)
}
func (s *TestSuite) TestCheckIsNotEqualForEqual(c *C) {
	c.Check(CheckIsNotEqual("123", "123"), Equals, false)
}

func (s *TestSuite) TestNotEqual(c *C) {
	c.Check(BuiltinTest("asa", "!=", "ishi"), IsNil)
}

// > operator

func (s *TestSuite) TestIsGreaterThan17and5(c *C) {
	c.Check(CheckIsGreaterThan(17, 5), Equals, true)
}
func (s *TestSuite) TestIsGreaterThan5and17(c *C) {
	c.Check(CheckIsGreaterThan(5, 17), Equals, false)
}

func (s *TestSuite) TestGreaterThan(c *C) {
	c.Check(BuiltinTest("17", ">", "5"), IsNil)
}

// >= operator

func (s *TestSuite) TestIsGreaterEqual17and17(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), Equals, true)
}
func (s *TestSuite) TestIsGreaterEqual17and5(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), Equals, true)
}
func (s *TestSuite) TestIsGreaterEqual5and17(c *C) {
	c.Check(CheckIsGreaterEqual(5, 17), Equals, false)
}

func (s *TestSuite) TestGreaterEqual(c *C) {
	c.Check(BuiltinTest("5", ">=", "5"), IsNil)
}

// < operator

func (s *TestSuite) TestIsLessThan17and5(c *C) {
	c.Check(CheckIsLessThan(17, 5), Equals, false)
}
func (s *TestSuite) TestIsLessThan5and17(c *C) {
	c.Check(CheckIsLessThan(5, 17), Equals, true)
}
func (s *TestSuite) TestLessThan(c *C) {
	c.Check(BuiltinTest("5", "<", "17"), IsNil)
}

// <= operator
func (s *TestSuite) TestIsLessEqual17and17(c *C) {
	c.Check(CheckIsLessEqual(17, 17), Equals, true)
}
func (s *TestSuite) TestIsLessEqual17and5(c *C) {
	c.Check(CheckIsLessEqual(17, 5), Equals, false)
}
func (s *TestSuite) TestIsLessEqual5and17(c *C) {
	c.Check(CheckIsLessEqual(5, 17), Equals, true)
}
func (s *TestSuite) TestLessEqual(c *C) {
	c.Check(BuiltinTest("5", "<=", "5"), IsNil)
}

// -f option

func (s *TestSuite) TestCheckFileReadExistFile(c *C) {
	c.Check(CheckFile("/bin/ls"), IsNil)
}
func (s *TestSuite) TestCheckFileFailsToReadInvalidPath(c *C) {
	c.Check(CheckFile("/bin/ls/not/found"), NotNil)
}
func (s *TestSuite) TestFileForValidPath(c *C) {
	c.Check(BuiltinTest("-f", "/bin/ls"), IsNil)
}
func (s *TestSuite) TestFileForInvalidPath(c *C) {
	c.Check(BuiltinTest("-f", "/bin/ls/not/found"), NotNil)
}

// -d option

func (s *TestSuite) TestCheckDirForExistDir(c *C) {
	c.Check(CheckDirectory("/bin"), IsNil)
}
func (s *TestSuite) TestCheckDirForNotExistDir(c *C) {
	c.Check(CheckDirectory("/bin/ls/not/found"), NotNil)
}
func (s *TestSuite) TestDirForValidPath(c *C) {
	c.Check(BuiltinTest("-d", "/bin"), IsNil)
}
func (s *TestSuite) TestDirForInvalidPath(c *C) {
	c.Check(BuiltinTest("-d", "/bin/ls/not/found"), NotNil)
}

// -n option
func (s *TestSuite) TestCheckNonZeroForOne(c *C) {
	c.Check(CheckNonZero("1"), Equals, true)
}
func (s *TestSuite) TestCheckNonZeroForZero(c *C) {
	c.Check(CheckNonZero(""), Equals, false)
}

func (s *TestSuite) TestNonZero(c *C) {
	c.Check(BuiltinTest("-n", "1"), IsNil)
}
func (s *TestSuite) TestNonZeroFailure(c *C) {
	c.Check(BuiltinTest("-n", ""), NotNil)
}

// -z option

func (s *TestSuite) TestCheckZeroForZero(c *C) {
	c.Check(CheckZero(""), Equals, true)
}

func (s *TestSuite) TestCheckZeroForNonZero(c *C) {
	c.Check(CheckZero("1"), Equals, false)
}

func (s *TestSuite) TestZero(c *C) {
	c.Check(BuiltinTest("-z", ""), IsNil)
}
func (s *TestSuite) TestZeroFailure(c *C) {
	c.Check(BuiltinTest("-z", "ok"), NotNil)
}

// -h option

func (s *TestSuite) TestCheckSlnForValidSln(c *C) {
	c.Check(CheckSymbolicLink("/var/run"), IsNil)
}
func (s *TestSuite) TestCheckSlnForInvalidSln(c *C) {
	c.Check(CheckSymbolicLink("/bin/ls"), NotNil)
}
func (s *TestSuite) TestSln(c *C) {
	c.Check(BuiltinTest("-h", "/var/run"), IsNil)
}
