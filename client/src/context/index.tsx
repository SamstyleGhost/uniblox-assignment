import { createContext, useContext, useEffect, useState } from "react";
import { UUIDTypes, v4 as uuidV4} from "uuid";
import axios, { AxiosResponse } from "axios";
import { UserContextProviderProps, UserContextType, CartItemType, CartResponseType } from "../types";

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserContextProvider = ({ children }: UserContextProviderProps) => {
  
  const [customerID, setCustomerID] = useState<string | null>(null);
  const [customerCart, setCustomerCart] = useState<CartItemType[]>([]);

  useEffect(() => {
    
    const cusID: string | null = localStorage.getItem("assignment_customer_id");
    setCustomerID(cusID);
    
    // The following if condition will be true when user visits the website first time (in real worlds, when they sign up instead of log in)
    // TODO: Restructure the code below to implement additional type safety and overall flow
    if(!cusID) {
      const cid : UUIDTypes = uuidV4();
      localStorage.setItem("assignment_customer_id", cid);
        axios.post("http://localhost:5000/api/v1/user", {
          user_id: cid
        }).then(res => {
          console.log(res);
          setCustomerID(cid);
        }).catch(err => {
          // ! Technically there would be more error handling here
          console.log(err)
        })
    }
    
  }, [])

  useEffect(() => {
    if(customerID) {
      axios.get(`http://localhost:5000/api/v1/user/${customerID}`)
        .then((res: AxiosResponse<CartResponseType>) => {
          setCustomerCart(res.data.user.cart);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  }, [customerID])

  return (
    <UserContext.Provider
      value={{
        customerID,
        customerCart,
        setCustomerCart
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

export const useUserContext = (): UserContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserContextProvider");
  }
  return context;
};