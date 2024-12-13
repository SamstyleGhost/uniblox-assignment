import { useCallback, useEffect, useState } from "react";
import { CartItemType, ItemType } from "../types"
import axios, { AxiosResponse } from "axios";
import { actualImages } from "../assets/images";

interface CartItemProps {
  item: CartItemType;
}

interface ItemTypeRes {
  item: ItemType;
}

const CartItem: React.FC<CartItemProps> = ({ item }) => {
  
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [currentItem, setCurrentItem] = useState<ItemType>();

  const getItemInfo = useCallback( async () => {
    try {
      const response : AxiosResponse<ItemTypeRes> = await axios.post("http://localhost:5000/api/v1/items", {
        item_id: item.item_id
      }) 
      
      setCurrentItem(response.data.item)
    } catch (error) {
      console.error(error);
    }
  }, [item.item_id])

  useEffect(() => {
    setIsLoading(true)
    getItemInfo()
    setIsLoading(false)
  }, [getItemInfo])

  if(isLoading || !currentItem) {
    return (
      <div>Loading</div>
    )
  }

  return (
    <div className="w-full flex justify-between bg-secondary rounded-2xl">
      <div className="flex flex-col p-2">
        <div className="flex md:flex-row flex-col md:gap-2 md:items-end">
          <h3 className="text-3xl">{currentItem.name}</h3>
          <h5 className="text-xl">{currentItem.vendor}</h5>
        </div>
        <div className="flex gap-2">
          <p>Quantity: {item.quantity}</p>
          <p>Cost: ${item.quantity * currentItem.price}</p>
        </div>
      </div>
      <div className="">
        <img src={actualImages[currentItem.item_id][0]} className="md:w-60 md:h-40 w-40 h-28" />
      </div>
    </div>
  );
}

export default CartItem