package gocart

/**
 * A struct for marshaling/unmarshaling data from config.yml
 */
type Configuration struct {

  Database struct {
    Type string
    Username string
    Password string
    Database string
    Host string
    Port string

    Cart struct {
      Table string

      Mappings struct {
        Index string
      }
    }

    Items struct {
      Table string

      Mappings struct {
        Index string
        Sku string
        Name string
        Price string
        Cost string
      }
    }
  }
}


