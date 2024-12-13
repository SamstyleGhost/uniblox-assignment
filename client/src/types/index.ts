import { UUIDTypes} from "uuid";

interface UserContextProviderProps {
  children: React.ReactNode;
}

interface CartItemType {
  item_id: number;
  name: string;
  vendor: string;
  description: string;
  images: string[];
  three_link: string;
  price: number;
  quantity: number;
}

interface CartResponseType {
  user: {
    id: UUIDTypes;
    cart: CartItemType[];
  }
}

interface UserContextType {
  customerID: string | null;
  customerCart: CartItemType[];
}

export { type UserContextProviderProps }
export { type CartItemType }
export { type CartResponseType }
export { type UserContextType }