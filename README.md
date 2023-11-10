# TODO-CLI

## Run the CLI

First, build the CLI:

```bash
go run main.go Make Canada Great Again
```

Then, run the CLI:

```bash
go run main.go
```

Output:
```bash
Make Canada Great Again
```

Check what's inside the todo.json file
```bash
cat todo.json

[{"Task":"David and Helen","Done":false,"CreatedAt":"2023-11-11T00:24:14.467859+08:00","CompletedAt":"0001-01-01T00:00:00Z"},{"Task":"Make Canada Great Again","Done":false,"CreatedAt":"2023-11-11T00:26:33.358406+08:00","CompletedAt":"0001-01-01T00:00:00Z"}]
```