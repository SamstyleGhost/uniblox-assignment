## Prereq
Need to expand this
- Require node,npm and Go installed
- My versions:
  - node: v22.11.0
  - npm: v10.9.0
  - go: v1.22.5

## Setup
1. git clone
2. cd server
3. go run main.go
4. cd client
5. npm install
6. npm run dev

## Features
API endpoints are available for:
- Adding user info to data
  - Users get a uuid assigned to them when they open the website. This id is stored in the localStorage in the frontend and is used to save user selected items in their respective carts
  - Checkout and discount functionalities also make use of this id
  - As long as the id field is not removed from the browser's local storage, every time a user visits the website, same id will be persisted
  - This is just a simulation for user auth
- Viewing all items
- Viewing singular item
- Adding item to user cart
  - Only item-id is taken from user. This would help in maintaining consistency in the data
- Checkout
  - Calculates total price of all cart items
  - For every nth order (I am using n = 3), generates a coupon code (uuid) that would be sent to the user. It is also simultaneously added to the orders.json file where the coupon code is stored in conjuction with the user that received it. Even if a coupon code is a valid one, but originates from the incorrect user, discount will not be applied
  - If the user has a valid coupon code, then the endpoint would bestow a discount of 10% on the current cart total, and the coupon code is deleted from the coupons list. This ensures that a coupon will only be used once.
  - After checkout is complete, the cart is emptied out, and the value resets to 0
  - Also checkout cannot be done if a cart is empty

## API endpoints
- Item endpoints
- User endpoints


## Photo credits:
- Photo by <a href="https://unsplash.com/@ryan_riggins?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Ryan Riggins</a> on <a href="https://unsplash.com/photos/white-ceramic-mug-on-brown-wooden-table-9v7UJS92HYc?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
- Photo by <a href="https://unsplash.com/@micheile?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">micheile henderson</a> on <a href="https://unsplash.com/photos/clothes-hanging-on-white-rack-FpPcoOAk5PI?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
- Photo by <a href="https://unsplash.com/@silvana_carlos?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Silvana Carlos</a> on <a href="https://unsplash.com/photos/green-blue-and-white-balloon-5Qajp1_80BA?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
- <a href="https://www.freepik.com/free-psd/view-sofa-interior-design-decor_320731945.htm#fromView=keyword&page=1&position=8&uuid=6ea8ce7f-abe5-4f0b-a81a-126e188d903a">Image by freepik</a>
- <a href="https://www.freepik.com/free-psd/view-sofa-interior-design-decor_320731638.htm#fromView=keyword&page=1&position=1&uuid=0b3b2ab2-a822-4cab-905e-c26f1332f3ee">Image by freepik</a>
- <a href="https://www.freepik.com/free-vector/leather-office-chair-different-angles-realistic-set-isolated-vector-illustration_63440324.htm#fromView=search&page=1&position=2&uuid=f60378ac-8d4e-4555-923c-fe7c6ee44c00">Image by macrovector on Freepik</a>
- <a href="https://www.freepik.com/free-photo/square-clock-indoors-still-life_43568881.htm#fromView=search&page=1&position=2&uuid=319e66da-ccb1-4ddc-bd3d-1d3a94116ead">Image by freepik</a>
- <a href="https://www.freepik.com/free-photo/retro-clock-alarm_1131678.htm#fromView=search&page=1&position=4&uuid=319e66da-ccb1-4ddc-bd3d-1d3a94116ead">Image by jigsawstocker on Freepik</a>
- <a href="https://www.freepik.com/free-photo/circular-clock-indoors-still-life_43568861.htm#fromView=search&page=1&position=16&uuid=319e66da-ccb1-4ddc-bd3d-1d3a94116ead">Image by freepik</a>
- <a href="https://www.freepik.com/free-photo/arranged-snack-wooden-table_1441183.htm#fromView=search&page=2&position=38&uuid=19a9e7a0-13b3-416e-b625-3d21d4e9a9e4">Image by freepik</a>
- <a href="https://www.freepik.com/free-photo/plant-vase-natural-holiday-flower_1044060.htm#fromView=search&page=5&position=26&uuid=19a9e7a0-13b3-416e-b625-3d21d4e9a9e4">Image by mrsiraphol on Freepik</a>
- <a href="https://www.freepik.com/free-ai-image/medieval-historical-rendering-knights_94937894.htm#fromView=search&page=1&position=0&uuid=f4c940d6-9c61-4d75-b5d3-b3791c152ac4">Image by freepik</a>

## 3d model credits:
- Sofa : https://skfb.ly/p8NZs

## TODO:
- Need to go over some of the status codes, mostly the error ones that are returned from helper functions