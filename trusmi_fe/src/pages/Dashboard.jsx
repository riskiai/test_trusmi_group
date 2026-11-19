import { useState } from 'react';
import OntimeLateEmployeeChart from '../components/charts/OntimeLateEmployeeChart';
import OntimeLatePieChart from '../components/charts/OntimeLatePieChart';
import SalesVsReportChart from '../components/charts/SalesVsReportChart';
import TotalKpiChart from '../components/charts/TotalKpiChart';
import { EmptyState, ErrorState, LoadingState } from '../components/common/StatusStates';
import SummaryCard from '../components/common/SummaryCard';
import Header from '../components/layout/Header';
import KpiSummaryTable from '../components/tables/KpiSummaryTable';
import TaskStatusTable from '../components/tables/TaskStatusTable';
import { useDashboardData } from '../hooks/useDashboardData';
import { formatPercent } from '../utils/formatters';

export default function Dashboard() {
  const [activeTab, setActiveTab] = useState('dashboard1');
  const {
    kpiRows,
    kpiMeta,
    ontimeLate,
    tasks,
    loading,
    refreshing,
    error,
    refresh,
  } = useDashboardData();

  return (
    <div className="dashboard-shell">
      <Header
        activeTab={activeTab}
        onTabChange={setActiveTab}
        onRefresh={refresh}
        refreshing={refreshing}
      />

      <main className="dashboard-content">
        {loading && <LoadingState />}

        {!loading && error && <ErrorState message={error} onRetry={refresh} />}

        {!loading && !error && kpiRows.length === 0 && <EmptyState />}

        {!loading && !error && kpiRows.length > 0 && activeTab === 'dashboard1' && (
          <>
            <p className="tab-description">
              Menampilkan rekap perhitungan KPI Sales dan Report untuk setiap karyawan berdasarkan target, realisasi, dan potongan keterlambatan.
            </p>

            <div className="summary-grid">
              <SummaryCard
                title="Total Karyawan"
                value={kpiMeta.total_karyawan}
              />
              <SummaryCard
                title="Rata-rata KPI"
                value={formatPercent(kpiMeta.rata_rata_kpi)}
              />
              <SummaryCard
                title="KPI Tertinggi"
                value={`${formatPercent(kpiMeta.kpi_tertinggi)} (${kpiMeta.karyawan_terbaik})`}
              />
              <SummaryCard
                title="KPI Terendah"
                value={formatPercent(kpiMeta.kpi_terendah)}
              />
            </div>

            <div className="charts-grid-2">
              <SalesVsReportChart data={kpiRows} />
              <TotalKpiChart data={kpiRows} />
            </div>

            <KpiSummaryTable data={kpiRows} />
          </>
        )}

        {!loading && !error && kpiRows.length > 0 && activeTab === 'dashboard2' && (
          <>
            <p className="tab-description">
              Menampilkan perbandingan jumlah serta persentase tasklist yang selesai tepat waktu (Ontime) dan terlambat (Late) beserta daftar tasklist.
            </p>

            <div className="summary-grid">
              <SummaryCard
                title="Total Task"
                value={ontimeLate.total_task}
              />
              <SummaryCard
                title="Task Ontime"
                value={ontimeLate.total_ontime}
              />
              <SummaryCard
                title="Task Late"
                value={ontimeLate.total_late}
              />
              <SummaryCard
                title="Persentase Ontime / Late"
                value={`${formatPercent(ontimeLate.persentase_ontime)} / ${formatPercent(ontimeLate.persentase_late)}`}
              />
            </div>

            <div className="charts-grid-2">
              <OntimeLatePieChart ontimeLate={ontimeLate} />
              <OntimeLateEmployeeChart perKaryawan={ontimeLate.per_karyawan} />
            </div>

            <TaskStatusTable tasks={tasks} />
          </>
        )}
      </main>
    </div>
  );
}
