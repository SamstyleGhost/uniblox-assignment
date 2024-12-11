
import { useRef } from 'react'
import { useGLTF } from '@react-three/drei'
import modelScene from './../assets/3d/round_table.glb'

const Table = (props) => {

  const tableRef = useRef()

  const { nodes, materials } = useGLTF(modelScene)
  return (
    <group ref={tableRef} {...props}>
      <group rotation={[-Math.PI / 2, 0, 0]} scale={100}>
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.RoundTableA_TopWood_0.geometry}
          material={materials.TopWood}
        />
        <mesh
          castShadow
          receiveShadow
          geometry={nodes.RoundTableA_LegsWood_0.geometry}
          material={materials.LegsWood}
        />
      </group>
    </group>
  )
}

export default Table