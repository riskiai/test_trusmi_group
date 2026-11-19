import { useMemo, useState } from 'react';
import { deriveTaskStatus } from '../../utils/formatters';

export default function TaskStatusTable({ tasks = [] }) {
  const [filterKaryawan, setFilterKaryawan] = useState('ALL');
  const [filterKpi, setFilterKpi] = useState('ALL');
  const [filterStatus, setFilterStatus] = useState('ALL');

  const karyawanList = useMemo(() => {
    const names = new Set();
    tasks.forEach((t) => {
      if (t.karyawan) names.add(t.karyawan);
    });
    return Array.from(names);
  }, [tasks]);

  const filteredTasks = useMemo(() => {
    return tasks.filter((item) => {
      const status = deriveTaskStatus(item.deadline, item.aktual, item.status);
      if (filterKaryawan !== 'ALL' && item.karyawan !== filterKaryawan) return false;
      if (filterKpi !== 'ALL' && item.kpi !== filterKpi) return false;
      if (filterStatus !== 'ALL' && status !== filterStatus) return false;
      return true;
    });
  }, [tasks, filterKaryawan, filterKpi, filterStatus]);

  return (
    <div className="card">
      <div className="card-header card-header-split">
        <h2 className="card-title">Data Tasklist Marketing</h2>

        <div className="filter-bar">
          <select
            className="filter-select"
            value={filterKaryawan}
            onChange={(e) => setFilterKaryawan(e.target.value)}
          >
            <option value="ALL">Semua Karyawan</option>
            {karyawanList.map((name) => (
              <option key={name} value={name}>
                {name}
              </option>
            ))}
          </select>

          <select
            className="filter-select"
            value={filterKpi}
            onChange={(e) => setFilterKpi(e.target.value)}
          >
            <option value="ALL">Semua KPI</option>
            <option value="Sales">Sales</option>
            <option value="Report">Report</option>
          </select>

          <select
            className="filter-select"
            value={filterStatus}
            onChange={(e) => setFilterStatus(e.target.value)}
          >
            <option value="ALL">Semua Status</option>
            <option value="On-Time">Ontime</option>
            <option value="Late">Late</option>
          </select>
        </div>
      </div>

      <div className="table-scroll">
        <table className="task-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Tasklist</th>
              <th>KPI</th>
              <th>Karyawan</th>
              <th>Deadline</th>
              <th>Aktual</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {filteredTasks.map((task, idx) => {
              const status = deriveTaskStatus(task.deadline, task.aktual, task.status);
              const isLate = status === 'Late';

              return (
                <tr key={task.id || task.tasklist}>
                  <td>{idx + 1}</td>
                  <td>{task.tasklist}</td>
                  <td>{task.kpi}</td>
                  <td>{task.karyawan}</td>
                  <td>{task.deadline}</td>
                  <td>{task.aktual}</td>
                  <td>
                    <span className={`status-badge ${isLate ? 'badge-late' : 'badge-ontime'}`}>
                      {isLate ? 'Late' : 'Ontime'}
                    </span>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
}
