
import { useGLTF } from '@react-three/drei'

import clockScene from './../assets/3d/clock.glb';
import { useRef } from 'react';

const Clock = (props) => {

  const clockRef = useRef();

  const { nodes, materials } = useGLTF(clockScene)
  return (
    <group ref={clockRef} {...props}>
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_frame_clock_frame_0.geometry}
        material={materials.clock_frame}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_glass_glass_clock_0.geometry}
        material={materials.glass_clock}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_face_clock_face_0.geometry}
        material={materials.clock_face}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_big_hand_clock_pointer_0.geometry}
        material={materials.clock_pointer}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_second_hand_clock_pointer_0.geometry}
        material={materials.clock_pointer}
        rotation={[0, 0, -2.618]}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.clock_small_hand_clock_pointer_0.geometry}
        material={materials.clock_pointer}
        rotation={[0, 0, -Math.PI / 3]}
      />
    </group>
  )
}

export default Clock;