package gocart

type CommandLine struct {

  // App config below related to command line usage
}
type RestApi struct {

}

type InitInterface interface {
  Init() error;
}

// Perform logic here
func (cli CommandLine) Init() error {
  mysql := MysqlConnection{
    host: "localhost", // cnf
    port: "33050", // cnf
    user: "pantheon", // cnf
    password: "pantheon", // cnf
    database: "pantheon", // cnf
    table: "gocart", // cnf ? maybe just create on initial run?
    table_index: "CartId",
  }
  //if err != nil {
  //  print(err);
  //}

  // Test logic below - not required for normal use
  GC, err := mysql.Connect();
  cart1 := GC.NewCart([]gocart.Item{}, 0.00);
  println(cart1.CartValue);
  println(cart1.CartId);

  return err;
}

func (rest RestApi) Init() error {

}
