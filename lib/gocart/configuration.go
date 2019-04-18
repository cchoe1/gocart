package gocart

type Configuration struct {

  Database struct {
    Type string;
    Username string;
    Database string;
    Host string;
    Port string;

    Cart struct {
      Table string;

      Mappings struct {
        Index string;
      }
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


