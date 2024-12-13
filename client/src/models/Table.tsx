
import { useRef } from 'react'
import { useGLTF } from '@react-three/drei'
import modelScene from './../assets/3d/table.glb'
import * as THREE from 'three'

const Table = (props) => {

  const tableRef = useRef<THREE.Mesh>()
  const { nodes, materials } = useGLTF(modelScene)

  return (
    <group ref={tableRef} {...props}>
      <group scale={0.01}>
        <group scale={[-1.3, 1.3, 1.3]}>
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Drawer_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[19.42, -15.586, -31.917]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Drawer_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-19.42, -15.586, 31.917]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Handle_Handle_0.geometry}
            material={materials.Handle}
            position={[23.905, -12.765, -34.545]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Handle_001_Handle_0.geometry}
            material={materials.Handle}
            position={[-23.905, -12.765, 34.545]}
            rotation={[-Math.PI / 2, 0, -Math.PI]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[31.1, -10.756, 2.261]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-29.525, -26.091, -22.309]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_002_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-31.01, -1.359, -15.885]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_003_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-5.233, -13.028, -28.043]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_004_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-39.524, -9.96, -2.267]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_005_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-56.629, -6.304, -17.001]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_006_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[31.01, -1.359, 15.885]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_007_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[5.37, -12.668, 22.837]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_008_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[29.525, -26.091, 22.309]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Inner_Shelf_009_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[56.629, -6.304, 17]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Panel_Left_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[56.216, -29.226, 0.323]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Panel_Left_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[59.158, -14.39, -0.032]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Panel_Right_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[-56.224, -29.221, 0]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Panel_Right_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-59.159, -12.429, -0.044]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Strip_Down_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[0.845, -28.898, 30.113]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Strip_Down_001_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[-0.626, -28.898, -30.112]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Strip_Top_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-20.689, -2.769, -30.285]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Strip_Top_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[20.689, -2.769, 30.285]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_1_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[56.364, -32.414, -30.719]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_1_001_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[-56.364, -32.414, 30.719]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_2_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[-56.365, -32.444, -30.739]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_2_001_Wood_Dark_0.geometry}
            material={materials.Wood_Dark}
            position={[56.365, -32.444, 30.739]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_3_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-56.825, -13.684, 32.687]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_3_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[56.825, -13.684, -32.687]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_4_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[56.143, -14.309, 32.36]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Leg_4_001_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[-56.143, -14.309, -32.36]}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Panel_Down_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            rotation={[-1.57, 0, 0]}
          />
          <mesh
            castShadow
            receiveShadow
            geometry={nodes.Table_Top_Wood_Light_0.geometry}
            material={materials.Wood_Light}
            position={[0, 0.143, 0]}
            rotation={[-1.57, 0, 0]}
          />
        </group>
      </group>
    </group>
  )
}

export default Table