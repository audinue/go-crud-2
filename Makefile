dev:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ." --open-browser=false

migrate:
	tern migrate --migrations ./migrations
