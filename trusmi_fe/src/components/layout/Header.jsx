export default function Header({ activeTab, onTabChange, onRefresh, refreshing }) {
  return (
    <header className="dashboard-header">
      <div className="header-inner">
        <div>
          <h1 className="brand-title">Dashboard KPI Marketing</h1>
          <p className="brand-subtitle">Trusmi Group</p>
        </div>

        <button
          type="button"
          className="btn-refresh"
          onClick={onRefresh}
          disabled={refreshing}
        >
          <svg
            className={`refresh-icon ${refreshing ? 'spin' : ''}`}
            width="15"
            height="15"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2.2"
            strokeLinecap="round"
            strokeLinejoin="round"
            aria-hidden="true"
          >
            <path d="M21 12a9 9 0 1 1-2.64-6.36" />
            <polyline points="21 3 21 9 15 9" />
          </svg>
          <span>{refreshing ? 'Memuat...' : 'Refresh'}</span>
        </button>
      </div>

      <nav className="tab-navigation">
        <button
          type="button"
          className={`tab-btn ${activeTab === 'dashboard1' ? 'active' : ''}`}
          onClick={() => onTabChange('dashboard1')}
        >
          Dashboard 1
        </button>
        <button
          type="button"
          className={`tab-btn ${activeTab === 'dashboard2' ? 'active' : ''}`}
          onClick={() => onTabChange('dashboard2')}
        >
          Dashboard 2
        </button>
      </nav>
    </header>
  );
}
