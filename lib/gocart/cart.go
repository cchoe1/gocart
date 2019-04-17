package gocart

/**
 * The actual cart data
 */
type cart struct {

  /**
   * The unique internal Cart ID
   */
  cart_id int64;

  /**
   * The items that exist in the cart
   */
  items []Item;

  /**
   * Recalculated at every addition/subtraction of items from the cart.
   */
  cart_value float64;

}

type CartInterface interface {
  GetId() int64;
  GetValue() float64;
  GetItems() []Item;
}

func (c cart) GetId() int64 {
  return c.cart_id;
}

func (c cart) GetItems() []Item {
  return c.items;
}

func (c cart) GetValue() float64 {
  return c.cart_value;
}
