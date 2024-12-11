import { useGLTF } from '@react-three/drei'

import chairScene from './../assets/3d/office_chair.glb';
import { useRef } from 'react';

const Chair = (props) => {

  const chairRef = useRef();

  const { nodes, materials } = useGLTF(chairScene)
  return (
    <group ref={chairRef} {...props}>
      <group scale={0.01}>
        <mesh
          geometry={nodes.Cube__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0, 9.791, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Circle__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0, 9.791, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Cylinder__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0.048, 46.99, 0.131]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Plane__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0.052, 105.109, 0.162]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Plane002__0.geometry}
          material={materials['Scene_-_Root']}
          position={[-0.002, -5.142, -0.015]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Plane003__0.geometry}
          material={materials['Scene_-_Root']}
          position={[-0.002, -8.249, -0.015]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Plane004__0.geometry}
          material={materials['Scene_-_Root']}
          position={[36.465, 105.109, 0.162]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.Plane005__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0, 9.791, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
        <mesh
          geometry={nodes.cable__0.geometry}
          material={materials['Scene_-_Root']}
          position={[0, 9.791, 0]}
          rotation={[-Math.PI / 2, 0, 0]}
        />
      </group>
    </group>
  )
}

export default Chair;