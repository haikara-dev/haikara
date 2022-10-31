install air to local

```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```


``` 
go run -mod=mod entgo.io/ent/cmd/ent init User
```


``` 
go generate ./ent
```