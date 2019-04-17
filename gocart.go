package main

import (
  "os"
  "fmt"
  "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

  arg1 := os.Args[1]
  arg2 := os.Args[2]

  print(arg1, arg2);

  mysql := MysqlConnection{
    host: "localhost", // cnf
    port: "33050", // cnf
    user: "pantheon", // cnf
    password: "pantheon", // cnf
    database: "pantheon", // cnf
    table: "gocart", // cnf ? maybe just create on initial run?
    table_index: "CartId",
  }
  _, err := mysql.Connect();

  if err != nil {
    print(err);
  }


  // @TODO: implement the basic steps to starting a cart

}

type MysqlConnection struct {
  host string;
  port string;
  user string;
  password string;
  database string;
  /**
   * The database table which holds data for possible products
   */
  table string;

  /**
   * The unique field on our table with which we can index our results by
   * - Keep this so we can use it as a type of abstraction for a filter?
   */
  table_index string;

}

type ConnectionInterface interface {
  Connect() (GoCart, error);
  disconnect() error;
  GetTable() string;
}

func (mysql MysqlConnection) GetTable() string {
  return mysql.table;
}

/**
 * Any connectors implementing our ConnectionInterface should be able to connect/disconnect to the appropriate persistence layer and return a GoCart instance
 * @TODO: Rename this func. Separate the insert from the general Connect() functionality--just because we connect to the DB does not mean we want to insert a cart
 */
func (mysql MysqlConnection) Connect() (GoCart, error) {
  dsn := mysql.user + ":" + mysql.password + "@tcp(" + mysql.host + ":" + mysql.port + ")/" + mysql.database + "?charset=utf8";
  var db *sql.DB;
  var err error;

  db, err = sql.Open("mysql", dsn);

  defer db.Close();

  // @TODO: Create a new record in the database at this point and then pull the record to retrieve the ID
  var id int64;
  var id_err error;
  //stmt := `
  //INSERT INTO ? (CartValue, Items)
  //VALUES (?, ?)
  //RETURNING CartId`;
  //insert, insert_err := db.Prepare(stmt);

  //insert, insert_err := db.Prepare("INSERT INTO ? (CartValue, Items) VALUES (?, ?) RETURNING CartId;");
  table := mysql.GetTable();

  //stmt := fmt.Sprint("INSERT INTO ", table, " (CartValue, Items) VALUES (", 0.00, ", \"{}\"", ")");
  //result, query_error := db.Query(stmt);

  return cart, err;
}

func (mysql MysqlConnection) disconnect() error {

  return nil;
}


/**
 * The main struct which holds our application data + config
 */

type goCart struct {

  /**
   * The unique internal Cart ID
   */
  CartId int64;

  /**
   * The items that exist in the cart
   */
  Items []Item;

  /**
   * Recalculated at every addition/subtraction of items from the cart.
   */
  CartValue float64;

  /**
   * A db connection
   */
  Db *sql.DB;
}

type GoCartInterface interface {
  new(items []Item, cart_value float64) goCart;

}

func (cart goCart) new(items []Item, cart_value float64) goCart {

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items) VALUES (?, ?)"), "0.00", "{}");

  if id_err != nil {
    fmt.Println(id_err);
  }

  if query_error != nil {
    fmt.Println(query_error);
  }
  id, id_err = result.LastInsertId();
  fmt.Println(id);

  cart := GoCart{
    CartId: id,
    CartValue: 0,
    Db: db,
  }
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
  ProductId int64;

  /**
   * Most likely the unique identifier
   */
  Sku string;

  /**
   * The product's name
   */
  ProductName string;

  /**
   * The product's description
   */
  ProductDescription string;

  /**
   * The product's cost
   */
  Cost float64;

  /**
   * The product's selling price
   */
  Price float64;

}
