import axios, { AxiosResponse } from "axios";
import { useCallback, useEffect, useState } from "react";
import { useUserContext } from "../context";
import { CartResponseType } from "../types";
import { NavigateFunction, useNavigate } from "react-router";
import { CartItem, LinkButton } from "../components";

const Cartpage: React.FC = () => {
  const { customerID, customerCart, setCustomerCart } = useUserContext();
  const navigate : NavigateFunction = useNavigate()
  
  const [coupon, setCoupon] = useState<string>("");

  const getUserCart = useCallback( async () => {

    let cid : string | null;

    try {
    
      // Checking if the customerID is null
      // If it is initially, then getting the id from localStorage
      // If it is not present in localStorage, then will redirect user to homepage
      if(customerID == null) {
        cid = localStorage.getItem("assignment_customer_id")
        if (cid == null) {
          navigate("/")
        }
      } else {
        cid = customerID
      }
      const response : AxiosResponse<CartResponseType> = await axios.get(`http://localhost:5000/api/v1/user/${cid}`)
      console.log(response);
      setCustomerCart(response.data.user.cart)
    } catch (error) {
      console.error(error);
    }
  }, [customerID, navigate, setCustomerCart])

  useEffect(() => {
    getUserCart()
  }, [getUserCart])
  
  return (
    <main className="w-full px-2">
      <div className="w-full flex flex-col gap-4 pt-4">
        {customerCart &&
          customerCart.map((item, index) => (
            <CartItem key={index} item={item} />
          ))}
      </div>
      <div className="w-full flex justify-center py-4 gap-2 items-center">
        <p>Have a coupon code? Enter it here:</p>
        <input
          type="text"
          value={coupon}
          onChange={(e) => setCoupon(e.target.value)}
          className="border border-secondary rounded-lg p-2"
        />
      </div>
      <div className="w-full flex justify-center py-4">
        <LinkButton
          to="/checkout"
          additional=""
          content="Checkout"
          state={{ coupon: coupon }}
        />
      </div>
    </main>
  );
}

export default Cartpage