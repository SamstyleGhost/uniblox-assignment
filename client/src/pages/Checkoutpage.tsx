import { useLocation } from "react-router"

const Checkoutpage = () => {
  
  const location = useLocation()
  const data = location.state
  
  if(data.coupon == ""){
    console.log();
  }

  return (
    <div>Checkoutpage</div>
  )
}

export default Checkoutpage