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
func (cli *CommandLine) Init() (GoCart) {
  mysql := MysqlConnection{
    host: "localhost", // cnf
    port: "33050", // cnf
    user: "pantheon", // cnf
    password: "pantheon", // cnf
    database: "pantheon", // cnf
    table: "gocart", // cnf ? maybe just create on initial run?
    table_index: "CartId",
  }

  // Test logic below - not required for normal use
  gocart := GoCart{
    Connection: mysql,
  }
  //GC, err := mysql.Connect();
  cart1 := gocart.NewCart([]Item{}, 0.00);
  println(cart1.GetValue());
  println(cart1.GetId());

  return gocart;
}

/**
 * Init() will initialize the API endpoints
 */
func (rest *RestApi) Init() error {
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
