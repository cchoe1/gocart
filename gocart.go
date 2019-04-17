package main

import (
  "os"
  "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

  arg1 := os.Args[1]
  arg2 := os.Args[2]

  mysql := MysqlConnection{
    Host="localhost",
    Port="3306",
    User="pantheon",
    Password="pantheon",
    Database="pantheon"
  }

  // @TODO: implement the basic steps to starting a cart

}

type MysqlConnection struct {
  var Host string;
  var Port string;
  var User string;
  var Password string;
  var Database string;
  /**
   * The database table which holds data for possible products
   */
  var table string;

  /**
   * The unique field on our table with which we can index our results by
   * - Keep this so we can use it as a type of abstraction for a filter?
   */
  var table_index string;
}

type ConnectionInterface interface {
  connect() (GoCart, error);
  disconnect() error;
}

/**
 * Any connectors implementing our ConnectionInterface should be able to connect/disconnect to the appropriate persistence layer and return a GoCart instance
 */
func (mysql MysqlConnection) connect() (GoCart, error) {
  dsn := mysql.User + ":" + mysql.Password + "@" + mysql.Host + "/" + mysql.Database;

  db, err := sql.Open("mysql", dsn);

  // @TODO: Create a new record in the database at this point and then pull the record to retrieve the ID
  cart := GoCart{
    CartId=123,
    CartValue=0,
    Db=db
  }
  return cart, err;
}

func (mysql MysqlConnection) disconnect() error {

}


/**
 * The main struct which holds our application data + config
 */

type GoCart struct {

  /**
   * The unique internal Cart ID
   */
  var CartId uint32;

  /**
   * The items that exist in the cart
   */
  var Items []Item;

  /**
   * Recalculated at every addition/subtraction of items from the cart.
   */
  var CartValue float64;

  /**
   * A db connection
   */
  var Db sql.DB;
}

/**
 * The Cart struct
 */
//type Cart struct {
//
//}

/**
 * The Item struct
 */
type Item struct {

  /**
   * The unique internal Product ID
   */
  var ProductId uint32;

  /**
   * Most likely the unique identifier
   */
  var Sku string;

  /**
   * The product's name
   */
  var ProductName string;

  /**
   * The product's description
   */
  var ProductDescription string;

  /**
   * The product's cost
   */
  var Cost float64;

  /**
   * The product's selling price
   */
  var Price float64;

}
