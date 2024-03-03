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
    fStore := fkshopify.New("store user"), "store password", "store name", "api version", "web hook token")
  ```
  
  - then U can call resources functions like for example if U want to fetch all customers on store:

  ```golang
    fStore.GetAllShopifyCustomers() 
  ```
