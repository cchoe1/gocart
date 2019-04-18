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

  /**
   * The owner of the cart
   */
  cart_owner int64;

  /**
   * The time of cart creation
   */
  created int64;

  /**
   *
   */
  updated int64;

}

type CartInterface interface {
  /* Non-public Methods */
  calculateValue() float64;

  /* Public Methods */
  GetId() int64;
  GetValue() float64;
  GetItems() []Item;
  GetOwner() int64;
  GetCreated() int64;
  GetUpdated() int64;

  SetId(id int64) *cart;
  SetValue(value float64) *cart;
  SetItems(items []Item) *cart;
  SetOwner(oid int64) *cart;
  SetCreated(created int64) *cart;
  SetUpdated(updated int64) *cart;

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

func (c cart) GetOwner() int64 {
  return c.cart_owner;
}

func (c cart) GetCreated() int64 {
  return c.created;
}

func (c cart) GetUpdated() int64 {
  return c.updated;
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

func (c *cart) SetOwner(oid int64) *cart {
  c.cart_owner = oid;
  return c;
}

func (c *cart) SetCreated(created int64) *cart {
  c.created = created;
  return c;
}

func (c *cart) SetUpdated(updated int64) *cart {
  c.updated = updated;
  return c;
}

func (c *cart) calculateValue() float64 {
  var value float64;
  for _, item := range c.items {
    value += item.GetItemPrice();
  }
  return value;
}
