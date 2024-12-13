import { useRef } from "react";
import { useGLTF } from "@react-three/drei";
import sofaScene from "./../assets/3d/sofa.glb";

const Sofa = (props) => {
  const modelRef = useRef();

  const { nodes, materials } = useGLTF(sofaScene);
  return (
    <group ref={modelRef} {...props}>
      <group position={[0, 1.428, 0.646]} scale={1.005}>
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Object_16.geometry}
          material={materials.Material}
        />
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Object_17.geometry}
          material={materials.Material}
        />
      </group>
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_4.geometry}
        material={materials.Material}
        position={[0, 1.301, 0.576]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_6.geometry}
        material={materials.Material}
        position={[0, 1.36, 0.576]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_8.geometry}
        material={materials["Material.002"]}
        position={[1.048, 1.259, 0.945]}
        rotation={[0, -Math.PI / 4, 0]}
        scale={0.014}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_10.geometry}
        material={materials.Material}
        position={[-0.003, 1.7, 0.268]}
        rotation={[-0.168, 0, 0]}
        scale={[1.013, 1, 1]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_12.geometry}
        material={materials.Material}
        position={[0.725, 1.7, 0.268]}
        rotation={[-0.168, 0, 0]}
        scale={[1.07, 1.025, 1.025]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_14.geometry}
        material={materials.Material}
        position={[-0.722, 1.7, 0.268]}
        rotation={[-0.168, 0, 0]}
        scale={[1.03, 1.014, 1.014]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_19.geometry}
        material={materials.Material}
        position={[0.722, 1.428, 0.646]}
        scale={1.006}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_21.geometry}
        material={materials.Material}
        position={[-0.721, 1.428, 0.646]}
        rotation={[-Math.PI, 0, -Math.PI]}
        scale={1.011}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_23.geometry}
        material={materials.Material}
        position={[-0.996, 1.638, 0.664]}
        rotation={[0, 0, 0.243]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_25.geometry}
        material={materials.Material}
        position={[0.997, 1.638, 0.664]}
        rotation={[0, 0, -0.288]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Object_27.geometry}
        material={materials.Material}
        position={[1.104, 1.687, 0.498]}
      />
    </group>
  );
};

export default Sofa;
