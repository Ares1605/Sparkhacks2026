import '../App.css'
import { useState } from 'react'
import Sidebar from '../components/SideBar';
import ProviderCard from '../components/ProviderCard';

function Sync() {
    const providers = ["Amazon", "Provider 2", "Provider 3"];
    const [providerStatus, setProviderStatus] = useState({
    Amazon: "notLoggedIn",
    "Provider 2": "loading",
    "Provider 3": "loggedIn"
});
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
                        onSync={() => {
                        handleSync(provider);
                        setProviderStatus({
                            ...providerStatus,
                            [provider]: "loggedIn" // or whatever makes sense after sync
                        });
                        }}
                        onUnlink={() => {
                        setProviderStatus({
                            ...providerStatus,
                            [provider]: "notLoggedIn"
                        });
                        if (expandedProvider === provider) {
                            setExpandedProvider(null);
                        }

                        }}
                        lastSyncTime={lastSyncTimes[provider]}
                        status={providerStatus[provider]}  // <-- new
                    />
                    ))}
                </div>

            <Sidebar/>
            </div>
        </div>
    );
}

export default Sync;