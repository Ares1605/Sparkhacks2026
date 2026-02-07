import '../App.css'
import { useState } from 'react'
import { Link } from 'react-router-dom'
import ProviderCard from '../components/ProviderCard';

function Sync() {
  const providers = ["Amazon", "Provider 2", "Provider 3"];
  const [expandedProvider, setExpandedProvider] = useState(null);
  const [lastSyncTimes, setLastSyncTimes] = useState({});

  const handleSync = (providerName) => {
    setLastSyncTimes({
      ...lastSyncTimes,
      [providerName]: new Date()
    });
  };

  return (
    <div>
        <Link 
        to="/" style={{ color: '#60a5fa',
         position: 'absolute',
         top: '10px', left: '10px',
         textDecoration: 'none' }}> Back to Home
        </Link>

        <h1 style={{ marginLeft: '20px' }}>Sync Dashboard</h1>

        {providers.map((provider) => (
          <ProviderCard 
            key={provider} 
            providerName={provider}
            isExpanded={expandedProvider === provider}
            onToggle={() => setExpandedProvider(expandedProvider === provider ? null : provider)}
            onSync={() => handleSync(provider)}
            lastSyncTime={lastSyncTimes[provider]}
          />
        ))}
    </div>
  );
}

export default Sync;