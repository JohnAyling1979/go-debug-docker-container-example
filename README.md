### Debugging Go application running in a Docker container

A simple example of debugging a Go application running in a Docker container.

## To Run
- docker compose up
- lauch debugger (This starts the go application)

## Notes
- To see code changes you do need to restart the debugger after each change. Not sure if there is a nodemon equivalent for Go. Will look into it.

## Debug configuration example
```
{
  "version": "0.2.0",
  "configurations": [
      {
          "name": "Debug into Docker",
          "type": "go",
          "request": "attach",
          "mode": "remote",
          "remotePath": "/go/src/go-container-debugging",
          "port": 2345,
          "host": "127.0.0.1",
          "showLog": true
      }
  ]
}
```