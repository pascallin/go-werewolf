# go-wolvesgame
echo said he is a good man

# structure

## folder

```shell script
|- internal
    |- game       // game logic
    |- cli        // cli interface
    |- cui        // cli ui interface
    |- commander  // game control commands
    |- werewolf   // app instance
|- pkg            
    |- tcp        // tcp client and server
```          

## development

```shell script
git clone https://github.com/pascallin/go-wolvesgame
go mod download
go run ./cmd/cli-wolvesgame/main.go
// or 
go run ./cmd/cui-wolvesgame/main.go
// PS: for win10 powershell, you should change shell text encode to unicode first
chcp 65001 
```

## cli interaction

please run help

## test

```shell script
// game test
go test ./internal/game -v
```

## TODO list

app logic
- players connection

network
- lan scan

ui
- cui design

user experience
- game local cache(tmp file), reconnect