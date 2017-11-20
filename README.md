# go-honeypot
A honeypot server written in Go. 

go-honeypot listens on specified ports for any communication. When an attacker attempts to send data on one of these ports it will log relevant detail to a database.

## Running go-honeypot

1. `git clone git@github.com:Mojachieee/go-honeypot.git`
2. `cd go-honeypot`
3. `go get`

4. Create a config.json file. Formatted as follows:
```json
{
    "db": {
        "host": "myhostname.com",
        "name": "mydatabasename",
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
