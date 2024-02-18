package main_test

import (
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Cmd", func() {
	var command *exec.Cmd
	var session *gexec.Session

	BeforeEach(func() {
		pathToPackageCalculatorCLI, err := gexec.Build("github.com/bonzofenix/package-calculator/cmd")
		Expect(err).NotTo(HaveOccurred())
		command = exec.Command(pathToPackageCalculatorCLI)
	})

	It("should start the server in port 8080 when no port is provided", func() {
		session = Run(command)
		Eventually(session.Out).Should(gbytes.Say("listening on port 8080"))
	})

	It("should allow port override via PORT env var", func() {
		command.Env = append(os.Environ(), "PORT=8081")
		session = Run(command)
		Eventually(session.Out).Should(gbytes.Say("listening on port 8081"))
	})

	It("should exit when it receives a  SIGINT", func() {
		session = Run(command)

		go func() {
			time.Sleep(1 * time.Second)
			session.Signal(os.Interrupt)
		}()

		Eventually(session.Out, 2*time.Second).Should(gbytes.Say("Server gracefully stopped"))
		Eventually(session).Should(gexec.Exit())
		Eventually(session).Should(gexec.Exit())
		Expect(session.ExitCode()).To(Equal(0))
	})

	AfterEach(func() {
		session.Kill()
		gexec.CleanupBuildArtifacts()
	})
})
