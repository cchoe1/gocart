package gocart

import (
  "database/sql"
  "fmt"
)
/**
 * The main struct which ties our application together and also serves as an high-level API
 */

type GoCart struct {

  /**
   * A db connection
   */
  Db *sql.DB;

  Connection MysqlConnection;

}

type GoCartInterface interface {
  Connect() error;
  GetTable() string;
  NewCart(items []Item, cart_value float64) *cart;
  GetCart(id int64) *cart;
}

func (gc *GoCart) Connect() error {
  dsn := gc.Connection.user + ":" + gc.Connection.password + "@tcp(" + gc.Connection.host + ":" + gc.Connection.port + ")/" + gc.Connection.database + "?charset=utf8";
  var db *sql.DB;
  var err error;

  db, err = sql.Open("mysql", dsn);
  //go_cart := GoCart{
  //  Db: db,
  //  Connection: mysql,
  //}
  gc.Db = db;
  return err;
}
func (gc GoCart) GetTable() string {
  return gc.Connection.table;
}

func (gc GoCart) NewCart(items []Item, cart_value float64) *cart {
  gc.Connect();

  var id int64;
  var id_err error;
  table := gc.GetTable();

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := gc.Db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items) VALUES (?, ?)"), "0.00", "{}");
  if query_error != nil {
    fmt.Println(query_error);
  }

  id, id_err = result.LastInsertId();
  // @TODO: Panic below
  if id_err != nil {
    fmt.Println(id_err);
  }

  cart := cart{
    cart_id: id,
    cart_value: 0,
    items: items,
  }

  gc.Db.Close();
  return &cart;
}

