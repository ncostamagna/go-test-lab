package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite_Unit(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

func (suite *ExampleSuite) TestExample() {
	suite.Equal(1, 1)
}

func (suite *ExampleSuite) TestExample2() {
	suite.Equal(2, 2)
}

// SetupSuite is called once before the tests in the suite are run
func (suite *ExampleSuite) SetupSuite() {
	suite.T().Log("SetupSuite") // log in the test output
}

// TearDownSuite is called once after the tests in the suite are run
// close database, etc.
func (suite *ExampleSuite) TearDownSuite() {
	suite.T().Log("TearDownSuite") // log in the test output
}

// SetupTest is called before each test
func (suite *ExampleSuite) SetupTest() {
	suite.T().Log("SetupTest") // log in the test output
}

// TearDownTest is called after each test
func (suite *ExampleSuite) TearDownTest() {
	suite.T().Log("TearDownTest") // log in the test output
}

func (suite *ExampleSuite) BeforeTest(suiteName, testName string) {
	suite.T().Logf("BeforeTest: %s, %s", suiteName, testName) // log in the test output
}

func (suite *ExampleSuite) AfterTest(suiteName, testName string) {
	suite.T().Logf("AfterTest: %s, %s", suiteName, testName) // log in the test output
}






type ExampleSuiteIntegration struct {
	suite.Suite
}

func TestExampleSuite_Integration(t *testing.T) {
	suite.Run(t, new(ExampleSuiteIntegration))
}

func (suite *ExampleSuiteIntegration) TestExample() {
	suite.Equal(1, 1)
}

func (suite *ExampleSuiteIntegration) TestExample2() {
	suite.Equal(2, 2)
}

func (suite *ExampleSuiteIntegration) SetupSuite() {
	suite.T().Log("SetupSuite_Integration")
}

func (suite *ExampleSuiteIntegration) TearDownSuite() {
	suite.T().Log("TearDownSuite_Integration")
}

func (suite *ExampleSuiteIntegration) SetupTest() {
	suite.T().Log("SetupTest_Integration")
}

func (suite *ExampleSuiteIntegration) TearDownTest() {
	suite.T().Log("TearDownTest_Integration")
}

func (suite *ExampleSuiteIntegration) BeforeTest(suiteName, testName string) {
	suite.T().Logf("BeforeTest_Integration: %s, %s", suiteName, testName)
}

func (suite *ExampleSuiteIntegration) AfterTest(suiteName, testName string) {
	suite.T().Logf("AfterTest_Integration: %s, %s", suiteName, testName)
}