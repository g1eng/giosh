package builtins

import (
	. "gopkg.in/check.v1"
	"testing"
)

type TestTestSuite struct{}

func init() {
	Suite(&TestTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

//test() with no argument exits with error
func (s *TestTestSuite) TestFailsWithErrorOnNoArgument(c *C) {
	c.Check(BuiltinTest(), NotNil)
}

//test() exits with unknown single option, like -c
func (s *TestTestSuite) TestFailsWithUnknownSingleOption(c *C) {
	c.Check(BuiltinTest("-c"), NotNil)
}

//test exits with unknown option with an argumnt
func (s *TestTestSuite) TestFailsWithUnknownOptionWithArgument(c *C) {
	c.Check(BuiltinTest("-c", "hoge"), NotNil)
}

//passed for valid option with an argument
func (s *TestTestSuite) TestPassedForValidOptionWithAnArgument(c *C) {
	c.Check(BuiltinTest("-d", "thisDir"), IsNil)
}

//exits with error on single =
func (s *TestTestSuite) TestFailsWithSingleOperator(c *C) {
	c.Check(BuiltinTest("="), NotNil)
}

//exits with error on missing operand for =
func (s *TestTestSuite) TestFailsWithMissingOperandEquation(c *C) {
	c.Check(BuiltinTest("asa93", "="), NotNil)
}

//exits with missing operand for <
func (s *TestTestSuite) TestFailsWithMissingOperandLessThan(c *C) {
	c.Check(BuiltinTest("asa93", "<"), NotNil)
}

//exits with validated Equal
func (s *TestTestSuite) TestIsInEqual(c *C) {
	c.Check(BuiltinTest("asa", "=", "asa"), IsNil)
}

//exits with validated Mathematical Equal
func (s *TestTestSuite) TestIsInMathematicalEqual(c *C) {
	c.Check(BuiltinTest("1283829", "=", "1283829"), IsNil)
}

//exits for stderr with NotEqual // STDERR
func (s *TestTestSuite) TestIsNotEqual(c *C) {
	c.Check(BuiltinTest("asa", "=", "ase"), NotNil)
}

// < operand is for mathematical
func (s *TestTestSuite) TestFailsForNonMathematicalGt(c *C) {
	c.Check(BuiltinTest("15", ">", "junior"), NotNil)
}

// > operand is for mathematical
func (s *TestTestSuite) TestFailsForNonMathematicalLt(c *C) {
	c.Check(BuiltinTest("65", "<", "senior"), NotNil)
}

// valid greater than
func (s *TestTestSuite) TestIsGt(c *C) {
	c.Check(BuiltinTest("65", ">", "15"), IsNil)
}

// valid less than
func (s *TestTestSuite) TestIsLt(c *C) {
	c.Check(BuiltinTest("314159265", ">", "65536"), IsNil)
}

// test cannot parse and evaluate float number
func (s *TestTestSuite) TestFailsOnFloatNumberLt(c *C) {
	c.Check(BuiltinTest("3.14159265", ">", "65536"), NotNil)
}

/// unit tests

func (s *TestTestSuite) TestCheckIsEqualForStrings(c *C) {
	c.Check(CheckIsEqual("asa", "asa"), IsNil)
}
func (s *TestTestSuite) TestCheckIsEqualForEquation(c *C) {
	c.Check(CheckIsEqual("123", "123"), IsNil)
}
func (s *TestTestSuite) TestCheckIsEqualForNotEqual(c *C) {
	c.Check(CheckIsEqual("123", "1"), NotNil)
}
func (s *TestTestSuite) TestCheckIsNotEqualForStrings(c *C) {
	c.Check(CheckIsNotEqual("asa", "asu"), IsNil)
}
func (s *TestTestSuite) TestCheckIsNotEqualForMathematical(c *C) {
	c.Check(CheckIsNotEqual("132", "231"), IsNil)
}
func (s *TestTestSuite) TestCheckIsNotEqualForEqual(c *C) {
	c.Check(CheckIsNotEqual("123", "123"), NotNil)
}
func (s *TestTestSuite) TestIsGreaterThan17and5(c *C) {
	c.Check(CheckIsGreaterThan(17, 5), IsNil)
}
func (s *TestTestSuite) TestIsGreaterThan5and17(c *C) {
	c.Check(CheckIsGreaterThan(5, 17), NotNil)
}
func (s *TestTestSuite) TestIsGreaterEqual17and17(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), IsNil)
}
func (s *TestTestSuite) TestIsGreaterEqual17and5(c *C) {
	c.Check(CheckIsGreaterEqual(17, 5), IsNil)
}
func (s *TestTestSuite) TestIsGreaterEqual5and17(c *C) {
	c.Check(CheckIsGreaterEqual(5, 17), NotNil)
}
func (s *TestTestSuite) TestIsLessThan17and5(c *C) {
	c.Check(CheckIsLessThan(17, 5), NotNil)
}
func (s *TestTestSuite) TestIsLessThan5and17(c *C) {
	c.Check(CheckIsLessThan(5, 17), IsNil)
}
func (s *TestTestSuite) TestIsLessEqual17and17(c *C) {
	c.Check(CheckIsLessEqual(17, 17), IsNil)
}
func (s *TestTestSuite) TestIsLessEqual17and5(c *C) {
	c.Check(CheckIsLessEqual(17, 5), NotNil)
}
func (s *TestTestSuite) TestIsLessEqual5and17(c *C) {
	c.Check(CheckIsLessEqual(5, 17), IsNil)
}
