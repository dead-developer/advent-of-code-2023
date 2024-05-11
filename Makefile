.PHONY: check run

day ?= 1
part ?= 1

check:
	go test -run TestSolution ./day_$(day)/part_$(part)
check-all:
	go test -run TestSolution ./day_*/part_*
run:
	go run day_$(day)/part_$(part)/solution.go
