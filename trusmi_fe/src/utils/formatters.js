export function formatPercent(value, decimals = 2) {
  if (value === null || value === undefined || Number.isNaN(Number(value))) {
    return '0%';
  }
  const num = Number(value);
  const formatted = Number.isInteger(num) ? num.toString() : num.toFixed(decimals);
  return `${formatted}%`;
}

export function formatSignedPercent(value, decimals = 2) {
  if (value === null || value === undefined || Number.isNaN(Number(value))) {
    return '0%';
  }
  const num = Number(value);
  const formatted = Number.isInteger(num) ? num.toString() : num.toFixed(decimals);
  return `${formatted}%`;
}

export function formatNumber(value) {
  if (value === null || value === undefined || Number.isNaN(Number(value))) {
    return '0';
  }
  return Number(value).toLocaleString('id-ID');
}

export function getKpiBadgeClass(score) {
  const num = Number(score || 0);
  if (num >= 90) return 'badge-excellent';
  if (num >= 75) return 'badge-good';
  if (num >= 60) return 'badge-warning';
  return 'badge-danger';
}

export function deriveTaskStatus(deadline, aktual, apiStatus) {
  if (apiStatus && (apiStatus === 'On-Time' || apiStatus === 'Late' || apiStatus === 'Pending')) {
    return apiStatus;
  }
  if (!aktual || !deadline) return 'Pending';
  return aktual <= deadline ? 'On-Time' : 'Late';
}
