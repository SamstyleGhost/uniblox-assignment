import { useGLTF } from '@react-three/drei'

import chairScene from './../assets/3d/gaming_chair.glb';
import { useRef } from 'react';

const Chair = (props) => {

  const chairRef = useRef();

  const { nodes, materials } = useGLTF(chairScene)
  return (
    <group ref={chairRef} {...props}>
       <mesh
        castShadow
        receiveShadow
        geometry={
          nodes.gaming_chair_progress_uvpolySurface238_gaming_chair_progress_uvblinn6_0.geometry
        }
        material={materials.gaming_chair_progress_uvblinn6}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={
          nodes.gaming_chair_progress_uvpolySurface238_gaming_chair_progress_uvblinn4_0.geometry
        }
        material={materials.gaming_chair_progress_uvblinn4}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={
          nodes.gaming_chair_progress_uvpolySurface238_gaming_chair_progress_uvblinn4_0_1.geometry
        }
        material={materials.gaming_chair_progress_uvblinn4}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={
          nodes.gaming_chair_progress_uvpolySurface238_gaming_chair_progress_uvblinn4_0_2.geometry
        }
        material={materials.gaming_chair_progress_uvblinn4}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.gaming_chair_progress_uvpolySurface238_blinn1_0.geometry}
        material={materials.blinn1}
      />
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.gaming_chair_progress_uvpolySurface238_blinn1_0_1.geometry}
        material={materials.blinn1}
      />
    </group>
  )
}

export default Chair;