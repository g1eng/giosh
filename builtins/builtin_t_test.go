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

//passed for valid option with an argument
func (s *TestSuite) TestPassedForValidOptionWithAnArgument(c *C) {
	c.Check(BuiltinTest("-d", "thisDir"), IsNil)
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

// < operand is for mathematical
func (s *TestSuite) TestFailsForNonMathematicalGt(c *C) {
	c.Check(BuiltinTest("15", ">", "junior"), NotNil)
}

// > operand is for mathematical
func (s *TestSuite) TestFailsForNonMathematicalLt(c *C) {
	c.Check(BuiltinTest("65", "<", "senior"), NotNil)
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

/// unit tests

func (s *TestSuite) TestCheckIsEqualForStrings(c *C) {
	c.Check(CheckIsEqual("asa", "asa"), IsNil)
}
func (s *TestSuite) TestCheckIsEqualForEquation(c *C) {
	c.Check(CheckIsEqual("123", "123"), IsNil)
}
func (s *TestSuite) TestCheckIsEqualForNotEqual(c *C) {
	c.Check(CheckIsEqual("123", "1"), NotNil)
}
func (s *TestSuite) TestCheckIsNotEqualForStrings(c *C) {
	c.Check(CheckIsNotEqual("asa", "asu"), IsNil)
}
func (s *TestSuite) TestCheckIsNotEqualForMathematical(c *C) {
	c.Check(CheckIsNotEqual("132", "231"), IsNil)
}
func (s *TestSuite) TestCheckIsNotEqualForEqual(c *C) {
	c.Check(CheckIsNotEqual("123", "123"), NotNil)
}
func (s *TestSuite) TestIsGreaterThan17and5(c *C) {
	c.Check(CheckIsGreaterThan(17, 5), IsNil)
}
func (s *TestSuite) TestIsGreaterThan5and17(c *C) {
	c.Check(CheckIsGreaterThan(5, 17), NotNil)
}
func (s *TestSuite) TestIsGreaterEqual17and17(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), IsNil)
}
func (s *TestSuite) TestIsGreaterEqual17and5(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), IsNil)
}
func (s *TestSuite) TestIsGreaterEqual5and17(c *C) {
	c.Check(CheckIsGreaterEqual(5, 17), NotNil)
}
func (s *TestSuite) TestIsLessThan17and5(c *C) {
	c.Check(CheckIsLessThan(17, 5), NotNil)
}
func (s *TestSuite) TestIsLessThan5and17(c *C) {
	c.Check(CheckIsLessThan(5, 17), IsNil)
}
func (s *TestSuite) TestIsLessEqual17and17(c *C) {
	c.Check(CheckIsLessEqual(17, 17), IsNil)
}
func (s *TestSuite) TestIsLessEqual17and5(c *C) {
	c.Check(CheckIsLessEqual(17, 5), NotNil)
}
func (s *TestSuite) TestIsLessEqual5and17(c *C) {
	c.Check(CheckIsLessEqual(5, 17), IsNil)
}
