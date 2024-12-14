import React from 'react'
import { Link } from 'react-router'

const Navbar: React.FC = () => {
  return (
    <nav className='bg-primary text-background py-4 w-full flex  flex-col items-center'>
      <div className='w-full max-w-page flex justify-between px-4'>
        <Link to="/">Home</Link>
        <Link to="/cart">Cart</Link>
      </div>
    </nav>
  )
}

export default Navbar
