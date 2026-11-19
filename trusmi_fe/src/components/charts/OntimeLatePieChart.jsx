import {
  Cell,
  Legend,
  Pie,
  PieChart,
  ResponsiveContainer,
  Tooltip,
} from 'recharts';
import { formatPercent } from '../../utils/formatters';

const COLORS = ['#10b981', '#ef4444'];

export default function OntimeLatePieChart({ ontimeLate }) {
  const totalOntime = ontimeLate?.total_ontime || 0;
  const totalLate = ontimeLate?.total_late || 0;
  const pctOntime = ontimeLate?.persentase_ontime || 0;
  const pctLate = ontimeLate?.persentase_late || 0;

  const pieData = [
    {
      name: 'Ontime',
      value: totalOntime,
      percent: pctOntime,
    },
    {
      name: 'Late',
      value: totalLate,
      percent: pctLate,
    },
  ];

  return (
    <div className="card">
      <div className="card-header">
        <h2 className="card-title">Grafik Persentase Ontime vs Late</h2>
      </div>

      <div className="pie-layout">
        <div className="chart-wrapper">
          <ResponsiveContainer width="100%" height={280}>
            <PieChart>
              <Pie
                data={pieData}
                cx="50%"
                cy="50%"
                innerRadius={60}
                outerRadius={95}
                paddingAngle={3}
                dataKey="value"
                label={({ name, percent }) => `${name}: ${Number(percent).toFixed(2)}%`}
              >
                {pieData.map((entry, idx) => (
                  <Cell key={entry.name} fill={COLORS[idx % COLORS.length]} />
                ))}
              </Pie>
              <Tooltip
                formatter={(value, name, props) => [
                  `${value} Task (${formatPercent(props?.payload?.percent)})`,
                  name,
                ]}
              />
              <Legend verticalAlign="bottom" height={36} />
            </PieChart>
          </ResponsiveContainer>
        </div>

        <div className="pie-stats-panel">
          <div className="pie-stat-item stat-ontime">
            <div className="pie-stat-label">Ontime</div>
            <div className="pie-stat-numbers">
              <strong>{totalOntime} Task</strong>
              <span>{formatPercent(pctOntime)}</span>
            </div>
          </div>

          <div className="pie-stat-item stat-late">
            <div className="pie-stat-label">Late</div>
            <div className="pie-stat-numbers">
              <strong>{totalLate} Task</strong>
              <span>{formatPercent(pctLate)}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
