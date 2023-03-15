build:
	go build -o /tmp/matrix_traverser main.go
	chmod 700 /tmp/matrix_traverser

run:
	/tmp/matrix_traverser
