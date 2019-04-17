package gocart
/**
 * The Item struct
 */
type Item struct {

  /**
   * The unique internal Product ID
   */
  ProductId int64;

  /**
   * Most likely the unique identifier
   */
  Sku string;

  /**
   * The product's name
   */
  ProductName string;

  /**
   * The product's description
   */
  ProductDescription string;

  /**
   * The product's cost
   */
  Cost float64;

  /**
   * The product's selling price
   */
  Price float64;

}

