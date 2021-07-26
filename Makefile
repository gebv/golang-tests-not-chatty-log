test1:
	go test -v -count 1 -race -timeout 5m ./tests1/... --run=TestOne
test2:
	go test -v -count 1 -race -timeout 5m ./tests2/... --run=TestOne
