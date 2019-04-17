package gocart

import "fmt"

/**
 * The actual cart data
 */
type cart struct {

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

}

type CartInterface interface {

}

func (gc GoCart) GetTable() string {
  return gc.Connection.table;
}
/**
 * Constructors
 */
func (gc GoCart) NewCart(items []Item, cart_value float64) *cart {
  var id int64;
  var id_err error;
  table := gc.GetTable();

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := gc.Db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items) VALUES (?, ?)"), "0.00", "{}");

  // @TODO: Panic below
  if id_err != nil {
    fmt.Println(id_err);
  }
  if query_error != nil {
    fmt.Println(query_error);
  }

  id, id_err = result.LastInsertId();

  cart := cart{
    CartId: id,
    CartValue: 0,
    Items: items,
  }
  return &cart;
}

