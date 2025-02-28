import React, { useEffect, useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import { FaEdit } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles.css';
import StudentMarksChart from './StudentChartPage';

const StudentDetails = () => {
    const [students, setStudents] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [formData, setFormData] = useState({ id: '', name: '', class: '', address: '' });
    const [editing, setEditing] = useState(false);
    const [selectedStudent, setSelectedStudent] = useState(null);
    const [marks, setMarks] = useState(null);
  const [marksLoading, setMarksLoading] = useState(false);
  const [showMarksModal, setMarksModal]=useState(false)
    const navigate=useNavigate();

    useEffect(() => {
        fetchStudents();
    }, []);

    const fetchStudents = async () => {
        console.log("Student Data Fetching starts here")
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
        console.log("Logout button clicked")
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
        const studentData = {
            name: formData.name,
            class: formData.class,
            address: formData.address
          };

        if (!formData.name || !formData.class || !formData.address) {
            
        alert('Please fill in all the fields (Name, Class, Address).');
        return; // Exit the function if validation fails
        }
        try{
            if (editing) {
                // Update student
                await axios.put(`http://localhost:8080/api/students/${formData.id}`, formData);
            } else{
                await axios.post('http://localhost:8080/api/students', studentData);
            }
            fetchStudents(); // Refresh student list
            closeModal();  // Close modal
        
        }catch (error) {
            console.error("Error creating student", error);
        }
        
        setShowModal(false); // Hide form after submission
        setFormData({ id: '', name: '', class: '', address: '' }); // Reset form
    };

    const handleEdit = (student) => {
        setFormData(student); // Set form data to current student details
        setEditing(true); // Enable editing mode
        setShowModal(true); // Show the modal form
    };

    const openCreateModal = () => {
        setFormData({ id: '', name: '', class: '', address: '' });
        setEditing(false);
        setShowModal(true);
    };

    const closeModal = () => {
        setShowModal(false);
    };

    const handleCancel = () => {
        setShowModal(false); // Hide form without any action
        setFormData({ id: '', name: '', class: '', address: '' }); // Reset form
        setEditing(false); // Reset to non-editing mode
    };
    
    return (
        <div class="container">
        <button className ='logout-button' onClick={()=>handleLogout()}> Logout </button>

        {/* Display Image Outside Table */}
        {/* <div className="image-container">
        <img src={studentimage} alt="Background" className="bg-image" />
        </div> */}

        <div className='student-table-container'>
            
    
        <button onClick= {openCreateModal}>Create New Student</button>
            {showModal && (
                <div className='modal'>
                    <div className='modal-content'>
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
                        <div className='form-actions'>
                        <button className='form-button' type="submit">{editing ? 'Update' : 'Create'}</button>
                        <button className='cancel-button' type="button" onClick={handleCancel}>Cancel</button>
                        </div>
                    </form>
                </div>
                </div>
            )}

            <table className='table-container' >
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
                        
                    <tr key={student.id}
                    style={{ cursor: "pointer" }}
                     onClick={() => navigate(`/charts/${student.id}`)}>
                        
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
             {/* Show Pie Chart after fetching marks
            {showMarksModal &&  (
            <div>
              <h2>Marks Distribution for Student ID: {selectedStudent.id}  Name:{selectedStudent.name}</h2>
              {marksLoading ? (
                <p>Loading Marks...</p>
              ) : marks ? (
                <StudentMarksChart marks={marks} />
              ) : (
                <p style={{ color: "red" }}>Failed to load marks</p>
              )} */}
            {/* </div>
          )} */}
      
        </div>
        </div>
    );
}



export default StudentDetails;


 

