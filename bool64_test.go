package bool64_test

import (
	"testing"

	"git.bangmarcel.art/marcel/bool64"
	"github.com/stretchr/testify/suite"
)

const (
	directory bool64.Bool64 = 1 << iota
	userRead
	userWrite
	userExecute
	groupRead
	groupWrite
	groupExecute
	otherRead
	otherWrite
	otherExecute
)

type Bool64TestSuite struct {
	suite.Suite
}

func TestBool64TestSuite(t *testing.T) {
	suite.Run(t, new(Bool64TestSuite))
}

func (s *Bool64TestSuite) TestSatisfies() {
	s.Run("should return true", func() {
		p := userRead | userWrite | userExecute | groupRead | groupWrite | otherRead | otherExecute
		ok := p.Satisfies(userRead)
		s.True(ok)
	})

	s.Run("should return true with multi flags", func() {
		p := userRead | userWrite | userExecute | groupRead | groupWrite | otherRead | otherExecute
		ok := p.Satisfies(userRead | groupRead)
		s.True(ok)
	})

	s.Run("should return false", func() {
		p := userRead | userWrite | userExecute | groupRead | groupWrite | otherRead | otherWrite
		ok := p.Satisfies(otherExecute)
		s.False(ok)
	})

	s.Run("should return false with multi flags", func() {
		p := userRead | userWrite | userExecute | groupRead | groupWrite | otherRead | otherWrite
		ok := p.Satisfies(otherExecute | userExecute)
		s.False(ok)
	})

	s.Run("should return false without correct flags", func() {
		p := userRead | userWrite | userExecute | groupRead | groupWrite | otherRead | otherWrite
		ok := p.Satisfies(otherExecute | directory)
		s.False(ok)
	})
}

func (s *Bool64TestSuite) TestHas() {
	s.Run("should return true if single flag is present", func() {
		p := userRead | userWrite
		s.True(p.Has(userRead))
		s.True(p.Has(userWrite))
	})

	s.Run("should return true if at least one of multi flags is present", func() {
		p := userRead | userWrite
		s.True(p.Has(userRead | userExecute))
		s.True(p.Has(userWrite | groupRead))
	})

	s.Run("should return false if flag is not present", func() {
		p := userRead | userWrite
		s.False(p.Has(userExecute))
		s.False(p.Has(groupRead))
	})

	s.Run("should return false if none of multi flags are present", func() {
		p := userRead | userWrite
		s.False(p.Has(userExecute | groupRead | otherRead))
	})
}

func (s *Bool64TestSuite) TestAdd() {
	s.Run("should add single flag", func() {
		var p bool64.Bool64
		p.Add(userRead)
		s.True(p.Has(userRead))
		s.False(p.Has(userWrite))
	})

	s.Run("should add multiple flags", func() {
		var p bool64.Bool64
		p.Add(userRead | userWrite)
		s.True(p.Has(userRead))
		s.True(p.Has(userWrite))
		s.True(p.Satisfies(userRead | userWrite))
	})

	s.Run("should be idempotent when adding existing flags", func() {
		p := userRead
		p.Add(userRead)
		s.Equal(userRead, p)
	})
}

func (s *Bool64TestSuite) TestRemove() {
	s.Run("should remove single flag", func() {
		p := userRead | userWrite
		p.Remove(userWrite)
		s.True(p.Has(userRead))
		s.False(p.Has(userWrite))
	})

	s.Run("should remove multiple flags", func() {
		p := userRead | userWrite | userExecute
		p.Remove(userWrite | userExecute)
		s.True(p.Has(userRead))
		s.False(p.Has(userWrite))
		s.False(p.Has(userExecute))
	})

	s.Run("should do nothing if flag to remove is not present", func() {
		p := userRead
		p.Remove(userWrite)
		s.Equal(userRead, p)
	})
}

func (s *Bool64TestSuite) TestToggle() {
	s.Run("should toggle single flag from off to on", func() {
		var p bool64.Bool64
		p.Toggle(userRead)
		s.True(p.Has(userRead))
	})

	s.Run("should toggle single flag from on to off", func() {
		p := userRead
		p.Toggle(userRead)
		s.False(p.Has(userRead))
		s.Equal(bool64.Bool64(0), p)
	})

	s.Run("should toggle multiple flags mixed state", func() {
		p := userRead
		p.Toggle(userRead | userWrite) // userRead: on -> off, userWrite: off -> on
		s.False(p.Has(userRead))
		s.True(p.Has(userWrite))
		s.Equal(userWrite, p)
	})
}
