package gocart

import (
  "database/sql"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "time"
  "strconv"
)
/**
 * The main struct which ties our application together and also serves as an high-level API
 */

type GoCart struct {

  /**
   * A potential db connection
   */
  Db *sql.DB;

  /**
   * The information related to a Mysql persistence layer
   */
  Connection MysqlConnection;

  /**
   * The application config
   */
  Config Configuration;

}

type GoCartInterface interface {
  /* Private Methods */
  loadConfig() error;

  /* Public Methods */
  // @TODO: Revise Connect() to be non-public
  Connect() error;
  GetTable() string;
  NewCart(items []Item, cart_value float64) *cart;
  GetCart(id int64) *cart;
  GetItem(id int64) *Item;
}

/**
 * Loads config from the config.yml at the project root
 */
func (gc *GoCart) loadConfig() error {

  text, err := ioutil.ReadFile("./config.yml");
  if err != nil {
    panic(err);
  }
  config := Configuration{}
  yaml_err := yaml.Unmarshal([]byte(text), &config);

  if yaml_err != nil {
    panic(yaml_err);
  }
  gc.Config = config;
  return nil;
}


func (gc *GoCart) Connect() error {
  dsn := gc.Connection.user + ":" + gc.Connection.password + "@tcp(" + gc.Connection.host + ":" + gc.Connection.port + ")/" + gc.Connection.database + "?charset=utf8";
  var db *sql.DB;
  var err error;

  db, err = sql.Open("mysql", dsn);
  gc.Db = db;
  return err;
}

func (gc GoCart) GetTable() string {
  return gc.Connection.table;
}

func (gc GoCart) NewCart(items []Item, cart_value float64) *cart {
  gc.Connect();
  defer gc.Db.Close();

  var id int64;
  var id_err error;
  created := time.Now().Unix();
  table := gc.GetTable();

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := gc.Db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items, Created) VALUES (?, ?, ?)"), "0.00", "{}", created);
  if query_error != nil {
    panic(query_error);
  }

  id, id_err = result.LastInsertId();
  if id_err != nil {
    panic(id_err);
  }

  cart := cart{
    CartId: id,
    CartValue: 0,
    Items: items,
    Created: time.Now().Unix(),
  }

  return &cart;
}

/**
 * Retrieves a cart
 */
func (gc GoCart) GetCart(id int64) *cart {
  gc.Connect();
  defer gc.Db.Close();

  table := gc.GetTable();
  // @TODO: refactor to use some abstraction that takes away the requirement for mysql
  row, err := gc.Db.Query(fmt.Sprint("SELECT CartId, Items, CartValue FROM ", table, " WHERE CartId = ?"), id);
  if err != nil {
    panic(err);
  }
  var cart_id int64;
  var cart_items []Item;
  var cart_value float64;
  row.Next();
  row.Scan(&cart_id, &cart_items, &cart_value);
  cart := cart{
    CartId: cart_id,
    Items: cart_items,
    CartValue: cart_value,
  }
  return &cart;
}

// @TODO: If we do implement the multi-DB connector, we need to change this arg to use that
// @TODO: Is there a better way to expose these config values to this method?
func (gc GoCart) SaveCart(cart cart) error {
  gc.Connect();
  gc.loadConfig();
  defer gc.Db.Close();
  // Update the updated field
  cart.SetUpdated(time.Now().Unix());

  index := gc.Config.Database.Cart.Mappings.Index;
  items := cart.GetItems();

  var item_ids string;
  // @TODO: Convert this to a method on the Item struct?
  // @TODO: Convert this to use the encode/json lib?
  for _, item := range items {
    //append(item_ids, []Item{item.GetId()}...);
    if item.GetItemId() != 0 {
    item_id := strconv.FormatInt(item.GetItemId(), 10);
      if len(item_ids) == 0 {
        item_ids = item_id;
        continue;
      }
      item_ids = item_ids + "," + item_id;
    }
  }
  item_ids = "{\"items\":[" + item_ids + "]}";
  fmt.Println(item_ids);

  _, err := gc.Db.Query(fmt.Sprint("UPDATE ", gc.Config.Database.Cart.Table, " SET ", 
    index, " = ?, ",
    "Items = ?, ",
    "CartOwner = ?, ",
    "CartValue = ?, ",
    "Created = ?, ",
    "Updated = ? ",
    "WHERE ", index, " = ", cart.GetId(),
  ),
    cart.GetId(),
    item_ids,
    cart.GetOwner(),
    cart.GetValue(),
    cart.GetCreated(),
    cart.GetUpdated(),
  );

  return err;
}


/**
 * Retrieves an individual item
 */
func (gc GoCart) GetItem(id int64) *Item {
  gc.Connect();
  // @TODO: Why do I have to load config again...
  gc.loadConfig();
  defer gc.Db.Close();

  var item_id int64;
  var item_name string;
  var item_cost float64;
  var item_price float64;

  table := gc.Config.Database.Items.Table;
  table_index := gc.Config.Database.Items.Mappings.Index;
  // @TODO: Is there a better way to compose this select statement?
  name_mapping := gc.Config.Database.Items.Mappings.Name;
  price_mapping := gc.Config.Database.Items.Mappings.Price;
  cost_mapping := gc.Config.Database.Items.Mappings.Cost;
  // @TODO: Should we implement the table index field?  well i think we have to...
  query := fmt.Sprint("SELECT ", table_index, ",", name_mapping, ",", price_mapping, ",", cost_mapping, " FROM ", table, " WHERE ", table_index, " = ", id);
  row, err := gc.Db.Query(query);

  if err != nil {
    panic(err);
  }
  row.Next();
  row.Scan(&item_id, &item_name, &item_cost, &item_price);
  // @TODO: should we implement the other fields?
  item := Item{
    item_id: item_id,
    item_name: item_name,
    item_cost: item_cost,
    item_price: item_price,
  }
  return &item;
}
