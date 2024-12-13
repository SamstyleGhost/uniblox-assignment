import React from 'react'

interface Props {
  additional: string;
  content: string
}

const Button : React.FC<Props> = ({ additional, content }) => {
  return (
    <button className={`px-6 py-3 rounded-full ${additional}bg-primary text-background hover:bg-accent hover:text-black border border-primary hover:border-black font-semibold font-skyer`}>{content}</button>
  )
}

export default Button