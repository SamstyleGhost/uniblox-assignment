import { UUIDTypes} from "uuid";

interface UserContextProviderProps {
  children: React.ReactNode;
}

interface CartItemType {
  id: number;
  name: string;
  vendor: string;
  description: string;
  images: string[];
  three_link: string;
  price: number;
  quantitu: number;
}

interface CartResponseType {
  success: boolean;
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