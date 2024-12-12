import { Route, Routes } from "react-router";
import { Homepage } from "./pages";

function App() {
  
  // const { setCustomerID } = useUserContext();
  // const [isLoading, setIsLoading] = useState(true);
  // 
  // // I am making use of localStorage.
  // // When the user enters the first time, a new uuid will be generated and set. Next time, the same id will be used for the user.
  // // Wont win any awards, but gets the job done right now
  // useEffect(() => {
  //   const customer_id: string | null = localStorage.getItem("assignment_customer_id");
  //   setCustomerID(customer_id);
  //   
  //   // The following if condition will be true when user visits the website first time (in real worlds, when they sign up instead of log in)
  //   if(!customer_id) {
  //     const cid : UUIDTypes = uuidV4();
  //     localStorage.setItem("assignment_customer_id", cid);
  //       axios.post("http://localhost:5000/api/v1/user", {
  //         id: cid
  //       }).then(function (res) {
  //         console.log(res);
  //         setCustomerID(cid);
  //       }).catch(function (err) {
  //         // ! Technically there would be more error handling here
  //         console.log(err)
  //       })
  //   }
  //   
  //   setIsLoading(false);
  // }, [setCustomerID])
  // 
  // if(isLoading) {
  //   return <div>Loading</div>
  // }

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
