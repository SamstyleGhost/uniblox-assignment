import axios, {AxiosResponse} from "axios"
import { CartItemType, CouponType } from "../types"
import { useUserContext } from "../context";
import { useEffect, useState } from "react";
import { UUIDTypes } from "uuid";

interface CheckoutResponseType {
  cart: CartItemType;
  original_price: number;
  discount: number;
  total_price: number;
  coupon: UUIDTypes;
}

const Checkoutpage = () => {

  const { customerID } = useUserContext()
  
  const [coupons, setCoupons] = useState<CouponType[]>([]);
  
  
  useEffect(() => {
    const checkForCoupons = async () => {
      try {
        const response: AxiosResponse<CouponType[]> = await axios.post(`${import.meta.env.VITE_BASE_URL}/user/coupons`,{
            user_id: customerID,
          });

        setCoupons(response.data);
      } catch (error) {
        console.error(error);
      }
    };

    checkForCoupons()
  }, [customerID]);
  
  useEffect(() => {
    const checkout = async () => {
      try {
        const response: AxiosResponse<CheckoutResponseType> = await axios.post(`${import.meta.env.VITE_BASE_URL}/user/checkout`, {
          user_id: customerID,
          coupon_code: coupons.length > 0 ? coupons[0] : "00000000-0000-0000-0000-000000000000"
        }) 
        
        console.log(response);
      } catch (error) {
        console.error(error);
      }

    };
    checkout()
  }, [coupons, customerID])

  return (
    <div>Checkoutpage</div>
  )
}

export default Checkoutpage