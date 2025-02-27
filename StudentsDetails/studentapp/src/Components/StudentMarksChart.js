import React from "react";
import { Chart } from "react-google-charts";

const StudentMarksChart = ({ marks }) => {
  if (!marks) return null;

  const chartData = [
    ["Subject", "Marks"],
    ["Maths", marks.maths],
    ["Science", marks.science],
    ["English", marks.english],
    ["Tamil", marks.tamil],
    ["Social Science", marks.socialSci],
  ];

  return (
    <div className="chart-wrapper">
      <div className="chart-box">
      <h3>Pie Chart</h3>
      <Chart chartType="PieChart" width="100%" height="200px" data={chartData} options={{ is3D: true }} />
      </div>

      <div className="chart-box">
      <h3>Bar Chart</h3>
      <Chart chartType="BarChart" width="100%" height="200px" data={chartData}  />
      </div>

      <div className="chart-box">
      <h3>Column Chart</h3>
      <Chart chartType="ColumnChart" width="100%" height="200px" data={chartData} />
      </div>

      <div className="chart-box">
      <h3>Line Chart</h3>
      <Chart chartType="LineChart" width="100%" height="200px" data={chartData} />
      </div>
    </div>
  );
};

export default StudentMarksChart;
