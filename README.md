# cli_tool_tree
Written in Golang

To run a program you need Go installed on your PC

How to start:
1. Clone repo in your go/src
"go run main.go"

By default files are shown
You can turn this of by changing "showFiles" default value to "false", and further enabling it with -f parametr when launching:
"go run main.go -f"

By default directory to be treed is current directory. You can give with -dir necessary directory to be treed.
"go run main.go -dir=/" - for root in mac/linux
