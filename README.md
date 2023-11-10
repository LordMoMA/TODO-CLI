# TODO-CLI

## Play with the CLI

```bash
➜  todo git:(main) ✗ ./todo -task "David's journey"
➜  todo git:(main) ✗ ./todo -list                  
 1: David's journey
➜  todo git:(main) ✗ ./todo -task "Helen comes with me"
➜  todo git:(main) ✗ ./todo -list                      
 1: David's journey
 2: Helen comes with me
➜  todo git:(main) ✗ ./todo -complete 1
➜  todo git:(main) ✗ ./todo -list      
X1: David's journey
 2: Helen comes with me
```

## tips

A benefit of using the flag package is that it provides automatic usage information if the user gives an invalid option or specifically requests help. 
You don’t have to do anything special to take advantage of this behavior either. 
Try it out by running your program with the -h option:

```bash
➜  todo git:(main) ✗ go build .
➜  todo git:(main) ✗ ./todo -h
Usage of ./todo:
  -complete int
        Item to be completed
  -list
        List all tasks
  -task string
        Task to be included in the TODO list
```


## Run the tests

```bash
go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | grep '^--- '
```