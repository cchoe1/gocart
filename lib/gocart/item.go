package gocart
/**
 * The Item struct
 */
type Item struct {

  /**
   * The unique internal Product ID
   */
  item_id int64;

  /**
   * Most likely the unique identifier
   */
  item_sku string;

  /**
   * The product's name
   */
  item_name string;

  /**
   * The product's description
   */
  item_description string;

  /**
   * The product's cost
   */
  item_cost float64;

  /**
   * The product's selling price
   */
  item_price float64;

}

type ItemInterface interface {
  /* Non-public Methods */
  /* Public Methods */
  GetItemId() int64;
  GetItemSku() string;
  GetItemName() string;
  GetItemDescription() string;
  GetItemCost() float64;
  GetItemPrice() float64;

  SetItemId() *Item;
  SetItemSku() *Item;
  SetItemName() *Item;
  SetItemDescription() *Item;
  SetItemCost() *Item;
  SetItemPrice() *Item;
}

func (item Item) GetItemId() int64 {
  return item.item_id;
}

func (item Item) GetItemSku() string {
  return item.item_sku;
}

func (item Item) GetItemName() string {
  return item.item_name;
}

func (item Item) GetItemDescription() string {
  return item.item_description;
}

func (item Item) GetItemCost() float64 {
  return item.item_cost;
}

func (item Item) GetItemPrice() float64 {
  return item.item_price;
}

func (item *Item) SetItemId(id int64) *Item {
  item.item_id = id;
  return item;
}

func (item *Item) SetItemSku(sku string) *Item {
  item.item_sku = sku;
  return item;
}

func (item *Item) SetItemName(item_name string) *Item {
  item.item_name = item_name;
  return item;
}

func (item *Item) SetItemDescription(item_description string) *Item {
  item.item_description = item_description;
  return item;
}

func (item *Item) SetItemCost(item_cost float64) *Item {
  item.item_cost = item_cost;
  return item;
}

func (item *Item) SetItemPrice(item_price float64) *Item {
  item.item_price = item_price;
  return item;
}
