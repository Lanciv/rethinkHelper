RethinkHelper
============
A helper to make it easier to use gorethink.

[![wercker status](https://app.wercker.com/status/71ceee02fbdb9665d5af55aa3a855e31/m "wercker status")](https://app.wercker.com/project/bykey/71ceee02fbdb9665d5af55aa3a855e31)

# Usage

```
package main

import rh "github.com/Lanciv/rethinkHelper"

func main() {

    db, _ := rh.Connect("localhost:28015", "test")

    // Get a collection with a table name.
    Users := db.NewCollection("users")

    // Insert a record
    userOne := map[string]string{
        "UserName": "test",
        "Password": "testPass",
    }

    ids, err := Users.Insert(userOne)

    if err != nil {
        panic(err)
    }

    userOne["ID"] = ids[0]

}
```
