package gocart

import (
  "database/sql"
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
  GetTable() string;
  NewCart(items []Item, cart_value float64) *cart;
  GetCart(id int64) *cart;
}

