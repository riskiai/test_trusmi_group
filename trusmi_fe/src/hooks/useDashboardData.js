import { useCallback, useEffect, useState } from 'react';
import { kpiApi } from '../services/api';

export function useDashboardData() {
  const [kpiRows, setKpiRows] = useState([]);
  const [kpiMeta, setKpiMeta] = useState({
    total_karyawan: 0,
    rata_rata_kpi: 0,
    kpi_tertinggi: 0,
    karyawan_terbaik: '-',
    kpi_terendah: 0,
  });
  const [ontimeLate, setOntimeLate] = useState({
    total_task: 0,
    total_ontime: 0,
    total_late: 0,
    persentase_ontime: 0,
    persentase_late: 0,
    per_karyawan: [],
    per_kpi: [],
  });
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [error, setError] = useState(null);

  const fetchAll = useCallback(async (isRefresh = false) => {
    try {
      if (isRefresh) {
        setRefreshing(true);
      } else {
        setLoading(true);
      }
      setError(null);

      const [dashboard1Res, dashboard2Res] = await Promise.all([
        kpiApi.getDashboard1(),
        kpiApi.getDashboard2(),
      ]);

      setKpiRows(Array.isArray(dashboard1Res?.data) ? dashboard1Res.data : []);
      if (dashboard1Res?.meta) {
        setKpiMeta(dashboard1Res.meta);
      }

      if (dashboard2Res?.data) {
        setOntimeLate({
          total_task: dashboard2Res.data.total_task ?? 0,
          total_ontime: dashboard2Res.data.total_ontime ?? 0,
          total_late: dashboard2Res.data.total_late ?? 0,
          persentase_ontime: dashboard2Res.data.persentase_ontime ?? 0,
          persentase_late: dashboard2Res.data.persentase_late ?? 0,
          per_karyawan: Array.isArray(dashboard2Res.data.per_karyawan)
            ? dashboard2Res.data.per_karyawan
            : [],
          per_kpi: Array.isArray(dashboard2Res.data.per_kpi)
            ? dashboard2Res.data.per_kpi
            : [],
        });
        setTasks(Array.isArray(dashboard2Res.data.tasks) ? dashboard2Res.data.tasks : []);
      }
    } catch (err) {
      setError(
        err?.message ||
          'Gagal terhubung ke server API Backend. Pastikan backend berjalan.'
      );
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchAll(false);
  }, [fetchAll]);

  return {
    kpiRows,
    kpiMeta,
    ontimeLate,
    tasks,
    loading,
    refreshing,
    error,
    refresh: () => fetchAll(true),
  };
}
