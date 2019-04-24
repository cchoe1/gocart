package gocart

import "net/http"
import "log"
import "strconv"
import "fmt"
import "encoding/json"
import "strings"
import "github.com/gorilla/mux"

type CommandLine struct {

  GoCart GoCart;
  // App config below related to command line usage
}
type RestApi struct {

  GoCart GoCart;

}

/**
 * InitInterface will expose the API methods whether it's by CLI or REST API
 */
type InitInterface interface {
  Init() (GoCart, error);
  GetCart(id int64) error;
}

// Perform logic here
// @TODO: Init() should not return any application stuff, maybe a status code if anything
func (cli *CommandLine) Init() (GoCart) {
  cli.GoCart.loadConfig();

  mysql := MysqlConnection{
    host: cli.GoCart.Config.Database.Host,
    port: cli.GoCart.Config.Database.Port,
    user: cli.GoCart.Config.Database.Username,
    password: cli.GoCart.Config.Database.Password,
    database: cli.GoCart.Config.Database.Database,
    table: cli.GoCart.Config.Database.Cart.Table,
    table_index: cli.GoCart.Config.Database.Cart.Mappings.Index,
  }
  mysql.EnsureCartTable();

  // Test logic below - not required for normal use
  gocart := GoCart{
    Connection: mysql,
  }
  //GC, err := mysql.Connect();
  _ = gocart.NewCart([]Item{}, 0.00);

  return gocart;
}

/**
 * Init() will initialize the API endpoints
 */
func (rest *RestApi) Init() error {
  rest.GoCart.loadConfig();
  mysql := MysqlConnection{
    host: rest.GoCart.Config.Database.Host,
    port: rest.GoCart.Config.Database.Port,
    user: rest.GoCart.Config.Database.Username,
    password: rest.GoCart.Config.Database.Password,
    database: rest.GoCart.Config.Database.Database,
    table: rest.GoCart.Config.Database.Cart.Table,
    table_index: rest.GoCart.Config.Database.Cart.Mappings.Index,
  }
  mysql.EnsureCartTable();

  rest.GoCart = GoCart{
    Connection: mysql,
  }
  router := mux.NewRouter();

  /**
   * GET request
   */
  //http.HandleFunc("/gocart/getCart", func(w http.ResponseWriter, r *http.Request) {
  router.HandleFunc("/gocart/getCart", func(w http.ResponseWriter, r *http.Request) {
    cart_id, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64);
    if err != nil {
      panic(err);
    }
    rest.GetCart(w, r, cart_id);
  }).Methods("GET");

  //http.HandleFunc("/gocart/addToCart", func(w http.ResponseWriter, r *http.Request) {
  router.HandleFunc("/gocart/addToCart", func(w http.ResponseWriter, r *http.Request) {
    cart_id, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64);
    if err != nil {
      panic(err);
    }
    items_qsp := r.URL.Query().Get("items");
    rest.AddToCart(w, r, cart_id, items_qsp);
    // @TODO: Print some error/success message
  }).Methods("POST");

  //http.HandleFunc("/gocart/get")

  log.Fatal(http.ListenAndServe(":9090", router))
  return nil;
}

func (rest *RestApi) GetCart(w http.ResponseWriter, r *http.Request, id int64) error {
  gc := rest.GoCart;
  cart := gc.GetCart(id);

  bytes, err := json.Marshal(cart);
  if err != nil {
    panic(err);
  }

  response := string(bytes);
  fmt.Fprintln(w, response);
  return nil;
}

func (rest *RestApi) AddToCart(w http.ResponseWriter, r *http.Request, cart_id int64, item_ids string) {

  ids := strings.Split(item_ids, ",");

  cart := rest.GoCart.GetCart(cart_id);
  for _, item_id := range ids {
    item_id, err := strconv.ParseInt(item_id, 10, 64);
    if err != nil {
      panic(err);
    }
    item := rest.GoCart.GetItem(item_id);
    cart.Add(*item);
  }
  rest.GoCart.SaveCart(*cart);
}
