# Shopify Generic functionality

- ## Shop

  - get store date(in case if neede) for example to display the
  owner, address ...etc of the store.

- ## Products

  - fetch all products(or products of specific collection).
  - push product images(will remove the old images).
  - update product cost per item.
  - update product quantity.

- ## Locations

  - fetch all locations.

- ## Customers

  - fetch all customers.

- ## Collection

  - fetch data of specific collection(like count of products).

- ## How to use?
  
  - install the package in your project
  - init your store like:
  
  ```golang
    myStore := fkshopify.New("store user"), "store password", "store name", "api version", "web hook token")
  ```
  
  - then U can call resources functions like for example if U want to fetch all customers on store:

  ```golang
    myStore.GetAllShopifyCustomers() 
  ```

- ## Web Hooks
  
  - the WebHookHandleActionEntity method handles create, update and delete
  webhooks for all resources just need to pass the right data like the examples bellow:
  
  ```go
      type Product struct {
        ID       int    `json:"id"`
        Title    string `json:"title"`
        BodyHTML string `json:"body_html,omitempty"`
      }

      type DeleteDate struct {
        ID       int  `json:"id"`
      }

      // handle the create or update webhook
      func yourHandler() {
        var payload Product
        body := ...  //the request body
        hmacHeader := .... // ge the X-Shopify-Hmac-SHA256 header
        data, err := fkshopify.WebHookHandleActionEntity(myStore, payload, hmacHeader, body)
        // handle the error

        // U have the data sended via webhook
        fmt.Printf("%#v\n", payload.Titlt)
      }

      // handle delete webhook

      func anotherHandler() {
        var payload DeleteDate
        // the rest code is the same ....
      }
  ```
