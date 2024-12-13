import { Route, Routes } from "react-router";
import { Homepage, Itemspage } from "./pages";
import { Navbar } from "./components";

function App() {

  return (
    <div className="w-full min-h-screen h-full flex flex-col items-center bg-background text-text">
      <Navbar />
      <div className="max-w-page w-full">
        <Routes>
          <Route path="/" element={<Homepage />} />
          <Route path="/items" element={<Itemspage />} />
        </Routes>
      </div>
    </div>
  );
}

export default App
