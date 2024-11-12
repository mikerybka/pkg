package nextjs

import (
	"os"
	"os/exec"
	"path/filepath"
)

func NewApp(dir string) (*App, error) {
	workdir := filepath.Dir(dir)
	err := os.MkdirAll(workdir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(dir)
	cmd := exec.Command("npx", "create-next-app@latest", "--typescript", "--eslint", "--tailwind", "--src-dir", "--app", "--import-alias", "@/*", name)
	cmd.Dir = workdir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return &App{
		Dir: dir,
	}, nil
}

type App struct {
	Dir string
}

func (a *App) SetFavicon(b []byte) error {
	path := filepath.Join(a.Dir, "src/app/favicon.ico")
	return os.WriteFile(path, b, os.ModePerm)
}

func (a *App) SetPage(path string, b []byte) error {
	path = filepath.Join(a.Dir, "src/app", path, "page.tsx")
	return os.WriteFile(path, b, os.ModePerm)
}

func (a *App) SetLayout(path string, b []byte) error {
	path = filepath.Join(a.Dir, "src/app", path, "layout.tsx")
	return os.WriteFile(path, b, os.ModePerm)
}

func (a *App) SetDockerfile(b []byte) error {
	path := filepath.Join(a.Dir, "Dockerfile")
	return os.WriteFile(path, b, os.ModePerm)
}
func (a *App) SetDockerignore(b []byte) error {
	path := filepath.Join(a.Dir, ".dockerignore")
	return os.WriteFile(path, b, os.ModePerm)
}
func (a *App) SetEnvLocal(b []byte) error {
	path := filepath.Join(a.Dir, ".env.local")
	return os.WriteFile(path, b, os.ModePerm)
}
func (a *App) SetEslintignore(b []byte) error {
	path := filepath.Join(a.Dir, ".eslintignore")
	return os.WriteFile(path, b, os.ModePerm)
}
func (a *App) SetGitignore(b []byte) error {
	path := filepath.Join(a.Dir, ".gitignore")
	return os.WriteFile(path, b, os.ModePerm)
}
func (a *App) SetTailwindConfig(b []byte) error {
	path := filepath.Join(a.Dir, "tailwind.config.ts")
	return os.WriteFile(path, b, os.ModePerm)
}
