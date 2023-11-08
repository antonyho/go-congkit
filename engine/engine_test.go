package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CangjieV3TestSuite struct {
	suite.Suite
	Engine
}

func (s *CangjieV3TestSuite) SetupTest() {
	engine, err := New(WithCangjieV3())
	if err != nil {
		s.FailNow("failed initialising the engine", err)
	}
	s.Engine = *engine
}

func (s *CangjieV3TestSuite) TestEngine_Compose() {
	var testCases = []struct {
		name      string
		radicals  string
		expected  []rune
		expectErr bool
	}{
		{"no match", "", []rune{}, false},
		{"single match", "", []rune{}, false},
		{"multiple matches", "", []rune{}, false},
		{"punctuation", "", []rune{}, false},
		{"db error", "", []rune{}, true},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Compose(testCase.radicals)
			if testCase.expectErr {
				assert.Error(s.T(), err)
			}
			assert.EqualValues(s.T(), testCase.expected, results)
		})
	}
}

func TestCangjieV3TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV3TestSuite))
}

type CangjieV5TestSuite struct {
	suite.Suite
	Engine
}

func (s *CangjieV5TestSuite) SetupTest() {
	engine, err := New(WithCangjieV5())
	if err != nil {
		s.FailNow("failed initialising the engine", err)
	}
	s.Engine = *engine
}

func (s *CangjieV5TestSuite) TestEngine_Compose() {}

func TestCangjieV5TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV5TestSuite))
}

type WithSimplifiedTestSuite struct {
	suite.Suite
	Engine
}

func (s *WithSimplifiedTestSuite) SetupTest() {
	engine, err := New(WithSimplified())
	if err != nil {
		s.FailNow("failed initialising the engine", err)
	}
	s.Engine = *engine
}

func (s *WithSimplifiedTestSuite) TestEngine_Compose() {}

func TestWithSimplifiedTestSuite(t *testing.T) {
	suite.Run(t, new(WithSimplifiedTestSuite))
}

type WithEasyTestSuite struct {
	suite.Suite
	Engine
}

func (s *WithEasyTestSuite) SetupTest() {
	engine, err := New(WithEasy())
	if err != nil {
		s.FailNow("failed initialising the engine", err)
	}
	s.Engine = *engine
}

func (s *WithEasyTestSuite) TestEngine_Compose() {}

func TestWithEasyTestSuite(t *testing.T) {
	suite.Run(t, new(WithEasyTestSuite))
}

type WithPredictionTestSuite struct {
	suite.Suite
	Engine
}

func (s *WithPredictionTestSuite) SetupTest() {
	engine, err := New(WithPrediction())
	if err != nil {
		s.FailNow("failed initialising the engine", err)
	}
	s.Engine = *engine
}

func (s *WithPredictionTestSuite) TestEngine_Compose() {}

func TestWithPredictionTestSuite(t *testing.T) {
	suite.Run(t, new(WithPredictionTestSuite))
}
