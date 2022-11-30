.PHONY: create
create:
	mkdir -p day$(day)/part$(part)
	echo "package main\n\
	\n\
	func main() {\n\
		// code goes here\n\
	}" > day$(day)/part$(part)/main.go

.PHONY: solve
solve:
	go run day$(day)/part$(part)/main.go $(input)
