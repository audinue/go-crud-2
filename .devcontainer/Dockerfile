FROM golang:bookworm

RUN go install golang.org/x/tools/gopls@latest \
  && go install honnef.co/go/tools/cmd/staticcheck@latest \
  && go install github.com/jackc/tern/v2@latest \
  && go install github.com/a-h/templ/cmd/templ@latest
