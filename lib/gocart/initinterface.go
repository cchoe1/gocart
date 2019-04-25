package gocart

import "net/http"
import "log"
import "strconv"
import "fmt"
import "encoding/json"
import "strings"
import "github.com/gorilla/mux"

type CommandLine struct {

  GoCart GoCart
  // App config below related to command line usage
}
type RestApi struct {

  GoCart GoCart

}

/**
 * InitInterface will expose the API methods whether it's by CLI or REST API
 */
type InitInterface interface {
  Init() (GoCart, error)
  GetCart(id int64) error
}

// @TODO: Init() should not return any application stuff, maybe a status code if anything
func (cli *CommandLine) Init() error {
  cli.GoCart.loadConfig()

  mysql := MysqlConnection{
    host: cli.GoCart.Config.Database.Host,
    port: cli.GoCart.Config.Database.Port,
    user: cli.GoCart.Config.Database.Username,
    password: cli.GoCart.Config.Database.Password,
    database: cli.GoCart.Config.Database.Database,
    table: cli.GoCart.Config.Database.Cart.Table,
    table_index: cli.GoCart.Config.Database.Cart.Mappings.Index,
  }
  mysql.EnsureCartTable()

  // Test logic below - not required for normal use
  gocart := GoCart{
    Connection: mysql,
  }
  //GC, err := mysql.Connect()
  _ = gocart.NewCart([]Item{}, 15)

  return nil
}

/**
 * Init() will initialize the API endpoints
 */
func (rest *RestApi) Init() error {
  rest.GoCart.loadConfig()
  mysql := MysqlConnection{
    host: rest.GoCart.Config.Database.Host,
    port: rest.GoCart.Config.Database.Port,
    user: rest.GoCart.Config.Database.Username,
    password: rest.GoCart.Config.Database.Password,
    database: rest.GoCart.Config.Database.Database,
    table: rest.GoCart.Config.Database.Cart.Table,
    table_index: rest.GoCart.Config.Database.Cart.Mappings.Index,
  }
  mysql.EnsureCartTable()

  rest.GoCart = GoCart{
    Connection: mysql,
  }
  router := mux.NewRouter()

  /**
   * GET request
   */
  //http.HandleFunc("/gocart/getCart", func(w http.ResponseWriter, r *http.Request) {
  router.HandleFunc("/gocart/getCart", func(w http.ResponseWriter, r *http.Request) {
    cart_id, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64)
    if err != nil {
      panic(err)
    }
    rest.GetCart(w, r, cart_id)
  }).Methods("GET")

  /**
   * POST request
   */
  router.HandleFunc("/gocart/addToCart", func(w http.ResponseWriter, r *http.Request) {
    cart_id, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64)
    if err != nil {
      panic(err)
    }
    items_qsp := r.URL.Query().Get("items")
    item_quantity := r.URL.Query().Get("quantity")

    ids := strings.Split(items_qsp, ",")
    for _, item_id := range ids {
      item_id, err := strconv.ParseInt(item_id, 10, 64)
      if err != nil {
        panic(err)
      }
      item_quantity, err := strconv.ParseInt(item_quantity, 10, 64)
      if err != nil {
        panic(err)
      }
      rest.AddToCart(w, r, cart_id, item_id, item_quantity)
    }
    // @TODO: Print some error/success message
  }).Methods("POST")

  log.Fatal(http.ListenAndServe(":9090", router))
  return nil
}

/**
 * GET request
 * /gocart/getCart
 *
 * Query String:
 *  - cart_id int64: The ID of the cart
 * Request Body:
 *  - n/a
 */
func (rest *RestApi) GetCart(w http.ResponseWriter, r *http.Request, id int64) error {
  gc := rest.GoCart
  cart := gc.GetCart(id)

  bytes, err := json.Marshal(cart)
  if err != nil {
    panic(err)
  }

  response := string(bytes)
  fmt.Fprintln(w, response)
  return nil
}

/**
 * POST request
 * /gocart/addToCart
 *
 * Query String:
 *  - n/a
 * Request Body:
 *  -
 */
//@TODO: Refactor to only receive 1 item ID and also quantity arg
func (rest *RestApi) AddToCart(w http.ResponseWriter, r *http.Request, cart_id int64, item_id int64, quantity int64) {

  //@ TODO: Need to check for quantity and increment if necessary

  cart := rest.GoCart.GetCart(cart_id)

  item := rest.GoCart.GetItem(item_id)
  item.SetItemQuantity(quantity)

  cart.Add(*item)
  rest.GoCart.SaveCart(*cart)
}
