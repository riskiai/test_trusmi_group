export function LoadingState() {
  return (
    <div className="state-box">
      <div className="spinner" />
      <h3>Memuat Data Dashboard dari API GoFiber...</h3>
      <p>Mengambil perhitungan KPI dan ketepatan waktu dari database MySQL.</p>
    </div>
  );
}

export function ErrorState({ message, onRetry }) {
  return (
    <div className="state-box state-error">
      <div className="state-icon">!</div>
      <h3>Gagal Memuat Data Dashboard</h3>
      <p>{message}</p>
      {onRetry && (
        <button type="button" className="btn-refresh" onClick={onRetry}>
          Coba Lagi
        </button>
      )}
    </div>
  );
}

export function EmptyState({ message = 'Belum ada data KPI untuk ditampilkan.' }) {
  return (
    <div className="state-box">
      <h3>Data Kosong</h3>
      <p>{message}</p>
    </div>
  );
}
