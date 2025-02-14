import React, { useEffect, useState } from 'react';
// import { useNavigate } from 'react-router-dom';
// import CreateStudent from './Components/Createstudent';
import { FaTrash } from "react-icons/fa";
import { FaEdit } from "react-icons/fa";

import axios from "axios";

const StudentDetails = () => {
    const [students, setStudents] = useState([]);
    const [newStudent, setNewStudent] = useState({id:'', name: '', class: '', address: '' });
    const [showForm, setShowForm] = useState(false);
    const [editStudentId, setEditStudentId] = useState(null);
    const [editStudent, setEditStudent] = useState({ id: '', name: '', class: '', address: '' });
    

    useEffect(() => {
        const fetchStudents = async () => {
            try {
                const response = await axios.get("/students");
                setStudents(response.data.data);
            } catch (error) {
                console.error("Error fetching students:", error);
            }
        };

        fetchStudents();
    }, []);
    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setNewStudent(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const handleFormSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8080/students', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(newStudent),
        });
        console.log(response.data)

        if (response.ok) {
            const result = await response.json();
            alert(result.status);
            setStudents([...students, newStudent]);
            setShowForm(false); // Hide the form after submission
            setNewStudent({ id: '', name: '', class: '', address: '' }); // Reset form fields
        } else {
            alert('Failed to create student');
        }
    };
    const handleUpdateStudent = (student) => {
        setEditStudentId(student.ID);
        setEditStudent(student);
    };

    const handleEditFormSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch(`http://localhost:8080/students/${editStudent.id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(editStudent),
        });

        if (response.ok) {
            const result = await response.json();
            alert(result.status);
            setStudents(students.map(student => student.ID === editStudent.id ? editStudent : student));
            setEditStudentId(null); // Exit edit mode
            setEditStudent({ id: '', name: '', class: '', address: '' }); // Reset form fields
        } else {
            alert('Failed to update student');
        }
    };

    const handleDeleteStudent = async (id) => {
        console.log(id)
        const response = await fetch(`http://localhost:8080/students/${id}`, {
            method: 'DELETE',
            credentials: 'include',
        });
        console.log(response)

        if (response.ok) {
            const result = await response.json();
            console.log("REsult:",result)
            
            alert(result.status);
            setStudents(students.filter(student => student.ID !== id));
        } else {
            alert('Failed to delete student');
        }
    };

    return (
        <div className="student-table-container">
            <button onClick={() => setShowForm(!showForm)}>
                {showForm ? 'Cancel' : 'Create New Student'}
            </button>
            {showForm && (
                <form onSubmit={handleFormSubmit}>
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
                    <button type="submit">Create Student</button>
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
                    
                <table >
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
                         return(<tr key={student.id}>
                         <td>{student.id}</td>
                         <td>{student.name}</td> 
                         <td>{student.class}</td>
                         <td>{student.address}</td>
                         <td>
                            <button onClick={() => handleUpdateStudent(()=>(student))}><FaEdit /></button>
                            <button onClick={() => handleDeleteStudent(()=>(student.id))}><FaTrash /></button>
                         </td>
                     </tr>)
                        
 })}
                 </tbody>
                    
                </table>
            </div>
        
    );
}

export default StudentDetails;


 

