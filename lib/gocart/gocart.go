package gocart

import (
  "database/sql"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)
/**
 * The main struct which ties our application together and also serves as an high-level API
 */

type Configuration struct {
  
  Database struct {
    Type string;
    Username string;
    Database string;
    Host string;
    Port string;

    Cart struct {
      Table string;
    }
    Items struct {
      Table string;

      Mappings struct {
        Index string;
        Name string;
        Price string;
        Cost string;
      }
    }
  }
}

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
  loadConfig() *Configuration;

  /* Public Methods */
  Connect() error;
  GetTable() string;
  NewCart(items []Item, cart_value float64) *cart;
  GetCart(id int64) *cart;
  GetItem(id int64) *Item;
}

/**
 * Loads config from the config.yml at the project root
 */
func (gc *GoCart) loadConfig() *GoCart {

  text, err := ioutil.ReadFile("./config.yml");
  if err != nil {
    panic(err);
  }
  config := Configuration{}
  yaml_err := yaml.Unmarshal([]byte(text), &config);

  if yaml_err != nil {
    panic(yaml_err);
  }
  fmt.Println(config.Database.Items.Table);

  gc.Config = config;
  return gc;
}


func (gc *GoCart) Connect() error {
  dsn := gc.Connection.user + ":" + gc.Connection.password + "@tcp(" + gc.Connection.host + ":" + gc.Connection.port + ")/" + gc.Connection.database + "?charset=utf8";
  var db *sql.DB;
  var err error;

  db, err = sql.Open("mysql", dsn);
  //go_cart := GoCart{
  //  Db: db,
  //  Connection: mysql,
  //}
  gc.Db = db;
  return err;
}
func (gc GoCart) GetTable() string {
  return gc.Connection.table;
}

func (gc GoCart) NewCart(items []Item, cart_value float64) *cart {
  gc.Connect();

  var id int64;
  var id_err error;
  table := gc.GetTable();

  // @TODO: Is this the best way to prepare this statement?
  result, query_error := gc.Db.Exec(fmt.Sprint("INSERT INTO ", table, " (CartValue, Items) VALUES (?, ?)"), "0.00", "{}");
  if query_error != nil {
    fmt.Println(query_error);
  }

  id, id_err = result.LastInsertId();
  // @TODO: Panic below
  if id_err != nil {
    fmt.Println(id_err);
  }

  cart := cart{
    cart_id: id,
    cart_value: 0,
    items: items,
  }

  gc.Db.Close();
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
  fmt.Println("SCANNED");
  row.Scan(&cart_id, &cart_items, &cart_value)
  fmt.Println(cart_id);
  cart := cart{
    cart_id: cart_id,
    items: cart_items,
    cart_value: cart_value,
  }
  return &cart;
}

/**
 * Retrieves an individual item
 */
func (gc GoCart) GetItem(id int64) *Item {
  gc.Connect();
  defer gc.Db.Close();

  // @TODO: Replace with the proper table reference.  We need to know what table to use.
  table := gc.GetTable();
  row, err := gc.Db.Query(fmt.Sprint("SELECT * FROM ", table, " WHERE ItemId = ?"), id);
  if err != nil {
    panic(err);
  }
  fmt.Println(row);
  item := Item{}
  return &item;
}
