import React, { useEffect, useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import { FaEdit } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles.css';

const StudentDetails = () => {
    const [students, setStudents] = useState([]);
    const [newStudent, setNewStudent] = useState({id:'', name: '', class: '', address: '' });
    const [showForm, setShowForm] = useState(false);
    const [editStudentId, setEditStudentId] = useState(null);
    const [editStudent, setEditStudent] = useState({ id: '', name: '', class: '', address: '' });
    const navigate=useNavigate();


    useEffect(() => {
    
        fetchStudents();
    }, []);

    const fetchStudents = async () => {
        try {
            // axios.defaults.baseURL = 'http://localhost:8080';
            const response = await axios.get("http://localhost:8080/api/students", { withCredentials: true });
            console.log("Student Response:",response.data.data)
            setStudents(response.data.data);
        } catch (error) {
            console.error("Error fetching students:", error);
            navigate('/login')
        }
    };

    
    // Handle delete data
    const handleDelete = (id,name) => {
        const confirmed = window.confirm(`Are you sure you want to delete this student?\n
            \nID: ${id}\nName: ${name}`);
        if(confirmed){
            axios.delete("http://localhost:8080/api/students/" +id)
            .then((res) => {
                console.log("Deleted student:", res);
                setStudents(students.filter(student => student.id !== id));
                // fetchStudents();
            })
            .catch((err) => {
                console.log(err);
                
            });
        }
    };
    // const handleLogout = async () => {
    //     try {
    //     const res = await axios.get('http://localhost:8080/api/logout', {
    //         withCredentials: true // This allows cookies to be sent and received
    //     });
    //     console.log(res.data);
    //     navigate('/');
    //     } catch (error) {
    //     console.error('Logout error:', error);
    //     }
    // };


    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setNewStudent(prevState => ({
            ...prevState,
            [name]: value
        }));
    };


    const handleFormSubmit = async () => {
        console.log("newstudent:",newStudent)
            try {
              const res = await axios.post('http://localhost:8080/api/students', {
                newStudent
              }, {
                withCredentials: true // This allows cookies to be sent and received
              });
              console.log("Response from CREATE API:", res)
              setStudents([...students, newStudent]);
              setShowForm(false); // Hide the form after submission
              setNewStudent({ id: '', name: '', class: '', address: '' }); // Reset form fields
              
            } catch (error) {
              console.error('Error creating student:', error);
              alert('Failed to create student');
            }
          };
    
    const handleUpdateStudent = (student) => {
        setEditStudentId(student.ID);
        setEditStudent(student);
    };

    const handleEditFormSubmit = async (e) => {
        try {
                  const res = await axios.put(`http://localhost:8080/api/students/${editStudent.id}`, {
                    editStudent
                  }, {
                    withCredentials: true // This allows cookies to be sent and received
                  });
                  console.log(res.data);
                  setStudents(students.map(student => student.ID === editStudent.id ? editStudent : student));
                  setEditStudentId(null); // Exit edit mode
                  setEditStudent({ id: '', name: '', class: '', address: '' }); // Reset form fields
                } catch (error) {
                  console.error('Error saving student details:', error);
                  alert('Failed to update student');
                }
              };

    
    return (
        <div className='Student-container'>
        <button onClick={() => setShowForm(!showForm)}>
            {showForm ? 'Cancel' : 'Create New Student'}
        </button>
            {showForm && (
                <form className='student-form' onSubmit={handleFormSubmit}>
                    <input
                        type="text"
                        name="id"
                        value={newStudent.id}
                        onChange={handleInputChange}
                        placeholder="ID"
                    />
                    <input
                        type="text"
                        name="name"
                        value={newStudent.name}
                        onChange={handleInputChange}
                        placeholder="Name"
                    />
                    <input
                        type="text"
                        name="class"
                        value={newStudent.class}
                        onChange={handleInputChange}
                        placeholder="Class"
                    />
                    <input
                        type="text"
                        name="address"
                        value={newStudent.address}
                        onChange={handleInputChange}
                        placeholder="Address"
                    />
                    <button className='create-button' type="submit">Create Student</button>
                </form>
            )}
                {editStudentId && (
                <form onSubmit={handleEditFormSubmit}>
                    <input
                        type="text"
                        name="id"
                        value={editStudent.id}
                        onChange={(e) => handleInputChange(e, 'edit')}
                        placeholder="ID"
                        readOnly
                    />
                    <input
                        type="text"
                        name="name"
                        value={editStudent.name}
                        onChange={(e) => handleInputChange(e, 'edit')}
                        placeholder="Name"
                    />
                    <input
                        type="text"
                        name="class"
                        value={editStudent.class}
                        onChange={(e) => handleInputChange(e, 'edit')}
                        placeholder="Class"
                    />
                    <input
                        type="text"
                        name="address"
                        value={editStudent.address}
                        onChange={(e) => handleInputChange(e, 'edit')}
                        placeholder="Address"
                    />
                    <button type="submit">Update Student</button>
                    <button onClick={() => setEditStudentId(null)}>Cancel</button>
                </form>
            )}

            <table  className='student-table-container'>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Student Name</th>
                        <th>Class</th>
                        <th>Address</th>
                        <th>Actions</th>
                        

                    </tr>
                </thead> 
                <tbody>
                    {students.map(student => {
                        return(
                        <tr key={student.id}>
                        <td>{student.id}</td>
                        <td>{student.name}</td> 
                        <td>{student.class}</td>
                        <td>{student.address}</td>
                        <td>
                            <FaEdit 
                            className='edit-button'
                            role="button"
                            onClick={() => handleUpdateStudent(()=>(student))}/>
                            

                            <FaTrash 
                            className='delete-button'
                            role="button"
                            onClick={()=>{handleDelete(student.id,student.name)}}/> 
                        </td>
                        </tr>)             
                    })}
                </tbody>      
            </table>
            
            
        </div>
    );

}



export default StudentDetails;


 

