import {
  Bar,
  BarChart,
  CartesianGrid,
  LabelList,
  Legend,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';
import { formatPercent } from '../../utils/formatters';

export default function TotalKpiChart({ data = [] }) {
  return (
    <div className="card">
      <div className="card-header">
        <h2 className="card-title">Grafik Total KPI Karyawan</h2>
      </div>

      <div className="chart-wrapper">
        <ResponsiveContainer width="100%" height={300}>
          <BarChart
            data={data}
            margin={{ top: 20, right: 16, left: 0, bottom: 8 }}
          >
            <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" vertical={false} />
            <XAxis
              dataKey="nama"
              tick={{ fill: '#374151', fontSize: 13 }}
              axisLine={{ stroke: '#d1d5db' }}
              tickLine={false}
            />
            <YAxis
              domain={[0, 110]}
              tickFormatter={(v) => `${v}%`}
              tick={{ fill: '#6b7280', fontSize: 12 }}
              axisLine={false}
              tickLine={false}
            />
            <Tooltip
              formatter={(value) => [formatPercent(value), 'Total KPI']}
              contentStyle={{
                borderRadius: '6px',
                border: '1px solid #d1d5db',
              }}
            />
            <Legend wrapperStyle={{ paddingTop: 10, fontSize: 13 }} />
            <Bar
              dataKey="total_kpi"
              name="Total KPI"
              fill="#10b981"
              radius={[4, 4, 0, 0]}
              maxBarSize={52}
            >
              <LabelList
                dataKey="total_kpi"
                position="top"
                formatter={(v) => `${v}%`}
                style={{ fill: '#1f2937', fontSize: 12, fontWeight: 600 }}
              />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
