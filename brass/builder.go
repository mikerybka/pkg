package brass

import (
	"time"

	"github.com/mikerybka/pkg/gitea"
)

type Builder struct {
	SrcDir      string
	GiteaClient *gitea.Client
}

func (b *Builder) Start() {
	for {
		// Pull the latest source code

		// Generate Go code from data/brass.dev
		// Update and build cmd
		// Upload build artifacts to builds/{os}-{arch}

		// Generate Next.js code from data/brass.dev
		// Update and build nextjs
		// Upload build artifacts to builds/nextjs

		// Build every hour
		time.Sleep(time.Hour)
	}
}
