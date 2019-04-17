package gocart

type CommandLine struct {

  // App config below related to command line usage
}
type RestApi struct {
  

}

type InitInterface interface {
  Init() (GoCart, error);
}

// Perform logic here
func (cli CommandLine) Init() (GoCart) {
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

func (rest RestApi) Init() error {

  return nil;
}
