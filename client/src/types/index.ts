import { UUIDTypes } from "uuid";

interface UserContextProviderProps {
  children: React.ReactNode;
}

interface ItemType {
  item_id: number;
  name: string;
  vendor: string;
  description: string;
  images: string[];
  three_link: string;
  price: number;
  quantity: number;
}

interface CartItemType {
  item_id: number;
  quantity: number;
}

interface UserType {
  user_id: UUIDTypes;
  cart: CartItemType[];
  total_cart_value: number;
}

interface CartResponseType {
  user: UserType;
}

interface UserContextType {
  customerID: string | null;
  customerCart: CartItemType[];
  setCustomerCart: React.Dispatch<React.SetStateAction<CartItemType[]>>;
}

export { type UserContextProviderProps }
export { type ItemType }
export { type CartItemType }
export { type UserContextType }
export { type UserType }
export { type CartResponseType }