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

export default function OntimeLateEmployeeChart({ perKaryawan = [] }) {
  return (
    <div className="card">
      <div className="card-header">
        <h2 className="card-title">Grafik Ontime vs Late per Karyawan</h2>
      </div>

      <div className="chart-wrapper">
        <ResponsiveContainer width="100%" height={280}>
          <BarChart
            data={perKaryawan}
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
              allowDecimals={false}
              domain={[0, 5]}
              tick={{ fill: '#6b7280', fontSize: 12 }}
              axisLine={false}
              tickLine={false}
            />
            <Tooltip />
            <Legend wrapperStyle={{ paddingTop: 10, fontSize: 13 }} />
            <Bar
              dataKey="total_ontime"
              name="Ontime"
              fill="#10b981"
              radius={[4, 4, 0, 0]}
              maxBarSize={42}
            >
              <LabelList
                dataKey="total_ontime"
                position="top"
                style={{ fill: '#1f2937', fontSize: 12, fontWeight: 600 }}
              />
            </Bar>
            <Bar
              dataKey="total_late"
              name="Late"
              fill="#ef4444"
              radius={[4, 4, 0, 0]}
              maxBarSize={42}
            >
              <LabelList
                dataKey="total_late"
                position="top"
                style={{ fill: '#1f2937', fontSize: 12, fontWeight: 600 }}
              />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
