import { useEffect } from "react"

import { v4 as uuidV4} from "uuid";
import { useUserContext } from "./context";

function App() {
  
  const { setCustomerID } = useUserContext();
  
  // I am making use of localStorage as a database for small use.
  // When the user enters the first time, a new uuid will be generated and set. Next time, the same id will be used for the user.
  // Wont win any awards, but gets the job done right now
  useEffect(() => {
    const customer_id = localStorage.getItem("assignment_customer_id");
    setCustomerID(customer_id);
    if(!customer_id) {
      const cid = uuidV4();
      localStorage.setItem("assignment_customer_id", cid);
      localStorage.setItem("assignment_customer_cart", "");
      setCustomerID(cid);
    }
  }, [setCustomerID])

  return (
    <div>Hello</div>
  )
}

export default App
