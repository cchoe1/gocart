// @TODO: Idea - we create a Persister/Entity object that will hold the configuration.    Then we can leavee these fields as non-public and then reference the Item struct on the Entity struct.    Item struct used for printing json but no persistence while Entity has the persistence.
package gocart
/**
 * The Item struct
 */
type Item struct {

    /**
     * The unique internal Product ID
     */
    ItemId int64

    /**
     * Most likely the unique identifier
     */
    ItemSku string

    /**
     * The product's name
     */
    ItemName string

    /**
     * The product's description
     */
    ItemDescription string

    /**
     * The product's cost
     */
    ItemCost float64

    /**
     * The product's selling price
     */
    ItemPrice float64

    /**
     * The quantity of items
     */
    ItemQuantity int64

    /**
     * Config necessary for mappings
     */
    config Configuration

}

/**
 * This is used for unmarshaling data from the database which holds it as a json TEXT
 */
type ItemItem struct {
    Id int64
    Quantity int64
}

type ItemInterface interface {
    /* Non-public Methods */
    /* Public Methods */
    GetItemId() int64
    GetItemSku() string
    GetItemName() string
    GetItemDescription() string
    GetItemCost() float64
    GetItemPrice() float64
    GetItemQuantity() int64

    SetItemId() *Item
    SetItemSku() *Item
    SetItemName() *Item
    SetItemDescription() *Item
    SetItemCost() *Item
    SetItemPrice() *Item
    SetItemQuantity() *Item
}

func (item Item) GetItemId() int64 {
    return item.ItemId
}

func (item Item) GetItemSku() string {
    return item.ItemSku
}

func (item Item) GetItemName() string {
    return item.ItemName
}

func (item Item) GetItemDescription() string {
    return item.ItemDescription
}

func (item Item) GetItemCost() float64 {
    return item.ItemCost
}

func (item Item) GetItemPrice() float64 {
    return item.ItemPrice
}

func (item Item) GetItemQuantity() int64 {
    return item.ItemQuantity
}

func (item *Item) SetItemId(id int64) *Item {
    item.ItemId = id
    return item
}

func (item *Item) SetItemSku(sku string) *Item {
    item.ItemSku = sku
    return item
}

func (item *Item) SetItemName(ItemName string) *Item {
    item.ItemName = ItemName
    return item
}

func (item *Item) SetItemDescription(ItemDescription string) *Item {
    item.ItemDescription = ItemDescription
    return item
}

func (item *Item) SetItemCost(ItemCost float64) *Item {
    item.ItemCost = ItemCost
    return item
}

func (item *Item) SetItemPrice(ItemPrice float64) *Item {
    item.ItemPrice = ItemPrice
    return item
}
func (item *Item) SetItemQuantity(item_quantity int64) *Item {
    item.ItemQuantity = item_quantity
    return item
}
