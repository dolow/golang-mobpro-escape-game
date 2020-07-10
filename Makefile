default: fmt compile run
fmt:
	go fmt main.go
compile:
	go build main.go
run:
	sudo ./main
