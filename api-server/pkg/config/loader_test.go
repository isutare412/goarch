package config_test

import (
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/goarch/api-server/pkg/config"
)

var _ = Describe("Loader", func() {
	It("loads config overwritten by environment variables", func() {
		os.Setenv("API_LOGGER_FORMAT", "TEST")

		cfg, err := config.LoadValidated(path.Join("..", "..", "config.yaml"))
		Expect(err).To(HaveOccurred())
		Expect(cfg.Logger.Format).To(BeEquivalentTo("TEST"))
	})
})
