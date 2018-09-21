# Lazada Client


## Usage
```
c := client.NewLazadaClient("APP_KEY", "APP_SECRET")
c.SetAccessToken("ACCESS_TOKEN")
```

### Get Orders
```
orders, err := c.GetOrders(client.GetOrdersParams{
  CreatedAfter: "2018-09-19T09:00:00+07:00",

  Limit:  "1",
  Offset: "0",
})

if err != nil {
	panic(err)
}
```