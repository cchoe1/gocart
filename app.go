/**
 * The architecture is as follows:
 *   Before any operations can be achieved, we must establish a connection with our persistence layer.
 *   This means that we must implement our ConnectionInterface using an appropriate Connector.  The implementation
 *     calls for the return value to be a GoCart struct which then holds the rest of the application functionality.
 *
 * @TODO: If we try to standardize everything from the start, we need a way to standardize the data retrieval.
 *    e.g. if the user is using PostgreSQL or NoSQL, then we need a way in the interface to define how to retrieve these records
 *      then the same methods can be called no matter what database we are using, as long as the ConnectionInterface is implemented correctly.
 *      - create basic SELECT, UPDATE, INSERT, and DELETE operations within the interface?  Then the user can implement the basic boilerplate
 *        required for those operations?
 *
 */
package main

import (
  "os"
  "fmt"
  "gocart/lib/gocart"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

  arg1 := os.Args[1]
  arg2 := os.Args[2]

  fmt.Println(arg1, arg2)

  command_line := gocart.CommandLine{}
  command_line.Init()

  rest_api := gocart.RestApi{}
  rest_api.Init()

}
