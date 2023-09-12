sort-test:main.go
	go build -o $@ $^
.PHONY:clean
clean:
	rm -f test