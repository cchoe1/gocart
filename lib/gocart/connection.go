package gocart

import "database/sql"

type ConnectionInterface interface {
  Init() (GoCart, error);
  Connect() (*sql.DB, error);
  disconnect() error;
}

