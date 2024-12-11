import { useEffect } from "react"

import { v4 as uuidV4} from "uuid";
import { useUserContext } from "./context";
import { Route, Routes } from "react-router";
import { Homepage } from "./pages";

function App() {
  
  const { setCustomerID } = useUserContext();
  
  // I am making use of localStorage.
  // When the user enters the first time, a new uuid will be generated and set. Next time, the same id will be used for the user.
  // Wont win any awards, but gets the job done right now
  useEffect(() => {
    const customer_id: string | null = localStorage.getItem("assignment_customer_id");
    setCustomerID(customer_id);
    if(!customer_id) {
      const cid = uuidV4();
      localStorage.setItem("assignment_customer_id", cid);
      setCustomerID(cid);
    }
  }, [setCustomerID])

  return (
    <div className="w-full min-h-screen h-full flex justify-center bg-background text-text">
      <div className="max-w-page w-full">
        <Routes>
          <Route path="/" element={<Homepage />} />
        </Routes>
      </div>
    </div>
  );
}

export default App
