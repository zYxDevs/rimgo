build:
  tailwindcss -i static/tailwind.css -o static/app.css -m
  CGO_ENABLED=0 go build -o rimgo -ldflags "-X codeberg.org/rimgo/rimgo/pages.VersionInfo=$(date '+%Y-%m-%d')-$(git rev-list --abbrev-commit -1 HEAD)"

dev:
  tailwindcss -i static/tailwind.css -o static/app.css -m -w &
  go run github.com/cosmtrek/air@latest -c .air.toml