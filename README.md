## Run
- docker compose up
- lauch debgger

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