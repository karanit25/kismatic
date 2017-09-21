package integration

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("kismatic docker registry feature", func() {
	BeforeEach(func() {
		dir := setupTestWorkingDir()
		os.Chdir(dir)
	})

	Describe("using an existing private docker registry", func() {
		ItOnAWS("should install successfully [slow]", func(aws infrastructureProvisioner) {
			WithInfrastructure(NodeCount{1, 1, 1, 0, 0}, Ubuntu1604LTS, aws, func(nodes provisionedNodes, sshKey string) {
				By("Installing an external Docker registry on one of the nodes")
				dockerRegistryPort := 8443
				caFile, err := deployDockerRegistry(nodes.etcd[0], dockerRegistryPort, sshKey)
				Expect(err).ToNot(HaveOccurred())
				opts := installOptions{
					dockerRegistryCAPath:   caFile,
					dockerRegistryIP:       nodes.etcd[0].PrivateIP,
					dockerRegistryPort:     dockerRegistryPort,
					dockerRegistryUsername: "kismaticuser",
					dockerRegistryPassword: "kismaticpassword",
				}
				err = installKismatic(nodes, opts, sshKey)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
