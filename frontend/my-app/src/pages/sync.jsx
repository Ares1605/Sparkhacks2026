import '../App.css'
import { useState } from 'react'
import { Link } from 'react-router-dom'
import Sidebar from '../components/SideBar';
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
        <div className= "layout">
            <div className= "main">
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

            <Sidebar/>
        </div>
    </div>
  );
}

export default Sync;