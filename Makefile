SRC_DIRS := cmd pkg # directories which hold app source (not vendored)

all: swagger mock

swagger:
	swag init -g cmd/service/app.go -o cmd/service/docs

mock:
	rm -f pkg/http/adding/mock_dataaccessor.go
	mockgen -package adding -self_package github.com/breiting/bookless/pkg/http/adding -destination pkg/http/adding/mock_dataaccessor.go github.com/breiting/bookless/pkg/http/adding DataAccessor
