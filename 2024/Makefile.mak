# Makefile

# Define the target to run the main.go file
run:
    go run main.go day1.go day2.go day3.go day4.go day5.go

# Define a target to build the project
build:
    go build -o advent_of_code main.go

# Define a target to clean up the build artifacts
clean:
    rm -f advent_of_code