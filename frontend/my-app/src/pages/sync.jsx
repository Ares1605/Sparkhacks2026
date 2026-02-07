import '../App.css'
import { Link } from 'react-router-dom'
import ProviderCard from '../components/ProviderCard';

function Sync() {
  return (
    <div>
        <Link 
        to="/" style={{ color: '#60a5fa',
         position: 'absolute',
         top: '10px', left: '10px',
         textDecoration: 'none' }}> Back to Home
        </Link>

        <h1>Sync Page</h1>
        <p>This is the sync page.</p>

        <ProviderCard />
    </div>
  );
}

export default Sync;