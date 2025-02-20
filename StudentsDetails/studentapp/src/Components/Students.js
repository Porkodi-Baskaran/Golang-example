import React, { useEffect, useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import { FaEdit } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles.css';

const StudentDetails = () => {
    const [students, setStudents] = useState([]);
    const [showForm, setShowForm] = useState(false);
    const [formData, setFormData] = useState({ id: '', name: '', class: '', address: '' });
    const [editing, setEditing] = useState(false);
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
    const handleLogout = async () => {
        try {
        const res = await axios.get('http://localhost:8080/api/logout', {
            withCredentials: true // This allows cookies to be sent and received
        });
        console.log(res.data);
        navigate('/');
        } catch (error) {
        console.error('Logout error:', error);
        }
    };

    const handleCreateOrUpdate = async () => {
        if (editing) {
            // Update student
            try {
                await axios.put(`http://localhost:8080/api/students/${formData.id}`, formData);
                fetchStudents(); // Refetch students after update
            } catch (error) {
                console.error("Error updating student", error);
            }
        } else {
            console.log("FormData:",formData)
            // Create new student
            try {
                await axios.post('http://localhost:8080/api/students', formData);
                // formData.name,formData.class,formData.address);
                fetchStudents(); // Refetch students after creation
            } catch (error) {
                console.error("Error creating student", error);
            }
        }
        setShowForm(false); // Hide form after submission
        setFormData({ id: '', name: '', class: '', address: '' }); // Reset form
    };

    const handleEdit = (student) => {
        setFormData(student); // Set form data to current student details
        setEditing(true); // Enable editing mode
        setShowForm(true); // Show the form
    };

    const handleCancel = () => {
        setShowForm(false); // Hide form without any action
        setFormData({ id: '', name: '', class: '', address: '' }); // Reset form
        setEditing(false); // Reset to non-editing mode
    };
    
    return (
        <div class="container">
              <button class="logout-button" onclick={handleLogout}>Logout</button>
        <div className='student-table-container'>
            
        
        <button onClick={() => setShowForm(true)}>New</button>
            {showForm && (
                <div>
                    <h3>{editing ? 'Edit Student' : 'Create New Student'}</h3>
                    <form onSubmit={(e) => { e.preventDefault(); handleCreateOrUpdate(); }}>
                        <input
                            type="text"
                            value={formData.name}
                            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                            placeholder="Name"
                        />
                        <input
                            type="text"
                            value={formData.class}
                            onChange={(e) => setFormData({ ...formData, class: e.target.value })}
                            placeholder="Class"
                        />
                        <input
                            type="text"
                            value={formData.address}
                            onChange={(e) => setFormData({ ...formData, address: e.target.value })}
                            placeholder="Address"
                        />
                        <button className='edit-button' type="submit">{editing ? 'Update' : 'Create'}</button>
                        <button className='delete-button' type="button" onClick={handleCancel}>Cancel</button>
                    </form>
                </div>
            )}

            <table  >
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
                            onClick={ () => handleEdit(student)}/>
                            

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
        </div>
    );

}



export default StudentDetails;


 

