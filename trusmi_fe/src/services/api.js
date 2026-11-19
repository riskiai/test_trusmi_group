/**
 * Menormalisasi VITE_API_BASE_URL agar:
 * - Di Development (VITE_API_BASE_URL=http://localhost:8090) -> http://localhost:8090/api/v1
 * - Di Production (VITE_API_BASE_URL=/api/v1 atau "")        -> /api/v1 (relative URL pada domain yang sama)
 * - Tidak pernah terjadi duplikasi prefix /api/v1/api/v1
 */
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
  // Cegah duplikasi jika caller tidak sengaja mengirim /api/v1/...
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
  getHealth: () => requestJson('/health'),
  getSummary: (karyawan = '') =>
    requestJson(karyawan ? `/kpi/summary?karyawan=${encodeURIComponent(karyawan)}` : '/kpi/summary'),
  getOntimeLate: () => requestJson('/kpi/ontime-late'),
  getTasks: ({ karyawan = '', kpi = '', status = '' } = {}) => {
    const params = new URLSearchParams();
    if (karyawan) params.set('karyawan', karyawan);
    if (kpi) params.set('kpi', kpi);
    if (status) params.set('status', status);
    const qs = params.toString();
    return requestJson(qs ? `/kpi/tasks?${qs}` : '/kpi/tasks');
  },
};
