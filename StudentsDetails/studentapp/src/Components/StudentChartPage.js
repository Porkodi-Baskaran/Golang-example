import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import StudentMarksChart from "./StudentMarksChart";

const StudentChartsPage = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [marks, setMarks] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get(`http://localhost:8080/api/marks/${id}`)
      .then((response) => {
        setMarks(response.data.data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching marks:", err);
        setError("Failed to load marks");
        setLoading(false);
      });
  }, [id]);

  return (
    <div className="charts-container">
      <button className="back-button" onClick={() => navigate("/students")}>Back to Student List</button>
      <h2>Student Marks Analysis (ID: {id})</h2>
      {loading ? (
        <p>Loading...</p>
      ) : error ? (
        <p style={{ color: "red" }}>{error}</p>
      ) : (
        <StudentMarksChart marks={marks} />
      )}
    </div>
  );
};

export default StudentChartsPage;



// import React, { useEffect, useState } from "react";
// import { Chart } from "react-google-charts";
// import axios from "axios";

// const StudentMarksChart = (Studid ) => {
//     const [studentMarks, setStudentMarks] = useState([]);
//     const [loading, setLoading] = useState(true);
//     const [error, setError] = useState(null);
    
//     useEffect(() => {
//             axios
//               .get("http://localhost:8080/api/marks/${id}") 
//               .then((response) => {
//                 console.log(response)
//                 if (response.status === 200) {
//                     console.log("Response:",response.data.data)
//                   setStudentMarks(response.data.data); 
//                 }
//                 setLoading(false);
//               })
//               .catch((err) => {
//                 console.error("Error fetching marks:", err);
//                 setError("Failed to load data");
//                 setLoading(false);
//               });
//           }, []);
        
//           const chartOptions = {
//             is3D: true, 
//           }
  
//             const chartData = [
//               ["Subject", "Marks"], 
//               ["Maths", studentMarks.maths],
//               ["Science", studentMarks.science],
//               ["English", studentMarks.english],
//               ["Tamil", studentMarks.tamil],
//               ["Social Science", studentMarks.socialSci],
//             ];

//             return (
//               <div key={studentMarks.studid} style={{ marginBottom: "30px" }}>
//                 <h3>Student ID: {studentMarks.studid}</h3>
//                 <Chart chartType="PieChart" width="100%" height="400px" 
//                 data={chartData} 
//                 options={{ ...chartOptions, title: `Marks for Student ID: ${studentMarks.studid} \nClass: ${studentMarks.class}` }} />
//               </div>
//             );
// };

// export default StudentMarksChart;
// ;