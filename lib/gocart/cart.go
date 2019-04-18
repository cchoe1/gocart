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
  /* Non-public Methods */
  calculateValue() float64;

  /* Public Methods */
  GetId() int64;
  GetValue() float64;
  GetItems() []Item;

  SetId(id int64) *cart;
  SetValue(value float64) *cart;
  SetItems(items []Item) *cart;
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

func (c *cart) SetId(id int64) *cart {
  c.cart_id = id;
  return c;
}

func (c *cart) SetValue(value float64) *cart {
  c.cart_value = value;
  return c;
}

func (c *cart) SetItems(items []Item) *cart {
  c.items = items;
  return c;
}

func (c *cart) calculateValue() float64 {
  var value float64;
  for _, item := range c.items {
    value += item.GetItemPrice();
  }
  return value;
}
