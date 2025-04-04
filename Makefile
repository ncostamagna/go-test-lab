test:
	go test -v ./suites/example/...

test-integration:
	go test -v -run Integration ./suites/example/...

test-unit:
	go test -v -run Unit ./suites/example/...



