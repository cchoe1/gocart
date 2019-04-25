package gocart

import "database/sql"

type WhereParam struct {
  field string
  operator string
  value string
}

type ConnectionInterface interface {
  Init() (GoCart, error)
  // @TODO: Keep the connect method here and then wrap it?  Require it to return a *sql.Db which we pass into our main GoCart struct?
  Connect() (*sql.DB, error)
  disconnect() error
}
