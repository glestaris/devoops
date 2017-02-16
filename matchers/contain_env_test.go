package matchers_test

import (
	"github.com/glestaris/devoops"
	. "github.com/glestaris/devoops/matchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainEnv", func() {
	It("should match environment variables that exist in the list", func() {
		envList := []string{"bananas=are", "the=best"}
		proc := devoops.Process{
			Env: envList,
		}

		Expect(proc).To(ContainEnv("bananas=are"))
	})

	It("should not match environment variables the do not exist in the list",
		func() {
			proc := devoops.Process{
				Env: []string{"bananas=are", "the=best"},
			}

			Expect(proc).NotTo(ContainEnv("bananas=best"))
		},
	)

	Context("when the actual is not a process object", func() {
		It("should error", func() {
			success, err := ContainEnv("banana=1").Match("banana=1")
			Expect(err).To(HaveOccurred())
			Expect(success).To(BeFalse())
		})
	})
})
