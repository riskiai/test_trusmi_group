export function getApiBaseUrl() {
  const raw = (import.meta.env.VITE_API_BASE_URL || '').trim().replace(/\/+$/, '');
  if (!raw) {
    return '/api/v1';
  }
  if (raw.endsWith('/api/v1')) {
    return raw;
  }
  if (raw.endsWith('/api')) {
    return `${raw}/v1`;
  }
  return `${raw}/api/v1`;
}

export function buildApiUrl(path = '') {
  const base = getApiBaseUrl();
  let cleanPath = String(path || '').trim();
  if (!cleanPath.startsWith('/')) {
    cleanPath = `/${cleanPath}`;
  }
  if (cleanPath.startsWith('/api/v1/')) {
    cleanPath = cleanPath.slice('/api/v1'.length);
  } else if (cleanPath === '/api/v1') {
    cleanPath = '';
  }
  return `${base}${cleanPath}`;
}

async function requestJson(path) {
  const url = buildApiUrl(path);
  const response = await fetch(url, {
    method: 'GET',
    headers: {
      Accept: 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error(`HTTP ${response.status}: Gagal memuat data dari server`);
  }

  return response.json();
}

export const kpiApi = {
  getDashboard1: () => requestJson('/dashboard-1'),
  getDashboard2: () => requestJson('/dashboard-2'),
};
