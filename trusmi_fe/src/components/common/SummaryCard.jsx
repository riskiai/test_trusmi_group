export default function SummaryCard({ title, value, subtitle }) {
  return (
    <div className="summary-card">
      <div className="summary-card-title">{title}</div>
      <div className="summary-card-value">{value}</div>
      {subtitle && <div className="summary-card-subtitle">{subtitle}</div>}
    </div>
  );
}
