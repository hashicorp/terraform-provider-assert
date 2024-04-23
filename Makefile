default: fmt lint tests-lint-fix generate copyright build install    

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

# Generate docs and copywrite headers
generate:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

copyright: ## Run copywrite (generate source code headers)
	copywrite headers

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...
	
tests-lint:
	@echo "==> Checking acceptance test terraform blocks code with terrafmt..."
	@terrafmt diff -f ./internal/provider --check --pattern '*_test.go' --quiet || (echo; \
		echo "Unexpected differences in acceptance test HCL formatting."; \
		echo "To see the full differences, run: terrafmt diff ./internal/provider --pattern '*_test.go'"; \
		echo "To automatically fix the formatting, run 'make tests-lint-fix' and commit the changes."; \
		exit 1)

tests-lint-fix:
	@echo "==> Fixing acceptance test terraform blocks code with terrafmt..."
	@find ./internal/provider -name "*_test.go" -exec sed -i ':a;N;$$!ba;s/fmt.Sprintf(`\n/fmt.Sprintf(`/g' '{}' \; # remove newlines for terrafmt
	@terrafmt fmt -f ./internal/provider --pattern '*_test.go'

.PHONY: build install lint generate fmt test testacc tests-lint tests-lint-fix copyright
