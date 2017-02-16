package matchers

import (
	"fmt"

	"github.com/glestaris/devoops"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type containEnvMatcher struct {
	expected string
}

func ContainEnv(expected string) types.GomegaMatcher {
	return &containEnvMatcher{
		expected: expected,
	}
}

func (m *containEnvMatcher) Match(actual interface{}) (
	success bool, err error,
) {
	v, ok := actual.(devoops.Process)
	if !ok {
		return false, fmt.Errorf(
			"Can only match against devoops.Process instances",
		)
	}

	for _, e := range v.Env {
		if e == m.expected {
			return true, nil
		}
	}
	return false, nil
}

func (m *containEnvMatcher) FailureMessage(actual interface{}) string {
	return format.Message(
		actual, "to have the environment variable", m.expected,
	)
}

func (m *containEnvMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(
		actual, "not to have the environment variable", m.expected,
	)
}
