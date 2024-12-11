import { useRef } from 'react'
import { useGLTF } from '@react-three/drei'
import knightsScene from './../assets/3d/knight_statue.glb'

const Knights = (props) => {
  
  const knightsRef = useRef();

  const { nodes, materials } = useGLTF(knightsScene)
  return (
    <group ref={knightsRef} {...props}>
      <group scale={0.01}>
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Knight_Statue_Sword_Statue_01_0.geometry}
          material={materials['Statue_01']}
          position={[-200, 0, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
          scale={100}
        />
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Knight_Statue_Shield_Statue_01_0.geometry}
          material={materials['Statue_01']}
          position={[200, 0, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
          scale={100}
        />
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Sword_Statue_01_0.geometry}
          material={materials.Statue_01}
          position={[-200, 0, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
          scale={100}
        />
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.Shield_Statue_01_0.geometry}
          material={materials.Statue_01}
          position={[200, 0, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
          scale={100}
        />
      </group>
    </group>
  )
}

export default Knights