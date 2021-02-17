# League Backend Challenge

Run web server
```
go run .
```

Send request with attached example file
```
curl -F 'file=@matrix.csv' "localhost:3000/echo"
```

Testing services
```
go test ./services
```