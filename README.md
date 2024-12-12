## Photo credits:
- Photo by <a href="https://unsplash.com/@ryan_riggins?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Ryan Riggins</a> on <a href="https://unsplash.com/photos/white-ceramic-mug-on-brown-wooden-table-9v7UJS92HYc?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
- Photo by <a href="https://unsplash.com/@micheile?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">micheile henderson</a> on <a href="https://unsplash.com/photos/clothes-hanging-on-white-rack-FpPcoOAk5PI?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
- Photo by <a href="https://unsplash.com/@silvana_carlos?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Silvana Carlos</a> on <a href="https://unsplash.com/photos/green-blue-and-white-balloon-5Qajp1_80BA?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>

## TODO:
- The status codes returned for some of the conditions is not right (in the API endpoints)
- Will have to optimize some endpoints as there is some repeated work being done (these would be easier and more efficient using a database)
- Right now, the add to cart enpoint sends back an item again which is added to cart, but will have to change it so that the user only sends back the itemID, which will allow the cart to be consistent with the items as the user should not be able to change the prices and other values in the payload sent