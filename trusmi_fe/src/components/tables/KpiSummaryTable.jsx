import {
  formatPercent,
  formatSignedPercent,
} from '../../utils/formatters';

export default function KpiSummaryTable({ data = [] }) {
  return (
    <div className="card">
      <div className="card-header">
        <h2 className="card-title">Tabel Perhitungan KPI Marketing</h2>
      </div>

      <div className="table-scroll">
        <table className="kpi-table">
          <thead>
            <tr>
              <th rowSpan={2} className="th-nama">
                Nama
              </th>
              <th colSpan={6} className="th-group-sales">
                Sales
              </th>
              <th colSpan={6} className="th-group-report">
                Report
              </th>
              <th rowSpan={2} className="th-group-kpi">
                KPI
              </th>
            </tr>
            <tr>
              <th className="th-sub-sales">Target</th>
              <th className="th-sub-sales">Actual</th>
              <th className="th-sub-sales">Pencapaian</th>
              <th className="th-sub-sales">Bobot Sales</th>
              <th className="th-sub-sales">Late Sales</th>
              <th className="th-sub-sales">Total Bobot</th>

              <th className="th-sub-report">Target</th>
              <th className="th-sub-report">Actual</th>
              <th className="th-sub-report">Pencapaian</th>
              <th className="th-sub-report">Actual Bobot</th>
              <th className="th-sub-report">Late Report</th>
              <th className="th-sub-report">Total Bobot</th>
            </tr>
          </thead>
          <tbody>
            {data.map((row) => (
              <tr key={row.karyawan_id || row.nama}>
                <td>
                  <strong>{row.nama}</strong>
                </td>

                <td className="td-num">{row.sales_target}</td>
                <td className="td-num">{row.sales_actual}</td>
                <td className="td-num">{formatPercent(row.sales_pencapaian)}</td>
                <td className="td-num">{formatPercent(row.sales_bobot)}</td>
                <td className={`td-num ${Number(row.sales_late) < 0 ? 'text-danger' : ''}`}>
                  {formatSignedPercent(row.sales_late)}
                </td>
                <td className="td-num">
                  <strong>{formatPercent(row.sales_total_bobot)}</strong>
                </td>

                <td className="td-num">{row.report_target}</td>
                <td className="td-num">{row.report_actual}</td>
                <td className="td-num">{formatPercent(row.report_pencapaian)}</td>
                <td className="td-num">{formatPercent(row.report_bobot)}</td>
                <td className={`td-num ${Number(row.report_late) < 0 ? 'text-danger' : ''}`}>
                  {formatSignedPercent(row.report_late)}
                </td>
                <td className="td-num">
                  <strong>{formatPercent(row.report_total_bobot)}</strong>
                </td>

                <td className="td-num td-final-kpi">
                  <strong>{formatPercent(row.total_kpi)}</strong>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
