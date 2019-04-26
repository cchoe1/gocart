package gocart

import (
    //"fmt"
)

/**
 * The actual cart data
 * Required columns:
 *     - CartId (int64)
 *     - Items (text)
 *     - CartValue (float64)
 *     - CartOwner (int64)
 *     - Created (int64)
 *     - Updated (int64)
 */
type cart struct {

    /**
     * The unique internal Cart ID
     */
    CartId int64

    /**
     * The items that exist in the cart
     */
    Items []Item

    /**
     * Recalculated at every addition/subtraction of items from the cart.
     */
    CartValue float64

    /**
     * The owner of the cart
     * 0 - Anonymous
     * -1 - Invalid user
     */
    CartOwner int64

    /**
     * The time of cart creation
     */
    Created int64

    /**
     * The time of last cart update
     */
    Updated int64

}

type CartInterface interface {
    /* Non-public Methods */
    calculateValue() float64

    /* Public Methods */
    GetId() int64
    GetValue() float64
    GetItems() []Item
    GetOwner() int64
    GetCreated() int64
    GetUpdated() int64

    SetId(id int64) *cart
    SetValue(value float64) *cart
    SetItems(items []Item) *cart
    SetOwner(oid int64) *cart
    SetCreated(created int64) *cart
    SetUpdated(updated int64) *cart

    Add(item Item) *cart
}

// @TODO: Should remove the config value for this index and just require it to be this way
func (c cart) GetId() int64 {
    return c.CartId
}

func (c cart) GetItems() []Item {
    return c.Items
}

func (c cart) GetValue() float64 {
    return c.CartValue
}

func (c cart) GetOwner() int64 {
    return c.CartOwner
}

func (c cart) GetCreated() int64 {
    return c.Created
}

func (c cart) GetUpdated() int64 {
    return c.Updated
}

func (c *cart) SetId(id int64) *cart {
    c.CartId = id
    return c
}

func (c *cart) SetValue(value float64) *cart {
    c.CartValue = value
    return c
}

func (c *cart) SetItems(items []Item) *cart {
    c.Items = items
    return c
}

func (c *cart) SetOwner(oid int64) *cart {
    c.CartOwner = oid
    return c
}

func (c *cart) SetCreated(created int64) *cart {
    c.Created = created
    return c
}

func (c *cart) SetUpdated(updated int64) *cart {
    c.Updated = updated
    return c
}

/**
 * Adds an item to the cart
 */
func (c *cart) Add(item Item) *cart {
    current_items := c.GetItems()
    exists := false
    var current_quantity int64

    for _, it := range current_items {
        if it.GetItemId() == item.GetItemId() {
            exists = true
            current_quantity = it.GetItemQuantity()
            c.Remove(it)
            break
        }
    }

    if exists {
        new_quantity := current_quantity + item.GetItemQuantity()
        item.SetItemQuantity(new_quantity)
    } else {
        item.SetItemQuantity(item.GetItemQuantity())
    }
    c.Items = append(c.Items, item)
    c.calculateValue()
    return c
}

func (c *cart) Remove(item Item) *cart {
    var i int
    exists := false
    for index, it := range c.Items {
        if it.GetItemId() == item.GetItemId() {
            exists = true
            i = index
            break
        }
    }
    if exists == false {
        panic("The item does not exist in this cart.")
    }

    c.Items[len(c.Items)-1], c.Items[i] = c.Items[i], c.Items[len(c.Items)-1]
    c.Items = c.Items[:len(c.Items)-1]
    return c
}

func (c *cart) calculateValue() float64 {
    var value float64
    for _, item := range c.Items {
        individual_val := item.GetItemPrice()
        quantity := item.GetItemQuantity()
        value += individual_val * float64(quantity)
    }
    c.CartValue = value
    return value
}

