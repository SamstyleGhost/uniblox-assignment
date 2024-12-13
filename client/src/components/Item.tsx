import { ItemType } from "../types"
import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import Chair from "../models/Chair";
import Sofa from "../models/Sofa";
import Knights from "../models/Knight";
import Clock from "../models/Clock";
import Table from "../models/Table";
import { useState } from "react";
import axios from "axios";
import { useUserContext } from "../context";

interface ItemProps {
  item: ItemType;
}

interface SceneProps {
  id: number
}

const SceneForThree : React.FC<SceneProps> = ({ id }) => {
  switch (id) {
    case 1:
      return <Chair scale={[0.15,0.15,0.15]} />
  
    case 2:
      return <Sofa scale={[1.8,1.8,1.8]} position={[0,-2,0]} />

    case 3:
      return <Clock scale={[4,4,4]} />

    case 4:
      return <Table scale={[2,2,2]} />

    case 5:
      return <Knights scale={[0.5,0.5,0.5]} />

    default:
      break;
  }
}

const Item : React.FC<ItemProps> = ({ item }) => {
  
  const { customerID, setCustomerCart } = useUserContext()

  const [quantity, setQuantity] = useState<number>(0);

  const handleSubmit = async (e : React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    
    try {
      const response = await axios.post(`${import.meta.env.VITE_BASE_URL}/user/add-to-cart`, {
        user_id: customerID,
        item_id: item.item_id,
        quantity: quantity,
      });
  
      console.log(response);
      setCustomerCart(response.data.cart);
    } catch (err) {
      console.error(err);
    }
  
    setQuantity(0);
  }
  
  if(!customerID) {
    return <div>Loading</div>
  }
  
  return (
    <div className="w-full h-full flex gap-4 md:flex-row flex-col pt-4">
      <div className="md:w-1/2 w-full flex flex-col py-8">
        <div className="w-full flex justify-between items-center">
          <h3 className="md:text-4xl font-medium">{item.name}</h3>
          <p className="md:text-2xl font-medium flex items-center italic">
            ${item.price}
          </p>
        </div>
        <h5 className="md:text-2xl">{item.vendor}</h5>
        <p className="py-4">{item.description}</p>
        <form onSubmit={(e) => handleSubmit(e)} className="flex gap-4">          
          <input 
            type="number" 
            value={quantity}
            placeholder="Enter quantity"
            onChange={(e) => setQuantity(e.target.value === '' ? 0 : Number(e.target.value))}
            className="border border-secondary rounded-lg p-2"
          />
          <button type="submit" className="px-4 py-2 rounded-full bg-primary text-background hover:bg-accent hover:text-black border border-primary hover:border-black font-skyer w-fit text-sm">
            Add To Cart
          </button>
        </form>
      </div>
      <div className="md:w-1/2 w-full">
        <Canvas>
          <OrbitControls />
          <directionalLight />
          <ambientLight />
          <SceneForThree id={item.item_id} />
        </Canvas>
      </div>
    </div>
  );
}

export default Item