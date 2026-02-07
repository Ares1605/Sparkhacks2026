import './ProviderCard.css';

export default function ProviderCard({ providerName, isExpanded, onToggle, onSync, onUnlink, lastSyncTime, status }) {
  const getTimeSinceSync = () => {};

  const renderContent = () => {
    if (status === "loading") return <div>Loading...</div>;
    if (status === "notLoggedIn") return (
      <>
        <input type="text" placeholder="User" className="input" />
        <input type="text" placeholder="Password" className="input" />
        <button onClick={onSync} className="syncButton">Sync</button>
      </>
    );
    if (status === "loggedIn") return (
      <>
        <button onClick={onSync} className="syncButton">Resync</button>
        <button onClick={onUnlink} className="unlinkButton">Unlink</button>
      </>
    );
    return null;
  };

  return (
    <div className="container">
      <div className="header">
        <button onClick={onToggle} className="providerButton">{providerName}</button>
        <span className="lastSyncText">{getTimeSinceSync()}</span>
      </div>

      <div className={`expandedContainer ${isExpanded ? '' : 'collapse'}`} style={{ maxHeight: isExpanded ? '500px' : '0' }}>
        <div className="card">
          {renderContent()}
        </div>
      </div>
    </div>
  );
}