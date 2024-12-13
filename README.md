## Prerequisite to run this
- Require node,npm and Go installed
- My versions:
  - node: v22.11.0
  - npm: v10.9.0
  - go: v1.22.5
- Also, the env file in the client is NOT a mistake

## Setup
1. git clone
2. cd server
3. go run main.go
4. cd ..
5. cd client
6. npm install
7. npm run dev

## Features
- The backend of the project is built with [Go](https://go.dev/doc/)
  - Used [fiber](https://github.com/gofiber/fiber) as the backend framework
- The frontend is built using React with [Vite](https://vite.dev/guide/)
  - Also integrated some 3d models using [ThreeJS](https://threejs.org/docs/index.html#manual/en/introduction/Creating-a-scene) & [R3F (react-three-fiber)](https://github.com/pmndrs/react-three-fiber)
  - Uses [TailwindCSS](https://tailwindcss.com/docs/installation/using-postcss) for styling
  - Uses Typescript
- There is no external database, however I have added persistency through json files as well as some use of browser localStorage


## Endpoints
API endpoints are available for:

- ```GET /items```
  - Sends all items from the items.json file
  - Response:
    ``` 
    {
    "items": [
        {
            "item_id": 1,
            "name": "Office Chair",
            "vendor": "OfficeCo Supplies",
            "description": "A comfortable and ergonomic office chair with adjustable height and wheels for mobility, ideal for long hours at the desk.",
            "images": [
                "chair1",
                "chair2"
            ],
            "three_link": "https://fab.com/s/2cf83e2ac5e8",
            "price": 49.99,
            "quantity": 25
        },
        ...
      ]
    }
    ```
- ```POST /items ```
  - Sends only the selected item
  - Request: 
  ```
  {
    "item_id": 3
  }
  ```
  - Response: 
  ```
  {
    "item": {
        "item_id": 3,
        "name": "Clock",
        "vendor": "Timekeepers Ltd.",
        "description": "A sleek and modern wall clock with a silent mechanism, suitable for homes or offices.",
        "images": [
            "clock1",
            "clock2",
            "clock3"
        ],
        "three_link": "https://fab.com/s/7534f812496b",
        "price": 19.99,
        "quantity": 100
    }
  }
  ```

- ``` GET /user/:id ```
  - Sends the user's cart
  - Response: 
  ```
  {
    "user": {
        "user_id": "85090ac9-9066-4979-8a64-91b4a6016bbc",
        "cart": [
            {
                "item_id": 2,
                "quantity": 6
            }
        ],
        "total_cart_value": 539.94
    }
  }
  ```

- ```POST /user ```
  - Adds the current user to the users file
  - Request: ``` user_id ```
  
- ```POST /user/add-to-cart```
  - Adds current item to the user's cart
  - Request: 
  ```
  {
    "user_id": "39490ad6-7cc2-42f6-8887-3f29749dbb42",
    "item_id": 5,
    "quantity": 6
  }
  ```
  - Response:
  ```
  {
    "cart": [
        {
            "item_id": 5,
            "quantity": 6
        }
    ]
  }
  ```
  
- ```POST /user/checkout```
  - Request1: 
  ```
  {
    "user_id": "39490ad6-7cc2-42f6-8887-3f29749dbb42"
  }
  ```
  - Response1:
  ```
  {
    "cart": [
        {
            "item_id": 5,
            "quantity": 6
        }
    ],
    "coupon": "00000000-0000-0000-0000-000000000000",
    "discount": 0,
    "original_price": 599.94,
    "total_price": 599.94
  }
  ```
  - Request2:
  ```
  {
    "user_id": "8d27c3f7-bc80-4994-bab7-e27db5f102c5",
    "coupon_code": "1972fcab-3b1c-4cae-b3db-31704b931973"
  }
  ```
  - Response2:
  ```
  {
    "cart": [
        {
            "item_id": 3,
            "quantity": 2
        },
        {
            "item_id": 4,
            "quantity": 5
        }
    ],
    "coupon": "00000000-0000-0000-0000-000000000000",
    "discount": 78.993004,
    "original_price": 789.93,
    "total_price": 710.937
  }
  ```
  - The nil uuid present in the *coupon* field signifies that the current transaction did not generate any coupon
  - Every *nth* transaction generates a discount coupon, which can be redeemed as shown in the above request
  
- ```POST /user/coupons```
  - Request:
  ```
  {
    "user_id": "847f38fa-bf27-43b1-a871-93c1810098c8"
  }
  ```
  - Response:
  ```
  [
    {
        "user_id": "847f38fa-bf27-43b1-a871-93c1810098c8",
        "coupon_code": "fc1d6f3e-91ff-4fd3-a809-a9d702485595"
    },
    {
        "user_id": "847f38fa-bf27-43b1-a871-93c1810098c8",
        "coupon_code": "7259d3fc-5e7f-4e23-8aae-1c5b81e03663"
    }
  ]
  ```


## Additional Info
- Adding user info to data
  - Users get a uuid assigned to them when they open the website. This id is stored in the localStorage in the frontend and is used to save user selected items in their respective carts
  - Checkout and discount functionalities also make use of this id
  - As long as the id field is not removed from the browser's local storage, every time a user visits the website, same id will be persisted
  - This is just a simulation for user auth
- Adding item to user cart
  - Only item-id is taken from user. This would help in maintaining consistency in the data
- Checkout
  - Calculates total price of all cart items
  - For every nth order (I am using n = 3), generates a coupon code (uuid) that would be sent to the user. It is also simultaneously added to the orders.json file where the coupon code is stored in conjuction with the user that received it. Even if a coupon code is a valid one, but originates from the incorrect user, discount will not be applied
  - If the user has a valid coupon code, then the endpoint would bestow a discount of 10% on the current cart total, and the coupon code is deleted from the coupons list. This ensures that a coupon will only be used once.
  - After checkout is complete, the cart is emptied out, and the value resets to 0
  - Also checkout cannot be done if a cart is empty

## Needed optimizations
- There is some redundant & repeated code in the server. Almost all of it is due to the work with JSON files. A database would fix that. And building an ORM simulating middleware would be a bit more time-consuming I suppose
- In hindsight, having a seperate file for coupons was a small mistake, just an additional field in the users file would have done the job better
- On the client side, havent really compressed the images and models much. Also, right now everything resides on the server itself, but using a storage system would ease the burden and also make working with the assets better.


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