import { BrowserRouter as Router, Route, Routes,Navigate} from 'react-router-dom';
import Login from './Components/Login';
import StudentDetails from './Components/Students';
import Register from './Components/Register';
import StudentChartsPage from './Components/StudentChartPage';
import EnableTOTP from './Components/EnableOTP';

function App() {
    // Check if user is logged in (Modify based on session storage or cookies)
    const isAuthenticated = !!localStorage.getItem("user_id"); 
    return (
        <div>
        <h1>Student Management Application</h1>
        <Router>
        <Routes>
                    {/* Public Routes */}
                    <Route path="/" element={<Login />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/students" element={<StudentDetails />} />
                    <Route path="/charts/:id"  element={<StudentChartsPage />} />
                    <Route path="/enable-2fa/:username"  element={<EnableTOTP />} />

                    
                    {/* <Route  path="/students" 
                        element={isAuthenticated ? <StudentDetails /> : <Navigate to="/login" />} 
                    />
                    <Route 
                        path="/charts/:id" 
                        element={isAuthenticated ? <StudentChartsPage /> : <Navigate to="/login" />} 
                    />
                    <Route 
                        path="/enable-2fa" 
                        element={isAuthenticated ? <EnableTOTP /> : <Navigate to="/login" />} 
                    /> */}
                </Routes>
        </Router>
        </div>
    );
}

// function LoginWrapper() {
//     const navigate = useNavigate();

//     const handleLoginSuccess = () => {
//         navigate('/students');
//         axios.defaults.baseURL = 'http://localhost:8080';

//     };

//     return <Login onLogin={handleLoginSuccess} />;
// }

export default App;


// function PageWrapper() {
//     const location = useLocation();
//     let heading = '';

//     if (location.pathname === '/login') {
//         heading = 'Login Page';
//     } else if (location.pathname === '/students') {
//         heading = 'Student Details';
//     } else if (location.pathname === '/register') {
//         heading = 'Admin User Registeration';
//     }else if (location.pathname === '/') {
//         heading = 'Login Page';
//     }

//     return (
//         <div>
//             <h1>Student Management Application</h1>
//             <h2>{heading}</h2>
//             {location.pathname === '/' || '/login' ? <LoginWrapper /> : location.pathname === '/students' ? <StudentDetails /> 
//             : <Register />}

//         </div>
//     );
// }

// import React, { useState } from 'react';
// import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
// import Login from './Components/Login';
// import StudentDetails from './Components/Students';
// import Cookies from 'js-cookie';

// function App() {
//     const [token, setToken] = useState('');
//     const [loggedIn, setLoggedIn] = useState(false);

//     return (
//         <Router>
//             <div className="App">
//                 <h1>School Mangement App</h1>
//                 <Routes>
//                     <Route
//                         path="/"
//                         element={
//                             loggedIn ? <Navigate to="/students" /> : <Login setToken={setToken} setLoggedIn={setLoggedIn} />
//                         }
//                     />
//                     <Route
//                         path="/students"
//                         element={
//                             loggedIn ? <StudentDetails token={token} /> : <Navigate to="/" />
//                         }
//                     />
//                 </Routes>
//             </div>
//         </Router>
//     );
// }


// export default App;



// // // import React from 'react';
// // // import Login from './Components/Login';
// // // import Register from './Components/Register';
// // // import StudentDetails from './Components/Students';

// // // import Cookies from 'js-cookie';

// // // const App = () => {
// // //     const token = Cookies.get('token');

// // //     console.log(token)
// // //     return (
// // //         <div>
// // //             <h1>Student Management App</h1>
// // //             {!token ? (
// // //                 <>
// // //                     <Register />
// // //                     <Login />
// // //                 </>
// // //             ) : (
// // //                 <>
                   
// // //                     <StudentDetails />
// // //                 </>
// // //             )}
// // //         </div>
// // //     );
// // // };

// // // export default App;
