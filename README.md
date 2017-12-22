# go-HoneyPot
A honeypot server written in Go. 

go-HoneyPot listens on specified ports for any communication. When an attacker attempts to send data on one of these ports it will log relevant details to a database.

## Running go-HoneyPot

1. `git clone git@github.com:Mojachieee/go-HoneyPot.git`
2. `cd go-HoneyPot`
3. `go get`

4. Create a config.json file. Formatted as follows:
```json
{
    "db": {
        "host": "myhostname.com",
        "name": "mydatabasename",
        "table": "mytablename",
        "username": "mydatabaseuser",
        "password": "mydatabasepassword"
    },
    "tcp": {
        "ports": [
            "1220", "5777"
        ]
    }
}
```

5. `go run honeypot.go`
