package engine_test

import (
	"path"
	"testing"

	congkit "github.com/antonyho/go-congkit/engine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	TestDBPath = "testdata/congkit.db"
)

type TestCase struct {
	name     string
	radicals string
	expected []rune
}

type CongkitV3TestSuite struct {
	suite.Suite
	congkit.Engine
}

func (s *CongkitV3TestSuite) SetupSuite() {
	engine := congkit.New(
		congkit.WithCongkitV3(),
		congkit.WithDatabase(TestDBPath),
	)
	s.Engine = *engine
}

func (s *CongkitV3TestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *CongkitV3TestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'倉'}},
		{"multiple matches", "hqi", []rune{'我', '牫', '𥫻'}},
		{"congkit v3 characters", "yhhqm", []rune{'產', '産'}},
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

func TestCongkitV3TestSuite(t *testing.T) {
	suite.Run(t, new(CongkitV3TestSuite))
}

type CongkitV5TestSuite struct {
	suite.Suite
	congkit.Engine
}

func (s *CongkitV5TestSuite) SetupSuite() {
	engine := congkit.New(
		congkit.WithCongkitV5(),
		congkit.WithDatabase(TestDBPath),
	)
	s.Engine = *engine
}

func (s *CongkitV5TestSuite) TearDownSuite() {
	s.Engine.Close()
}

func (s *CongkitV5TestSuite) TestEngineEncode() {
	var testCases = []TestCase{
		{"no match", "abcd", []rune{}},
		{"single match", "oiar", []rune{'倉'}},
		{"multiple matches", "hqi", []rune{'我', '牫', '𥫻'}},
		{"congkit v5 character", "yhhqm", []rune{'産'}},
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

func TestCongkitV5TestSuite(t *testing.T) {
	suite.Run(t, new(CongkitV5TestSuite))
}

type WithSimplifiedTestSuite struct {
	suite.Suite
	congkit.Engine
}

func (s *WithSimplifiedTestSuite) SetupSuite() {
	engine := congkit.New(
		congkit.WithSimplified(),
		congkit.WithDatabase(TestDBPath),
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
		{"congkit v5 character", "yhhqm", []rune{'产'}},
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
	congkit.Engine
}

func (s *WithEasyTestSuite) SetupSuite() {
	engine := congkit.New(
		congkit.WithEasy(),
		congkit.WithDatabase(TestDBPath),
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
	congkit.Engine
}

func (s *WithPredictionTestSuite) SetupSuite() {
	engine := congkit.New(
		congkit.WithPrediction(),
		congkit.WithDatabase(TestDBPath),
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
	tmpDir := t.TempDir()
	notExistDBName := "notexist.db"
	notExistDbPath := path.Join(tmpDir, notExistDBName)
	engine := congkit.New(congkit.WithDatabase(notExistDbPath))
	_, err := engine.Encode("oiar")
	assert.Error(t, err)

	assert.NoError(t, engine.Close())
}

func TestEngineSetOption(t *testing.T) {
	engine := congkit.New()

	assert.False(t, engine.OutputSimplified)

	engine.Set(congkit.WithSimplified())
	assert.True(t, engine.OutputSimplified)

	assert.NoError(t, engine.Close())
}

func TestEngineEncodeForMultiRadicalSetsWord(t *testing.T) {
	wordWithMultipleRadicalSets := '曰'

	engine := congkit.New(congkit.WithDatabase(TestDBPath))

	results, err := engine.Encode("a")
	assert.NoError(t, err)
	assert.Contains(t, results, wordWithMultipleRadicalSets)

	results, err = engine.Encode("xa")
	assert.NoError(t, err)
	assert.Contains(t, results, wordWithMultipleRadicalSets)

	assert.NoError(t, engine.Close())
}
