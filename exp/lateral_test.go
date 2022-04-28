package exp_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yz89122/goqu/v10/exp"
)

type lateralExpressionSuite struct {
	suite.Suite
}

func TestLateralExpressionSuite(t *testing.T) {
	suite.Run(t, &lateralExpressionSuite{})
}

func (les *lateralExpressionSuite) TestClone() {
	le := exp.NewLateralExpression(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{}))
	les.Equal(exp.NewLateralExpression(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{})), le.Clone())
}

func (les *lateralExpressionSuite) TestExpression() {
	le := exp.NewLateralExpression(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{}))
	les.Equal(le, le.Expression())
}

func (les *lateralExpressionSuite) TestLateral() {
	le := exp.NewLateralExpression(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{}))
	les.Equal(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{}), le.Table())
}

func (les *lateralExpressionSuite) TestAs() {
	le := exp.NewLateralExpression(newTestAppendableExpression(`SELECT * FROM "test"`, []interface{}{}))
	les.Equal(exp.NewAliasExpression(le, "foo"), le.As("foo"))
}
