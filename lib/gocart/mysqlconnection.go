package gocart

import "database/sql"

/**
 * Any connectors implementing our ConnectionInterface should be able to connect/disconnect to the appropriate persistence layer and return a GoCart instance
 * @TODO: Rename this func. Separate the insert from the general Connect() functionality--just because we connect to the DB does not mean we want to insert a cart
 */

type MysqlConnection struct {
  host string
  port string
  user string
  password string
  database string
  /**
   * The database table which holds data for possible products
   */
  table string

  /**
   * The unique field on our table with which we can index our results by
   * - Keep this so we can use it as a type of abstraction for a filter?
   */
  table_index string

}

/**
 * If an existing `gocart` table does not already exist, then we create one
 */
func (mysql MysqlConnection) EnsureCartTable() error {

  dsn := mysql.user + ":" + mysql.password + "@tcp(" + mysql.host + ":" + mysql.port + ")/" + mysql.database + "?charset=utf8"

  var db *sql.DB
  var err error

  db, err = sql.Open("mysql", dsn)
  if err != nil {
    panic(err)
  }

  check_query := `
    SELECT count(*)
    FROM gocart
  `
  _, err = db.Query(check_query)

  if err != nil {
    new_table_query := `
      CREATE TABLE gocart(
        CartId INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
        Items TEXT,
        CartValue FLOAT,
        CartOwner INT UNSIGNED,
        Created INT UNSIGNED,
        Updated INT UNSIGNED
      )
    `
    _, err = db.Exec(new_table_query)
  }
  return nil
}
