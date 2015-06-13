moves
=====

Package moves implements a [Moves][1] API client in Go.


Usage
-----

Already have an access token?

```go
api := moves.New(token)
summary, err := api.Storyline("2014-07-20")
```

Need a token or just want more control?

```go
t := &moves.Transport{
  Key:        os.Getenv("MOVES_CLIENT_ID"),
  Secret:     os.Getenv("MOVES_CLIENT_SECRET"),
  TokenCache: moves.CacheFile("token_cache"),
}

fmt.Println("Enter the following URL in your browser.")
fmt.Println(t.AuthCodeURL("foo"))

var code string
fmt.Println("Follow the instructions in the browser.")
fmt.Println("When your browser redirects, grab the code from the query string.")
fmt.Printf("Paste the code here: ")
fmt.Scanf("%s", &code)

token, err := t.Exchange(code)
if err != nil {
  log.Fatal(err)
}

api := t.Client()
storyline, err := api.Storyline("2014-07-20")
if err != nil {
  log.Fatal(err)
}

fmt.Println(token)
fmt.Println(storyline)
```

For usage information, see the [package documentation][2] or
the [API documentation][3].


Disclaimer
----------

This package uses data from Moves but is not endorsed or certified by Moves.
Moves is a trademark of ProtoGeo Oy.


License
-------

Copyright (c) 2014 by Philip Nelson. See [LICENSE][4] for details.


[1]: https://moves-app.com
[2]: https://godoc.org/github.com/pnelson/moves
[3]: https://dev.moves-app.com/
[4]: https://github.com/pnelson/moves/blob/master/LICENSE
