package gocart

import "net/http"
import "log"
import "strconv"

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
    host: cli.GoCart.Configuration.Database.Host, // cnf
    port: cli.GoCart.Configuration.Database.Port, // cnf
    user: cli.GoCart.Configuration.Database.Username, // cnf
    password: cli.GoCart.Configuration.Database.Password, // cnf
    database: cli.GoCart.Configuration.Database.Database, // cnf
    table: cli.GoCart.Configuration.Database.Cart.Table, // cnf ? maybe just create on initial run?
    table_index: cli.GoCart.Configuration.Database.Mappings.Index,
  }

  // Test logic below - not required for normal use
  gocart := GoCart{
    Connection: mysql,
  }
  //GC, err := mysql.Connect();
  cart1 := gocart.NewCart([]Item{}, 0.00);

  return gocart;
}

/**
 * Init() will initialize the API endpoints
 */
func (rest *RestApi) Init() error {
  rest.GoCart.loadConfig();
  mysql := MysqlConnection{
    host: "localhost", // cnf
    port: "33050", // cnf
    user: "pantheon", // cnf
    password: "pantheon", // cnf
    database: "pantheon", // cnf
    table: "gocart", // cnf ? maybe just create on initial run?
    table_index: "CartId",
  }

  rest.GoCart = GoCart{
    Connection: mysql,
  }

  /**
   * GET request
   */
  http.HandleFunc("/gocart/getCart", func(w http.ResponseWriter, r *http.Request) {
    cart_id, err := strconv.ParseInt(r.URL.Query().Get("cart_id"), 10, 64);
    if err != nil {
      panic(err);
    }
    rest.GetCart(w, r, cart_id);
  });
  //http.HandleFunc("/gocart/get")

  log.Fatal(http.ListenAndServe(":9090", nil))
  return nil;
}

func (rest *RestApi) GetCart(w http.ResponseWriter, r *http.Request, id int64) error {
  gc := rest.GoCart;
  _ = gc.GetCart(id);
  return nil;
}
