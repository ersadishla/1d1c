test:
	@go test -race -v 1d1c/leetcode/problem

coverage:
	@go test -timeout 30s --tags=integration -race -v -cover -coverprofile=coverage.out 1d1c/leetcode/problem 2>&1
	@go tool cover -func=coverage.out

cover:
	@go tool cover -html=coverage.out
