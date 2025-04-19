package utils

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i TokenGenerator -o ./mocks -s "_mock.go"
