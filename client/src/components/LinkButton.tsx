import { Link } from 'react-router';

interface Props {
  to: string;
  additional: string;
  content: string;
  state?: object;
}

const LinkButton : React.FC<Props> = ({ to, additional, content, state }) => {
  return (
    <Link to={to} className={`px-6 py-3 rounded-full ${additional}bg-primary text-background hover:bg-accent hover:text-black border border-primary hover:border-black font-semibold font-skyer`} state={state}>{content}</Link>
  )
}

export default LinkButton