package engine

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CangjieV3TestSuite struct {
	suite.Suite
	Engine
}

func (s *CangjieV3TestSuite) SetupTest() {}

func (s *CangjieV3TestSuite) TestEngine_Get() {}

func TestCangjieV3TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV3TestSuite))
}

type CangjieV5TestSuite struct {
	suite.Suite
	Engine
}

func (s *CangjieV5TestSuite) SetupTest() {}

func (s *CangjieV5TestSuite) TestEngine_Get() {}

func TestCangjieV5TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV5TestSuite))
}
