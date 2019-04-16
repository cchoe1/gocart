package main

import (
  "os"
  "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

  
}


//type Connection interface {
//  connect() GoCart;
//  disconnect() error;
//}
type MysqlConnection struct {
  var host string;
  var port string;
  var user string;
  var password string;
  var dsn string;
}

type ConnectionInterface interface {
  connect() GoCart;
  disconnect() error;
}

/**
 * Any connectors implementing our ConnectionInterface should be able to connect/disconnect to the appropriate persistence layer and return a GoCart instance
 */
func (mysql MysqlConnection) connect() GoCart {
  
}

func (mysql MysqlConnection) disconnect() error {

}


/**
 * The main struct which holds our application data + config
 */

type GoCart struct {

  /**
   * The database table which holds data for possible products
   */
  var db_table string;

  /**
   * The unique field on our table with which we can index our results by
   * - Keep this so we can use it as a type of abstraction for a filter?
   */
  var table_index string;

  /**
   * The attached cart which is an abstraction for the database layer
   */
  var Cart Cart;
}

/**
 * The Cart struct
 */
type Cart struct {

  /**
   * The unique internal Cart ID
   */
  var CartId uint32;

  /**
   * The items that exist in the cart
   */
  var items []Item;

  /**
   * Recalculated at every addition/subtraction of items from the cart.
   */
  var CartValue float64;
}

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
