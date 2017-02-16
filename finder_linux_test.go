// +build linux

package devoops_test

import (
	"os"
	"os/exec"

	"github.com/glestaris/devoops"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("LinuxProcessFinder", func() {
	var sess *gexec.Session

	BeforeEach(func() {
		var err error

		path, err := exec.LookPath("sleep")
		Expect(err).NotTo(HaveOccurred())

		sess, err = gexec.Start(
			&exec.Cmd{
				Path: path,
				Args: []string{"sleep", "100"},
				Env: []string{
					"banana=best_fruit",
					"environ=1",
					"hello=world",
				},
				Dir: "/tmp",
			},
			GinkgoWriter, GinkgoWriter,
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		sess.Kill()
		Expect(sess.Wait()).To(gexec.Exit())
	})

	Describe("FindByPid", func() {
		var pid int

		BeforeEach(func() {
			pid = sess.Command.Process.Pid
		})

		It("should return a process with the given environment variables", func() {
			proc, err := devoops.FindByPid(pid)
			Expect(err).NotTo(HaveOccurred())

			Expect(proc.Env).To(ContainElement("banana=best_fruit"))
			Expect(proc.Env).To(ContainElement("environ=1"))
			Expect(proc.Env).To(ContainElement("hello=world"))
		})

		It("should return a process with its parent's environment variables "+
			"as well",
			func() {
				proc, err := devoops.FindByPid(pid)
				Expect(err).NotTo(HaveOccurred())

				Expect(proc.Env).To(ConsistOf(append(
					os.Environ(),
					[]string{
						"banana=best_fruit",
						"environ=1",
						"hello=world",
					}...,
				)))
			})
	})
})
