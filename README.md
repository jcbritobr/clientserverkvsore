# Client and Server Key Value Store

This is a simple key value store. You can insert, store and change data into kvstore by a client cli.
The kvstore will act like a todo task server.


* **Compile the client** - In clients folder run the following command
```
$ go build -v
```

* **Compile the server** -  In server folder run the following command
```
$ go build -v
```

* **Run the server** - 
```
$ ./server
[GIN-debug] GET    /ping                     --> github.com/jcbritobr/cstodo/server/router.setupRoute.func1 (3 handlers)
[GIN-debug] POST   /insert                   --> github.com/jcbritobr/cstodo/server/router.setupRoute.func2 (3 handlers)
[GIN-debug] GET    /load                     --> github.com/jcbritobr/cstodo/server/router.setupRoute.func3 (3 handlers)
[GIN-debug] GET    /list                     --> github.com/jcbritobr/cstodo/server/router.setupRoute.func4 (3 handlers)
[GIN-debug] POST   /doneundone               --> github.com/jcbritobr/cstodo/server/router.setupRoute.func5 (3 handlers)
5 (3 handlers)
```

* **Insert in kvstore** - Run the command 
```
$ ./client -insert -data="{\"title\":\"Cup of tea\",\"description\":\"Drink a cup of tea\", \"done\": false}"
$ {"uuid":"7536ed76-c169-4043-8106-a54592af9dcf"}
```

* **List data in kvstore** - To list data from kvstore
```
$ ./client -list
+--------------------------------------+------------+--------------------+-------+
|                 UUID                 |   TITLE    |    DESCRIPTION     | DONE  |
+--------------------------------------+------------+--------------------+-------+
| 7536ed76-c169-4043-8106-a54592af9dcf | Cup of tea | Drink a cup of tea | false |
+--------------------------------------+------------+--------------------+-------+
```

* **Switch done field** - To switch the done field value in kvstore
```
$ ./client -dud -data="{\"uuid\":\"7536ed76-c169-4043-8106-a54592af9dcf\"}"
$ {"done":true}
```

* **List commands available**
```
$ ./client -h
$ Usage of ./client:
  -data string
        Pass data as json to kvstore (default "{\"message\":\"default\"}")
  -dud
        Switches the Done field value. Needs data flag
  -insert
        Inserts data in kvstore. Needs data flag
  -list
        Lists all data from kvstore
```
