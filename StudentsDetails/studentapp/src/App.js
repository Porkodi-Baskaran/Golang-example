import { BrowserRouter as Router, Route, Routes, useNavigate ,useLocation} from 'react-router-dom';
import Login from './Components/Login';
import StudentDetails from './Components/Students';
import Register from './Components/Register';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<PageWrapper />} />
                <Route path="/students" element={<PageWrapper />} />
                <Route path="/register" element={<PageWrapper />} />
            </Routes>
        </Router>
    );
}

function PageWrapper() {
    const location = useLocation();
    let heading = '';

    if (location.pathname === '/') {
        heading = 'Login Page';
    } else if (location.pathname === '/students') {
        heading = 'Student Details';
    } else if (location.pathname === '/register') {
        heading = 'Admin User Registeration';
    }

    return (
        <div>
            <h1>Student Management Application</h1>
            <h2>{heading}</h2>
            {location.pathname === '/' ? <LoginWrapper /> : location.pathname === '/students' ? <StudentDetails /> : <Register />}
        </div>
    );
}

function LoginWrapper() {
    const navigate = useNavigate();

    const handleLoginSuccess = () => {
        navigate('/students');
    };

    return <Login onLogin={handleLoginSuccess} />;
}

export default App;


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
