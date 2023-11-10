package engine_test

import (
	"testing"

	cangjie "github.com/antonyho/go-cangjie/engine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestCase struct {
	name     string
	radicals string
	expected []rune
}

type CangjieV3TestSuite struct {
	suite.Suite
	cangjie.Engine
}

func (s *CangjieV3TestSuite) SetupSuite() {
	engine := cangjie.New(
		cangjie.WithCangjieV3(),
		cangjie.WithDatabase("testdata/cangjie.db"),
	)
	s.Engine = *engine
}

func (s *CangjieV3TestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *CangjieV3TestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'倉'}},
		{"multiple matches", "hqi", []rune{'我', '牫', '𥫻'}},
		{"cangjie v3 characters", "yhhqm", []rune{'產', '産'}},
		{"punctuation", "zxad", []rune{'。'}},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Encode(testCase.radicals)
			s.NoError(err)
			s.ElementsMatch(results, testCase.expected)
		})
	}
}

func TestCangjieV3TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV3TestSuite))
}

type CangjieV5TestSuite struct {
	suite.Suite
	cangjie.Engine
}

func (s *CangjieV5TestSuite) SetupSuite() {
	engine := cangjie.New(
		cangjie.WithCangjieV5(),
		cangjie.WithDatabase("testdata/cangjie.db"),
	)
	s.Engine = *engine
}

func (s *CangjieV5TestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *CangjieV5TestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'倉'}},
		{"multiple matches", "hqi", []rune{'我', '牫', '𥫻'}},
		{"cangjie v5 character", "yhhqm", []rune{'産'}},
		{"punctuation", "zxad", []rune{'。'}},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Encode(testCase.radicals)
			s.NoError(err)
			s.ElementsMatch(results, testCase.expected)
		})
	}
}

func TestCangjieV5TestSuite(t *testing.T) {
	suite.Run(t, new(CangjieV5TestSuite))
}

type WithSimplifiedTestSuite struct {
	suite.Suite
	cangjie.Engine
}

func (s *WithSimplifiedTestSuite) SetupSuite() {
	engine := cangjie.New(
		cangjie.WithSimplified(),
		cangjie.WithDatabase("testdata/cangjie.db"),
	)
	s.Engine = *engine
}

func (s *WithSimplifiedTestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *WithSimplifiedTestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'仓'}},
		{"multiple matches", "hqi", []rune{'我', '牫', '𥫻'}},
		{"cangjie v5 character", "yhhqm", []rune{'产'}},
		{"punctuation", "zxad", []rune{}},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Encode(testCase.radicals)
			s.NoError(err)
			s.ElementsMatch(results, testCase.expected)
		})
	}
}

func TestWithSimplifiedTestSuite(t *testing.T) {
	suite.Run(t, new(WithSimplifiedTestSuite))
}

type WithEasyTestSuite struct {
	suite.Suite
	cangjie.Engine
}

func (s *WithEasyTestSuite) SetupSuite() {
	engine := cangjie.New(
		cangjie.WithEasy(),
		cangjie.WithDatabase("testdata/cangjie.db"),
	)
	s.Engine = *engine
}

func (s *WithEasyTestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *WithEasyTestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"normal", "kx", []rune{'㿕', '癠', '𡚒', '𤟅'}},
		{"exceed radicals length", "kxj", []rune{'㿕', '癠', '𡚒', '𤟅'}},
		{"single radical", "s", []rune{'尸'}},
		{"multiple matches on single radical", "a", []rune{'日', '曰'}},
		{"punctuations", "zd", []rune{'。', '「', '﹏'}},
	}

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Encode(testCase.radicals)
			s.NoError(err)
			s.ElementsMatch(results, testCase.expected)
		})
	}
}

func TestWithEasyTestSuite(t *testing.T) {
	suite.Run(t, new(WithEasyTestSuite))
}

type WithPredictionTestSuite struct {
	suite.Suite
	cangjie.Engine
}

func (s *WithPredictionTestSuite) SetupSuite() {
	engine := cangjie.New(
		cangjie.WithPrediction(),
		cangjie.WithDatabase("testdata/cangjie.db"),
	)
	s.Engine = *engine
}

func (s *WithPredictionTestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *WithPredictionTestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'倉'}},
		{"multiple matches", "nsm", []rune{'刍', '張', '戼', '𩔘'}},
		{"punctuation", "zxad", []rune{'。'}},
	}

	s.True(s.Engine.Prediction)

	for _, testCase := range testCases {
		s.T().Run(testCase.name, func(t *testing.T) {
			results, err := s.Engine.Encode(testCase.radicals)
			s.NoError(err)
			s.ElementsMatch(results, testCase.expected)
		})
	}
}

func TestWithPredictionTestSuite(t *testing.T) {
	suite.Run(t, new(WithPredictionTestSuite))
}

func TestEngineInvalidDB(t *testing.T) {
	engine := cangjie.New(cangjie.WithDatabase("notexist.db"))
	_, err := engine.Encode("oiar")
	assert.Error(t, err)

	assert.NoError(t, engine.Close())
}

func TestEngineSetOption(t *testing.T) {
	engine := cangjie.New()

	assert.False(t, engine.OutputSimplified)

	engine.Set(cangjie.WithSimplified())
	assert.True(t, engine.OutputSimplified)

	assert.NoError(t, engine.Close())
}

func TestEngineEncodeForMultiRadicalSetsWord(t *testing.T) {
	wordWithMultipleRadicalSets := '曰'

	engine := cangjie.New(cangjie.WithDatabase("testdata/cangjie.db"))

	results, err := engine.Encode("a")
	assert.NoError(t, err)
	assert.Contains(t, results, wordWithMultipleRadicalSets)

	results, err = engine.Encode("xa")
	assert.NoError(t, err)
	assert.Contains(t, results, wordWithMultipleRadicalSets)

	assert.NoError(t, engine.Close())
}
