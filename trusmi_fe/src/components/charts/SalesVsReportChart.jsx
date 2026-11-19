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

export default function SalesVsReportChart({ data = [] }) {
  return (
    <div className="card">
      <div className="card-header">
        <h2 className="card-title">Grafik Bobot Sales vs Report</h2>
      </div>

      <div className="chart-wrapper">
        <ResponsiveContainer width="100%" height={300}>
          <BarChart
            data={data}
            margin={{ top: 20, right: 16, left: 0, bottom: 8 }}
            barGap={8}
          >
            <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" vertical={false} />
            <XAxis
              dataKey="nama"
              tick={{ fill: '#374151', fontSize: 13 }}
              axisLine={{ stroke: '#d1d5db' }}
              tickLine={false}
            />
            <YAxis
              domain={[0, 60]}
              tickFormatter={(v) => `${v}%`}
              tick={{ fill: '#6b7280', fontSize: 12 }}
              axisLine={false}
              tickLine={false}
            />
            <Tooltip
              formatter={(value) => formatPercent(value)}
              contentStyle={{
                borderRadius: '6px',
                border: '1px solid #d1d5db',
              }}
            />
            <Legend wrapperStyle={{ paddingTop: 10, fontSize: 13 }} />
            <Bar
              dataKey="sales_total_bobot"
              name="Total Bobot Sales"
              fill="#3b82f6"
              radius={[4, 4, 0, 0]}
              maxBarSize={44}
            >
              <LabelList
                dataKey="sales_total_bobot"
                position="top"
                formatter={(v) => `${v}%`}
                style={{ fill: '#1f2937', fontSize: 12, fontWeight: 600 }}
              />
            </Bar>
            <Bar
              dataKey="report_total_bobot"
              name="Total Bobot Report"
              fill="#ef4444"
              radius={[4, 4, 0, 0]}
              maxBarSize={44}
            >
              <LabelList
                dataKey="report_total_bobot"
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
