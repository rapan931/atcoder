run:
	go build main.go
	./main < input | tee output
.PHONY: run

base:
	cp -i ./_template/main.go .
.PHONY: base

m:
	./shell/move_file.sh
.PHONY: m

mc:
	./shell/move_file_contest.sh
.PHONY: mc
