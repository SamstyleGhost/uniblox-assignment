import { motion} from "motion/react";

import { LinkButton } from "../components";
import Sofa from "../models/Sofa";
import { Canvas, useFrame } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import * as THREE from 'three'
import { useRef } from "react";

const RotatingSofa = () => {
  const sofaRef = useRef<THREE.Mesh>(null);

  useFrame(() => {
    if (sofaRef.current) {
      sofaRef.current.rotation.y += 0.01;
    }
  });

  return <Sofa ref={sofaRef} scale={[1.8, 1.8, 1.8]} position={[0, -3, 0]} />;
}

const Homepage: React.FC = () => {

  return (
    <main className="w-full h-full overflow-hidden">
      <section className="w-full py-16 md:py-28 flex md:flex-row md:items-center md:gap-16 flex-col justify-between">
        <div className="w-full md:w-1/2 h-full relative">
          {/* <img src={sofa} className="object-contain w-full h-full" /> */}
          <Canvas>
            <OrbitControls />
            <directionalLight intensity={2}/>
            <spotLight />
            <RotatingSofa />
          </Canvas>
        </div>
        <div className="w-full md:w-1/2 flex flex-col gap-8">
          <div className="flex flex-col">
            {["One", "Stop", "Shop"].map((val, index) => (
              <motion.span key={index}
                animate={{ x: [1000, 0] }}
                transition={{ duration: 1, easing: "easeIn", delay: (index / 2) }}
                className="font-bold text-7xl md:text-[112px] font-agu md:text-left text-center"
              >{val}</motion.span>
            ))}
          </div>
        </div>
      </section>
      <section className="w-full flex justify-center">
          <div className="flex justify-center md:justify-start">
            <LinkButton to="/items" additional="" content="See all Items" />            
          </div>
      </section>
    </main>
  );
}

export default Homepage