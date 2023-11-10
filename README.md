# TODO-CLI

## Play with the CLI

```bash
➜  todo git:(main) ✗ go run main.go -list
➜  todo git:(main) ✗ go run main.go -task "David and Helen"
➜  todo git:(main) ✗ go run main.go -task "Make Canada Great Again"
➜  todo git:(main) ✗ go run main.go -list                          
David and Helen
Make Canada Great Again
➜  todo git:(main) ✗ go run main.go -complete "Find an apartment in CA"
➜  todo git:(main) ✗ go run main.go -list                          
David and Helen
Make Canada Great Again
Find an apartment in CA
➜  todo git:(main) ✗ go run main.go -complete 3                        
➜  todo git:(main) ✗ go run main.go -list                              
David and Helen
Make Canada Great Again
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

Check what's inside the todo.json file
```bash
cat todo.json

[{"Task":"David and Helen","Done":false,"CreatedAt":"2023-11-11T00:24:14.467859+08:00","CompletedAt":"0001-01-01T00:00:00Z"},{"Task":"Make Canada Great Again","Done":false,"CreatedAt":"2023-11-11T00:26:33.358406+08:00","CompletedAt":"0001-01-01T00:00:00Z"}]
```

## Run the tests

```bash
go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | grep '^--- '
```