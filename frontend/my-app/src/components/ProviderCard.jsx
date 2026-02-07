export default function ProviderCard({ providerName, isExpanded, onToggle, onSync, lastSyncTime }) {
  const getTimeSinceSync = () => {
    if (!lastSyncTime) {
      return 'Never synced';
    }

    const now = new Date();
    const diff = Math.floor((now - lastSyncTime) / 1000);

    if (diff < 60) {
      return "Recently Synced";
    } else if (diff < 3600) {
      return `${Math.floor(diff / 60)}m ago`;
    } else if (diff < 86400) {
      return `${Math.floor(diff / 3600)}h ago`;
    } else {
      return `${Math.floor(diff / 86400)}d ago`;
    }
  };

  return (
    <div style={{ marginBottom: '10px', marginLeft: '20px' }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: '15px' }}>
        <button 
          onClick={onToggle}
          style={{
            padding: '15px',
            backgroundColor: '#3b82f6',
            color: 'white',
            border: 'none',
            borderRadius: '8px',
            cursor: 'pointer',
            fontSize: '16px',
            fontWeight: 'bold',
          }}
        >
          {providerName}
        </button>

        <span style={{ color: 'white', fontSize: '14px' }}>
          {getTimeSinceSync()}
        </span>
      </div>

      <div style={{
        maxHeight: isExpanded ? '500px' : '0',
        overflow: 'hidden',
        transition: 'max-height 0.6s ease-in-out',
      }}>
        <div style={{ 
          border: '1px solid #cbd5e1', 
          borderRadius: '8px', 
          padding: '15px',
          width: '500px',
          backgroundColor: '#1f2937',
          color: 'white',
          marginTop: '10px',
          boxSizing: 'border-box'
        }}>
          <input 
            type="text" 
            placeholder="User"
            style={{
              width: '100%',
              height: '30px',
              marginBottom: '10px',
              borderRadius: '6px',
              border: '1px solid #cbd5e1',
              fontSize: '14px',
              boxSizing: 'border-box'
            }}
          />
          
          <input 
            type="text" 
            placeholder="Password"
            style={{
              width: '100%',
              height: '30px',
              marginBottom: '10px',
              borderRadius: '6px',
              border: '1px solid #cbd5e1',
              fontSize: '14px',
              boxSizing: 'border-box'
            }}
          />

          <button 
            onClick={onSync}
            style={{
              width: '25%',
              padding: '10px',
              backgroundColor: '#10b981',
              color: 'white',
              border: 'none',
              borderRadius: '6px',
              cursor: 'pointer',
              fontSize: '14px',
              fontWeight: 'bold',
              boxSizing: 'border-box'
            }}
          >
            Sync
          </button>

          <button 
            onClick={onSync}
            style={{
              width: '25%',
              padding: '10px',
              backgroundColor: '#b91010',
              color: 'white',
              border: 'none',
              borderRadius: '6px',
              cursor: 'pointer',
              fontSize: '14px',
              fontWeight: 'bold',
              boxSizing: 'border-box'
            }}
          >
            Unlink
          </button>
        </div>
      </div>
    </div>
  );
}