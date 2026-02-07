import { useState } from 'react';

export default function ProviderCard() {
  const [isExpanded, setIsExpanded] = useState(false);

  return (
    <div style={{ 
      border: '1px solid #cbd5e1', 
      borderRadius: '8px', 
      padding: '15px',
      width: '500px',
      backgroundColor: '#1f2937',
      color: 'white',
      marginBottom: '10px',
      boxSizing: 'border-box'
    }}>
      <button 
        onClick={() => setIsExpanded(!isExpanded)}
        style={{
          width: '100px',
          backgroundColor: '#3b82f6',
          color: 'white',
          border: 'none',
          borderRadius: '6px',
          cursor: 'pointer',
          fontSize: '16px',
          fontWeight: 'bold',
        }}
      >
        Amazon
      </button>

      {isExpanded && (
        <div style={{ marginTop: '16px' }}>
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

          <button style={{
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
          }}>
            Sync
          </button>
        </div>
      )}
    </div>
  );
}