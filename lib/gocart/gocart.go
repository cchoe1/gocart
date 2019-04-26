package gocart

import (
  "database/sql"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  //"io"
  "time"
  "encoding/json"
)
/**
 * The main struct which ties our application together and also serves as an high-level API
 */

type GoCart struct {

  /**
   * A potential db connection
   */
  Db *sql.DB

  /**
   * The information related to a Mysql persistence layer
   */
  Connection MysqlConnection

  /**
   * The application config
   */
  Config Configuration

}

type GoCartInterface interface {
  /* Private Methods */
  loadConfig() error

  /* Public Methods */
  // @TODO: Revise Connect() to be non-public
  Connect() error
  GetTable() string
  NewCart(items []Item, cart_value float64) *cart
  GetCart(id int64) *cart
  GetItem(id int64) *Item
}

/**
 * Loads config from the config.yml at the project root
 */
func (gc *GoCart) loadConfig() error {

  text, err := ioutil.ReadFile("./config.yml")
  if err != nil {
    panic(err)
  }
  config := Configuration{}
  yaml_err := yaml.Unmarshal([]byte(text), &config)

  if yaml_err != nil {
    panic(yaml_err)
  }
  gc.Config = config
  return nil
}


func (gc *GoCart) Connect() error {
  dsn := gc.Connection.user + ":" + gc.Connection.password + "@tcp(" + gc.Connection.host + ":" + gc.Connection.port + ")/" + gc.Connection.database + "?charset=utf8"
  var db *sql.DB
  var err error

  db, err = sql.Open("mysql", dsn)
  gc.Db = db
  return err
}

func (gc GoCart) GetTable() string {
  return gc.Connection.table
}

func (gc GoCart) NewCart(items []Item, owner int64) *cart {
  gc.Connect()
  defer gc.Db.Close()

  var id int64
  var id_err error
  created := time.Now().Unix()
  table := gc.GetTable()

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := gc.Db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items, CartOwner, Created, Updated) VALUES (?, ?, ?, ?, ?)"), "0.00", "[]", 0, created, 0)
  if query_error != nil {
    panic(query_error)
  }

  id, id_err = result.LastInsertId()
  if id_err != nil {
    panic(id_err)
  }

  cart := cart{
    CartId: id,
    CartValue: 0.00,
    Items: items,
    Created: created,
    CartOwner: owner,
  }
  cart.calculateValue()
  return &cart
}

/**
 * Retrieves a cart
 */
func (gc GoCart) GetCart(id int64) *cart {
  gc.Connect()
  defer gc.Db.Close()

  table := gc.GetTable()
  // @TODO: refactor to use some abstraction that takes away the requirement for mysql
  row, err := gc.Db.Query(fmt.Sprint("SELECT CartId, Items, CartValue, CartOwner, Created, Updated FROM ", table, " WHERE CartId = ?"), id)
  if err != nil {
    panic(err)
  }
  defer row.Close()
  var cart_id int64
  var cart_items []byte
  var cart_value float64
  var cart_owner int64
  var created int64
  var updated int64

  row.Next()
  err = row.Scan(&cart_id, &cart_items, &cart_value, &cart_owner, &created, &updated)
  if err != nil {
    panic(err)
  }

  item_list := []ItemItem{}
  err = json.Unmarshal(cart_items, &item_list)
  if err != nil {
    panic(err)
  }

  items := []Item{}
  for _, item_item := range item_list {
    item := gc.GetItem(item_item.Id)
    item.SetItemQuantity(item_item.Quantity)
    items = append(items, *item)
  }
  cart := cart{
    CartId: cart_id,
    Items: items,
    CartValue: cart_value,
    CartOwner: cart_owner,
    Created: created,
    Updated: updated,
  }
  cart.calculateValue()
  // @TODO: YUCK.  Need to refactor so that our entity structs have access to Save() itself without calling this GoCart struct
  gc.SaveCart(cart)

  return &cart
}

// @TODO: If we do implement the multi-DB connector, we need to change this arg to use that
// @TODO: Is there a better way to expose these config values to this method?
func (gc GoCart) SaveCart(cart cart) error {
  gc.Connect()
  gc.loadConfig()
  defer gc.Db.Close()
  // Update the updated field
  cart.SetUpdated(time.Now().Unix())

  index := gc.Config.Database.Cart.Mappings.Index
  items := cart.GetItems()

  var item_info []ItemItem
  // @TODO: Convert this to a method on the Item struct?
  // @TODO: Convert this to use the encode/json lib?
  for _, item := range items {
    //append(item_ids, []Item{item.GetId()}...)
    if item.GetItemId() != 0 {
      item_item := ItemItem{
        Id: item.GetItemId(),
        Quantity: item.GetItemQuantity(),
      }
      item_info = append(item_info, item_item)
    }
  }
  //item_ids = "{\"items\":[" + item_ids + "]}"
  marshaled_items, err := json.Marshal(item_info)
  if err != nil {
    panic(err)
  }

  _, err = gc.Db.Query(fmt.Sprint("UPDATE ", gc.Config.Database.Cart.Table, " SET ",
    index, " = ?, ",
    "Items = ?, ",
    "CartOwner = ?, ",
    "CartValue = ?, ",
    "Created = ?, ",
    "Updated = ? ",
    "WHERE ", index, " = ", cart.GetId(),
  ),
    cart.GetId(),
    marshaled_items,
    cart.GetOwner(),
    cart.GetValue(),
    cart.GetCreated(),
    cart.GetUpdated(),
  )

  return err
}

/**
 * Retrieves an individual item
 */
func (gc GoCart) GetItem(id int64) *Item {
  gc.Connect()
  // @TODO: Why do I have to load config again...
  gc.loadConfig()
  defer gc.Db.Close()

  var item_id int64
  var item_sku string;
  var item_name string
  var item_cost float64
  var item_price float64

  table := gc.Config.Database.Items.Table
  // @TODO: Is there a better way to compose this select statement?
  table_index := gc.Config.Database.Items.Mappings.Index
  sku_mapping := gc.Config.Database.Items.Mappings.Sku
  name_mapping := gc.Config.Database.Items.Mappings.Name
  price_mapping := gc.Config.Database.Items.Mappings.Price
  cost_mapping := gc.Config.Database.Items.Mappings.Cost
  // @TODO: Should we implement the table index field?  well i think we have to...
  query := fmt.Sprint("SELECT ", table_index, ",", sku_mapping, ",", name_mapping, ",", price_mapping, ",", cost_mapping, " FROM ", table, " WHERE ", table_index, " = ", id)
  row, err := gc.Db.Query(query)

  if err != nil {
    panic(err)
  }
  row.Next()
  row.Scan(&item_id, &item_sku, &item_name, &item_price, &item_cost)
  // @TODO: should we implement the other fields?
  item := Item{
    ItemId: item_id,
    ItemSku: item_sku,
    ItemName: item_name,
    ItemCost: item_cost,
    ItemPrice: item_price,
  }
  return &item
}
