import { useEffect, useState } from "react"
import axios, { AxiosResponse } from "axios"

import { CartItemType } from "../types"
import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import Table from "../models/Table";
import { Item } from "../components";

interface ItemsResponseType {
  items : CartItemType[];
}

const Itemspage : React.FC = () => {

  const [items, setItems] = useState<CartItemType[] | null>([]);

  const getAllItems = async () => {
    try {
      const response : AxiosResponse<ItemsResponseType> = await axios.get("http://localhost:5000/api/v1/items")
      setItems(response.data.items)
    } catch (error) {
      console.error(error);
    }
  }
  
  useEffect(() => {
    getAllItems(); 
  }, [])

  return (
    <main className="w-full max-md:px-2 relative">
      {items && items.map((item, index) => (
        <Item key={index} item={item} />
      ))}
      {/* <Canvas className="bg-gray-400">
        <directionalLight />
        <spotLight />
        <OrbitControls />
        <Table position={[0, 0, 0]} scale={[4,4,4]} />
      </Canvas> */}
    </main>
  )
}

export default Itemspage